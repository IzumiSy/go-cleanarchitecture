package usecases

import (
	"context"
	"go-cleanarchitecture/domains"
	"go-cleanarchitecture/domains/models"
	"go-cleanarchitecture/domains/models/user"
)

type GetTodosOutputPort interface {
	domains.OutputPort
	Write(todos models.Todos)
}

type GetTodosUsecase struct {
	Ctx        context.Context
	OutputPort GetTodosOutputPort
	TodosDao   domains.TodosRepository
	Logger     domains.Logger
}

func (uc GetTodosUsecase) Build() domains.AuthorizedUsecase {
	return domains.NewAuthorizedUsecase(uc.OutputPort, func(currentUserID user.ID) {
		todos, err := uc.TodosDao.GetByUserID(currentUserID)
		if err.NotNil() {
			uc.Logger.Errorf(uc.Ctx, err.Error())
			uc.OutputPort.Raise(err)
			return
		}

		uc.OutputPort.Write(todos)
	})
}
