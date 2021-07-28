package domains

import (
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models"
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

type Authorizer interface {
	Run() (models.Session, error)
}

func (uc AuthorizedUsecase) Run(authorizer Authorizer) {
	var (
		INVALID_AUTHORIZATION = errors.Preconditional("Invalid authorization")
	)

	session, err := authorizer.Run()
	if err != nil {
		uc.outputPort.Raise(INVALID_AUTHORIZATION)
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
