package usecases

import (
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/authentication"
)

type AuthenticateOutputPort interface {
	domains.OutputPort
	Write(session models.Session)
}

type AuthenticateParam struct {
	Email    string
	Password string
}

type AuthenticateUsecase struct {
	OutputPort        AuthenticateOutputPort
	AuthenticationDao domains.AuthenticationRepository
	SessionDao        domains.SessionRepository
	Logger            domains.Logger
}

func (uc AuthenticateUsecase) Build(params AuthenticateParam) domains.UnauthorizedUsecase {
	return domains.NewUnauthorizedUsecase(func() {
		// [ユーザーの認証を行うユースケース]
		// "ログイン"でも命名はよかったが、今後外部APIとして認証を実装したりする可能性を考えると
		// 人間以外のアクタも考慮し抽象化して"認証"と表現したくなったのでこの命名としている。

		var (
			USER_NOT_FOUND   = errors.Invalid("User not found")
			INVALID_PASSWORD = errors.Invalid("Invalid password")
		)

		email, err := authentication.NewEmail(params.Email)
		if err.NotNil() {
			uc.Logger.Warn(err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		auth, err, exists := uc.AuthenticationDao.GetByEmail(email)
		if err.NotNil() {
			uc.Logger.Error(err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		if !exists {
			uc.OutputPort.Raise(USER_NOT_FOUND)
			return
		}

		loginHash := authentication.NewHash(params.Password)
		if auth.Hash().Value() != loginHash.Value() {
			uc.OutputPort.Raise(INVALID_PASSWORD)
			return
		}

		session := models.NewSession(auth.User())
		if err := uc.SessionDao.Store(session); err.NotNil() {
			uc.Logger.Error(err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		uc.OutputPort.Write(session)
	})
}
