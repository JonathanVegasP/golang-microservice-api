package services

import (
	"flutter-store-api/domain/entity"
	"flutter-store-api/infrastructure/dtos"
)

type IUserController interface {
	GetUser(id *uint64) *dtos.CreateUserResponse
	GetUserByLogin(login *dtos.Login) *dtos.LoginResponse
	GetAllUsers() []entity.User
	CreateUser(user *dtos.CreateUser) *dtos.CreateUserResponse
	UpdateUser(user *dtos.UpdateUser, id *uint64) bool
	DeleteUser(id *uint64) bool
}
