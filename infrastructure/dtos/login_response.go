package dtos

import "flutter-store-api/infrastructure/auth"

type LoginResponse struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

func (l *LoginResponse) CreateToken(id interface{}) {
	l.Token = auth.CreateJWT(id, nil)
}
