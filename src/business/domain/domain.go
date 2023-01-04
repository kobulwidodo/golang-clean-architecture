package domain

import (
	"go-clean/src/business/domain/user"

	"gorm.io/gorm"
)

type Domains struct {
	User user.Interface
}

func Init(db *gorm.DB) *Domains {
	d := &Domains{
		User: user.Init(db),
	}

	return d
}
