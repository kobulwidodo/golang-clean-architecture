package user

import (
	"context"
	userDom "go-clean/src/business/domain/user"
	"go-clean/src/business/entity"
	"go-clean/src/lib/auth"
	"go-clean/src/lib/log"
)

type Interface interface {
	Create(ctx context.Context, params entity.CreateUserParams) (entity.User, error)
	GetByUID(ctx context.Context, uid string) (entity.User, error)
}

type user struct {
	log  log.Interface
	auth auth.Interface
	user userDom.Interface
}

func Init(log log.Interface, auth auth.Interface, ud userDom.Interface) Interface {
	u := &user{
		log:  log,
		auth: auth,
		user: ud,
	}

	return u
}

func (u *user) Create(ctx context.Context, params entity.CreateUserParams) (entity.User, error) {
	user := entity.User{
		UID:      params.UID,
		Name:     params.Name,
		Email:    params.Email,
		ImageUrl: params.ImageUrl,
	}

	newUser, err := u.user.CreateUser(ctx, user)
	if err != nil {
		return user, err
	}

	return newUser, nil
}

func (u *user) GetByUID(ctx context.Context, uid string) (entity.User, error) {
	return u.user.GetByUID(ctx, uid)
}
