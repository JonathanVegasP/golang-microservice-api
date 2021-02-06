package dtos

import (
	"flutter-store-api/infrastructure/security"
)

type CreateUser struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (u *CreateUser) HashPassword() {
	u.Password = security.CreateHash([]byte(u.Email), []byte(u.Password))
}
