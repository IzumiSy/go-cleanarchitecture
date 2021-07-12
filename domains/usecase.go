package domains

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models/session"
	"go-cleanarchitecture/domains/models/user"
)

type OutputPort interface {
	Raise(err errors.Domain)
}

// 認証済みユースケース基盤

type AuthorizedUsecase struct {
	usecase    func(currentUserID user.ID)
	outputPort OutputPort
}

func (uc AuthorizedUsecase) Run(sessionDao SessionRepository, sessionID string) {
	var (
		INVALID_SESSION_ID = errors.Invalid("Invalid session ID")
	)

	_sessionID, err := session.NewID(sessionID)
	if err.NotNil() {
		uc.outputPort.Raise(INVALID_SESSION_ID)
		return
	}

	session, err, exists := sessionDao.Get(_sessionID)
	if err.NotNil() {
		uc.outputPort.Raise(INVALID_SESSION_ID)
		return
	}

	if !exists {
		uc.outputPort.Raise(INVALID_SESSION_ID)
		return
	}

	uc.usecase(session.UserID())
}

func NewAuthorizedUsecase(outputPort OutputPort, usecase func(currentUserID user.ID)) AuthorizedUsecase {
	return AuthorizedUsecase{
		usecase:    usecase,
		outputPort: outputPort,
	}
}

// 未認証ユースケース基盤

type UnauthorizedUsecase struct {
	usecase func()
}

func (uc UnauthorizedUsecase) Run() {
	uc.usecase()
}

func NewUnauthorizedUsecase(usecase func()) UnauthorizedUsecase {
	return UnauthorizedUsecase{usecase}
}
