package usecases

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/authentication"
	"go-cleanarchitecture/domains/models/user"
)

type SignupOutputPort interface {
	domains.OutputPort
	Write(todo models.Authentication)
}

type SignupParam struct {
	Email    string
	Password string
	UserName string
}

type signupUsecase struct {
	outputPort        SignupOutputPort
	authenticationDao domains.AuthenticationRepository
	logger            domains.Logger
}

func NewSignupUsecase(
	outputPort SignupOutputPort,
	authenticationDao domains.AuthenticationRepository,
	logger domains.Logger,
) signupUsecase {
	return signupUsecase{outputPort, authenticationDao, logger}
}

func (usecase signupUsecase) Execute(params SignupParam) {
	// [新規ユーザーのサインナップを行うユースケース]
	// バリデーションルールは以下
	// - すでに同じメールアドレスで登録されている場合にはサインナップ不可

	var (
		EMAIL_INVALID = errors.Invalid("Email must not be duplicated")
	)

	email, err := authentication.NewEmail(params.Email)
	if err.NotNil() {
		usecase.logger.Warn(err.Error())
		usecase.outputPort.Raise(err)
		return
	}

	_, err, exists := usecase.authenticationDao.GetByEmail(email)
	if err.NotNil() {
		usecase.logger.Error(err.Error())
		usecase.outputPort.Raise(err)
		return
	}

	if exists {
		usecase.outputPort.Raise(EMAIL_INVALID)
		return
	}

	err, userName := user.NewName(params.UserName)
	if err.NotNil() {
		usecase.logger.Warn(err.Error())
		usecase.outputPort.Raise(err)
		return
	}

	hash := authentication.NewHash(params.Password)
	auth := models.NewAuthentication(email, hash, userName)
	if err = usecase.authenticationDao.Store(auth); err.NotNil() {
		usecase.logger.Error(err.Error())
		usecase.outputPort.Raise(err)
		return
	}

	usecase.outputPort.Write(auth)
}
