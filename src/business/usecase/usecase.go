package usecase

import (
	"go-clean/src/business/domain"
	"go-clean/src/business/usecase/user"
	"go-clean/src/lib/auth"
	"go-clean/src/lib/log"
)

type Usecase struct {
	User user.Interface
}

func Init(log log.Interface, auth auth.Interface, d *domain.Domains) *Usecase {
	uc := &Usecase{
		User: user.Init(log, auth, d.User),
	}

	return uc
}
