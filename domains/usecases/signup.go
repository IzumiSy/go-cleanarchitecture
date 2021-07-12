package usecases

import "go-cleanarchitecture/domains"

type SignupOutputPort interface {
	domains.OutputPort
}

type SignupParam struct {
	Email    string
	Password string
}

type signupUsecase struct {
	outputPort SignupOutputPort
	logger     domains.Logger
}

func (usecase signupUsecase) Execute(params SignupParam) {
	// todo
}
