package dtos

import (
	"flutter-store-api/infrastructure/security"
)

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l *Login) HashPassword() {
	l.Password = security.CreateHash([]byte(l.Email), []byte(l.Password))
}
