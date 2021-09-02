package authentication

import (
	"crypto/sha1"
	"fmt"
	"go-cleanarchitecture/domains/errors"
	"go-cleanarchitecture/domains/models/entity"
	"net/mail"
)

type CreatedAt struct {
	// 認証情報の作成日時を表す値オブジェクト
	entity.Time_
}

type Hash struct {
	// 認証用のパスワードハッシュを表現する値オブジェクト

	value string
}

func GenerateHash(value string) Hash {
	h := sha1.New()
	h.Write([]byte(value))
	return Hash{fmt.Sprintf("%x", h.Sum(nil))}
}

func NewHash(value string) Hash {
	return Hash{value}
}

func (hash Hash) Value() string {
	return hash.value
}

type Email struct {
	// 認証用のメールアドレスを表現する値オブジェクト

	value string
}

func NewEmail(value string) (Email, errors.Domain) {
	empty := Email{}

	email, err := mail.ParseAddress(value)
	if err != nil {
		return empty, errors.Preconditional("Invalid email")
	} else {
		return Email{email.Address}, errors.None
	}
}

func (email Email) Value() string {
	return email.value
}
