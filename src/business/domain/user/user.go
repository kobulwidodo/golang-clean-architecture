package user

import (
	"context"
	"errors"
	"go-clean/src/business/entity"
	"go-clean/src/lib/auth"
	"go-clean/src/lib/codes"
	customError "go-clean/src/lib/errors"
	"go-clean/src/lib/log"

	"gorm.io/gorm"
)

type Interface interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	GetByUID(ctx context.Context, uid string) (entity.User, error)
}

type user struct {
	log  log.Interface
	db   *gorm.DB
	auth auth.Interface
}

func Init(log log.Interface, db *gorm.DB, auth auth.Interface) Interface {
	u := &user{
		log:  log,
		db:   db,
		auth: auth,
	}

	return u
}

func (u *user) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	tx := u.db.Begin()
	defer tx.Rollback()

	if err := tx.Create(&user).Error; err != nil {
		return user, err
	}

	if err := tx.Commit().Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *user) GetByUID(ctx context.Context, uid string) (entity.User, error) {
	var user entity.User
	res := u.db.Where("uid = ?", uid).Take(&user)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return user, customError.NewWithCode(codes.CodeSQLRecordDoesNotExist, res.Error.Error())
	} else if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return user, customError.NewWithCode(codes.CodeSQLRead, res.Error.Error())
	}

	return user, nil
}
