package usecase

import (
	"go-clean/src/business/domain"
	"go-clean/src/business/usecase/user"
	"go-clean/src/lib/auth"
)

type Usecase struct {
	User user.Interface
}

func Init(auth auth.Interface, d *domain.Domains) *Usecase {
	uc := &Usecase{
		User: user.Init(d.User, auth),
	}

	return uc
}
