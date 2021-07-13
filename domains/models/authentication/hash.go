package authentication

import (
	"crypto/sha1"
	"fmt"
)

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
