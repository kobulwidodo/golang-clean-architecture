package domain

import (
	"go-clean/src/business/domain/user"
	"go-clean/src/lib/auth"
	"go-clean/src/lib/log"

	"gorm.io/gorm"
)

type Domains struct {
	User user.Interface
}

func Init(log log.Interface, db *gorm.DB, auth auth.Interface) *Domains {
	d := &Domains{
		User: user.Init(log, db, auth),
	}

	return d
}
