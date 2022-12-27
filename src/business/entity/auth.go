package entity

import (
	firebase_auth "firebase.google.com/go/auth"
)

type FirebaseUserParam struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type FirebaseUser struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DisplayName string `json:"display_name"`
	PhotoURL    string `json:"photo_url"`
	IsDisabled  bool   `json:"is_disabled"`
}

type AuthUserInfo struct {
	AuthUser      AuthUser            `json:"auth_user"`
	FirebaseToken firebase_auth.Token `json:"firebase_token"`
}

type AuthUser struct {
	ID          uint   `json:"id"`
	UID         string `json:"uid"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}
