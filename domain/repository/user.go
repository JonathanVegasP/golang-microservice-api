package repository

import "flutter-store-api/domain/entity"

type IUserRepository interface {
	Get(id *uint64) (*entity.User, error)
	GetByEmail(email *string) (*entity.User, error)
	GetAll() ([]entity.User, error)
	Create(user *entity.User) bool
	Save(user *entity.User) bool
	Delete(id *uint64) bool
}
