package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UID      string
	Name     string
	Email    string
	ImageUrl string
}

type CreateUserParams struct {
	Name     string `json:"name"`
	UID      string `json:"-"`
	Email    string `json:"email"`
	ImageUrl string `json:"imamge_url"`
}

func (u *User) ConvertToAuthUser() AuthUser {
	return AuthUser{
		ID:          u.ID,
		UID:         u.UID,
		Email:       u.Email,
		DisplayName: u.Name,
	}
}
