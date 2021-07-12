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
}
