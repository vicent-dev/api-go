package auth

import (
	user "api/src/user/domain"
	"fmt"
)

type jwt string

type headerJWT map[string]string

var header = headerJWT{
	"alg": "HS256",
	"typ": "JWT",
}

func NewJWT(u *user.User) jwt {
	fmt.Printf("%v", header)
	return jwt("aaa")
}

func (jwt *jwt) GetClaims() (*user.User, error) {
	return nil, nil
}
