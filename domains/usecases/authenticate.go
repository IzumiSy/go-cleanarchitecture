package usecases

import (
	"context"
	"time"

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
	Ctx               context.Context
	OutputPort        AuthenticateOutputPort
	AuthenticationDao domains.AuthenticationRepository
	SessionDao        domains.SessionRepository
	Logger            domains.Logger
	Publisher         domains.EventPublisher
}

func (uc AuthenticateUsecase) Build(params AuthenticateParam) domains.UnauthorizedUsecase {
	return domains.NewUnauthorizedUsecase(func() {
		// [ユーザーの認証を行うユースケース]
		// "ログイン"でも命名はよかったが、今後外部APIとして認証を実装したりする可能性を考えると
		// 人間以外のアクタも考慮し抽象化して"認証"と表現したくなったのでこの命名としている。

		var (
			USER_NOT_FOUND   = errors.Preconditional("User not found")
			INVALID_PASSWORD = errors.Preconditional("Invalid password")
		)

		email, err := authentication.NewEmail(params.Email)
		if err.NotNil() {
			uc.Logger.Warnf(uc.Ctx, err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		auth, err, exists := uc.AuthenticationDao.GetByEmail(email)
		if err.NotNil() {
			uc.Logger.Errorf(uc.Ctx, err.Error())
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
			uc.Logger.Errorf(uc.Ctx, err.Error())
			uc.OutputPort.Raise(err)
			return
		}
		uc.Logger.Info(fmt.Sprintf("New session stored: %s", session.ID()))

		event := UserAuthenticatedEvent{
			UserID:    auth.User().ID().String(),
			CreatedAt: time.Now(),
		}

		// Eventのpublishに失敗しても意図的にエラーはレスポンスせずErrorのレポートのみとしておく
		// 非同期処理のエラーは別途手動で復旧作業を行う
		if err := uc.Publisher.Publish(event); err.NotNil() {
			uc.Logger.Errorf(uc.Ctx, "Failed publishing event: %s", err.Error())
		} else {
			uc.Logger.Infof(uc.Ctx, "Event published: %s", event.ID())
		}

		uc.OutputPort.Write(session)
	})
}

type UserAuthenticatedEvent struct {
	UserID    string
	CreatedAt time.Time
}

func (UserAuthenticatedEvent) Name() domains.EventName {
	return domains.UserAuthenticated
}

func (UserAuthenticatedEvent) ID() domains.EventID {
	return domains.NewEventID()
}
