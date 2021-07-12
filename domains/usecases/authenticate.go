package usecases

import "go-cleanarchitecture/domains"

type AuthenticateOutputPort interface {
	domains.OutputPort
	Write()
}

type AuthenticateParam struct {
	Email    string
	Password string
}

type authenticateUsecase struct {
	outputPort AuthenticateOutputPort
	logger     domains.Logger
}

func (usecase authenticateUsecase) Execute(params AuthenticateParam) {
	// todo
}
