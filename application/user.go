package application

import (
	"flutter-store-api/domain/entity"
	"flutter-store-api/domain/repository"
	"flutter-store-api/domain/services"
	"flutter-store-api/infrastructure/dtos"
	"flutter-store-api/infrastructure/mapper"
)

func NewUserService(repo repository.IUserRepository) services.IUserController {
	return &controller{
		repo: repo,
	}
}

type controller struct {
	repo repository.IUserRepository
}

func (u *controller) GetUserByLogin(login *dtos.Login) *dtos.LoginResponse {
	user, err := u.repo.GetByEmail(&login.Email)

	if err != nil || user == nil {
		return nil
	}

	login.HashPassword()

	if user.Password != login.Password {
		return nil
	}

	response := &dtos.LoginResponse{
		Name: user.Name,
	}

	response.CreateToken(user.ID)

	return response
}

func (u *controller) GetUser(id *uint64) *dtos.CreateUserResponse {
	entity, err := u.repo.Get(id)

	if err != nil {
		return nil
	}

	var user dtos.CreateUserResponse

	mapper.AutoMapper(entity, &user)

	return &user
}

func (u *controller) GetAllUsers() []entity.User {
	user, err := u.repo.GetAll()

	if err != nil {
		return nil
	}

	return user
}

func (u *controller) CreateUser(user *dtos.CreateUser) *dtos.CreateUserResponse {
	user.HashPassword()

	var entity entity.User

	mapper.AutoMapper(user, &entity)

	if result := u.repo.Create(&entity); result {
		var dto dtos.CreateUserResponse
		mapper.AutoMapper(&entity, &dto)
		return &dto
	}

	return nil
}

func (u *controller) UpdateUser(user *dtos.UpdateUser, id *uint64) bool {
	var entity entity.User

	mapper.AutoMapper(user, &entity)

	(&entity).ID = *id

	return u.repo.Save(&entity)
}

func (u *controller) DeleteUser(id *uint64) bool {
	return u.repo.Delete(id)
}
