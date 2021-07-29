package errors

import (
	"fmt"

	"golang.org/x/xerrors"
)

var _ xerrors.Wrapper = Domain{}

type Domain struct {
	// ドメイン層におけるエラーを表現する型

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

// [事前条件違反]
// 主にバリデーションエラーなどユースケースの実行条件を満たさないことを示す
func Preconditional(reason string) Domain {
	return Domain{
		type_:  PreconditionalError,
		reason: reason,
	}
}

// [事後条件違反]
// ユースケースの実行は開始したが外部装置などの不具合でユースケース実行が行えないことを示す
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

func (e Domain) Is(other error) bool {
	return xerrors.Is(e, other)
}

func (e Domain) IsType(type_ ErrorType) bool {
	return e.type_ == type_
}
