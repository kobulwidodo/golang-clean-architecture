package user

import (
	"go-clean/src/business/entity"

	"gorm.io/gorm"
)

type Interface interface {
	Create(user entity.User) (entity.User, error)
	GetByUsername(username string) (entity.User, error)
	GetById(id uint) (entity.User, error)
}

type user struct {
	db *gorm.DB
}

func Init(db *gorm.DB) Interface {
	a := &user{
		db: db,
	}

	return a
}

func (a *user) Create(user entity.User) (entity.User, error) {
	if err := a.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (a *user) GetByUsername(username string) (entity.User, error) {
	user := entity.User{}

	if err := a.db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (a *user) GetById(id uint) (entity.User, error) {
	user := entity.User{}

	if err := a.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
