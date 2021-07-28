package errors

import (
	"fmt"

	"golang.org/x/xerrors"
)

var _ xerrors.Wrapper = Domain{}

type Domain struct {
	// ドメイン層におけるエラーを表現する型
	// 基本的にはバリデーションエラーを実装するのに使う

	err    ErrorType
	reason string
}

func (e Domain) Unwrap() error {
	return e.err
}

func (e Domain) Error() string {
	return fmt.Sprintf("%s: %s", e.err.Error(), e.reason)
}

type ErrorType error

var (
	None = Domain{}

	PreconditionalError  ErrorType = xerrors.New("preconditional error")
	PostconditionalError ErrorType = xerrors.New("postconditional error")
)

func Preconditional(reason string) Domain {
	return Domain{
		err:    PreconditionalError,
		reason: reason,
	}
}

func Postconditional(err error) Domain {
	if err == nil {
		return None
	}

	return Domain{
		err:    PostconditionalError,
		reason: err.Error(),
	}
}

func (e Domain) NotNil() bool {
	return !e.Is(None)
}

func (e Domain) Is(other Domain) bool {
	return xerrors.Is(e, other)
}

func (e Domain) IsType(type_ ErrorType) bool {
	return xerrors.Is(e.err, type_)
}
