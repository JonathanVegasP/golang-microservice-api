package dtos

import "flutter-store-api/infrastructure/security"

type UpdateUserPassword struct {
	Password string `json:"password" binding:"required,min=8"`
}

func (u *UpdateUserPassword) HashPassword(email []byte) {
	u.Password = security.CreateHash(email, []byte(u.Password))
}
