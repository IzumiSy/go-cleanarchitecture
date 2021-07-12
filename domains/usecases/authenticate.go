package usecases

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/authentication"
)

type AuthenticateOutputPort interface {
	domains.OutputPort
	Write(auth models.Authentication)
}

type AuthenticateParam struct {
	Email    string
	Password string
}

type authenticateUsecase struct {
	outputPort        AuthenticateOutputPort
	authenticationDao domains.AuthenticationRepository
	logger            domains.Logger
}

func AuthenticateUsecase(
	outputPort AuthenticateOutputPort,
	authenticationDao domains.AuthenticationRepository,
	logger domains.Logger,
) authenticateUsecase {
	return authenticateUsecase{outputPort, authenticationDao, logger}
}

func (usecase authenticateUsecase) Execute(params AuthenticateParam) {
	// [ユーザーの認証を行うユースケース]
	// "ログイン"でも命名はよかったが、今後外部APIとして認証を実装することを考えると
	// あえて抽象化して"認証"と表現したくなったのでこの命名としている。

	var (
		USER_NOT_FOUND   = errors.Invalid("User not found")
		INVALID_PASSWORD = errors.Invalid("Invalid password")
	)

	email, err := authentication.NewEmail(params.Email)
	if err.NotNil() {
		usecase.logger.Warn(err.Error())
		usecase.outputPort.Raise(err)
		return
	}

	auth, err, exists := usecase.authenticationDao.GetByEmail(email)
	if err.NotNil() {
		usecase.logger.Error(err.Error())
		usecase.outputPort.Raise(err)
		return
	}

	if !exists {
		usecase.outputPort.Raise(USER_NOT_FOUND)
		return
	}

	loginHash := authentication.NewHash(params.Password)
	if auth.Hash().Value() != loginHash.Value() {
		usecase.outputPort.Raise(INVALID_PASSWORD)
		return
	}

	//
	// TODO: ここで認証トークンを払いだす
	//

	usecase.outputPort.Write(auth)
}
