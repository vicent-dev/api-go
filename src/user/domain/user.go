package domain

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Email    string
	Password string
}

func (u *User) HashPassword() (string, error) {

	hPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(hPassword), nil
}

func (u *User) VerifyPassword(plainPassword string) bool {
	return nil == bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))
}
