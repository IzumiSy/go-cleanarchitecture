package errors

import (
	"fmt"

	"golang.org/x/xerrors"
)

var _ xerrors.Wrapper = Domain{}

type Domain struct {
	// ドメイン層におけるエラーを表現する型
	// 基本的にはバリデーションエラーを実装するのに使う

	type_  ErrorType
	reason string
}

func (e Domain) Unwrap() error {
	return xerrors.New(e.type_.name)
}

func (e Domain) Error() string {
	return fmt.Sprintf("%s: %s", e.type_.name, e.reason)
}

type ErrorType struct {
	name string
}

var (
	None = Domain{}

	PreconditionalError  ErrorType = ErrorType{"preconditional error"}
	PostconditionalError ErrorType = ErrorType{"postconditional error"}
)

func Preconditional(reason string) Domain {
	return Domain{
		type_:  PreconditionalError,
		reason: reason,
	}
}

func Postconditional(err error) Domain {
	if err == nil {
		return None
	}

	return Domain{
		type_:  PostconditionalError,
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
	return e.type_ == type_
}
