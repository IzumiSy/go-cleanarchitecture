package usecases

import (
	"fmt"

	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/authentication"
	"go-cleanarchitecture/domains/models/user"
)

type SignupOutputPort interface {
	domains.OutputPort
	Write(auth models.Authentication)
}

type SignupParam struct {
	Email    string
	Password string
	UserName string
}

type SignupUsecase struct {
	OutputPort        SignupOutputPort
	AuthenticationDao domains.AuthenticationRepository
	Logger            domains.Logger
	Publisher         domains.EventPublisher
}

func (uc SignupUsecase) Build(params SignupParam) domains.UnauthorizedUsecase {
	return domains.NewUnauthorizedUsecase(func() {
		// [新規ユーザーのサインナップを行うユースケース]
		// バリデーションルールは以下
		// - すでに同じメールアドレスで登録されている場合にはサインナップ不可

		var (
			EMAIL_INVALID = errors.Preconditional("Email must not be duplicated")
		)

		email, err := authentication.NewEmail(params.Email)
		if err.NotNil() {
			uc.Logger.Warn(err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		_, err, exists := uc.AuthenticationDao.GetByEmail(email)
		if err.NotNil() {
			uc.Logger.Error(err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		if exists {
			uc.OutputPort.Raise(EMAIL_INVALID)
			return
		}

		userName, err := user.NewName(params.UserName)
		if err.NotNil() {
			uc.Logger.Warn(err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		hash := authentication.NewHash(params.Password)
		auth := models.NewAuthentication(email, hash, userName)
		if err = uc.AuthenticationDao.Store(auth); err.NotNil() {
			uc.Logger.Error(err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		event := UserSignedUpEvent{
			UserID: auth.User().ID().String(),
			Email:  auth.Email().Value(),
			Name:   auth.User().Name().Value(),
		}
		if err := uc.Publisher.Publish(event); err != nil {
			uc.Logger.Error(fmt.Sprintf("Failed publishing event: %s", err.Error()))
		}

		uc.OutputPort.Write(auth)
	})
}

type UserSignedUpEvent struct {
	UserID string
	Email  string
	Name   string
}

func (UserSignedUpEvent) ID() domains.DomainEventID {
	return domains.UserSignedUp
}
