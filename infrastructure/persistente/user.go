package persistente

import (
	"flutter-store-api/domain/entity"
	"flutter-store-api/domain/repository"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) repository.IUserRepository {
	return &userRepo{
		db: db,
	}
}

type userRepo struct {
	db *gorm.DB
}

func (u *userRepo) GetByEmail(email *string) (*entity.User, error) {
	var user entity.User

	if err := u.db.Find(&user, entity.User{Email: *email}).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) Get(id *uint64) (*entity.User, error) {
	var user entity.User

	if err := u.db.Find(&user, entity.User{ID: *id}).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) GetAll() ([]entity.User, error) {
	var users []entity.User

	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepo) Create(user *entity.User) bool {
	err := u.db.Create(user).Error

	return err == nil
}

func (u *userRepo) Save(user *entity.User) bool {
	err := u.db.Model(user).Updates(user).Error

	return err == nil
}

func (u *userRepo) Delete(id *uint64) bool {
	err := u.db.Delete(&entity.User{ID: *id}).Error

	return err == nil
}
