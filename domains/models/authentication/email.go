package authentication

import (
	"go-cleanarchitecture/domains/errors"
	"net/mail"
)

type Email struct {
	// 認証用のメールアドレスを表現する値オブジェクト

	value string
}

func NewEmail(value string) (Email, errors.Domain) {
	empty := Email{}

	email, err := mail.ParseAddress(value)
	if err != nil {
		return empty, errors.Invalid("Invalid email")
	} else {
		return Email{email.Address}, errors.None
	}
}

func (email Email) Value() string {
	return email.value
}
