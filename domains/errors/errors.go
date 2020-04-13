package errors

import (
	"golang.org/x/xerrors"
)

type Domain struct {
	// ドメイン層におけるエラーを表現する型
	// 基本的にはバリデーションエラーを実装するのに使う

	err error
}

var (
	None = Domain{}
)

// バリデーションエラーのメッセージをつくる関数
func Invalid(reason string) Domain {
	return Domain{
		err: xerrors.Errorf("Domain error: %w", xerrors.New(reason)),
	}
}

// 永続化装置などアプリケーション外部のエラーをラップするための関数
func External(err error) Domain {
	if err == nil {
		return None
	}

	return Domain{
		err: xerrors.Errorf("External error: %w", err),
	}
}

func (e Domain) NotNil() bool {
	return !e.Is(None)
}

func (e Domain) Is(other Domain) bool {
	return xerrors.Is(e.err, other.err)
}

func (e Domain) Value() error {
	return e.err
}
