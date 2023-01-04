package auth

import (
	"context"
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type contextKey string

const (
	userAuthInfo contextKey = "UserAuthInfo"
)

type Interface interface {
	SetUserAuthInfo(ctx context.Context, user User, token string) context.Context
	GetUserAuthInfo(ctx context.Context) (UserAuthInfo, error)
	GenerateToken(user User) (string, error)
	GenerateGuestToken() (string, error)
}

type auth struct {
}

func Init() Interface {
	return &auth{}
}

func (a *auth) SetUserAuthInfo(ctx context.Context, user User, token string) context.Context {
	userAuth := UserAuthInfo{
		User:  user,
		Token: token,
	}

	return context.WithValue(ctx, userAuthInfo, userAuth)
}

func (a *auth) GetUserAuthInfo(ctx context.Context) (UserAuthInfo, error) {
	userContext := ctx.Value(userAuthInfo)
	user, ok := userContext.(UserAuthInfo)
	if !ok {
		return user, errors.New("failed to get user auth")
	}

	return user, nil
}

func (a *auth) GenerateToken(user User) (string, error) {
	claim := jwt.MapClaims{}
	claim["id"] = user.ID
	claim["is_admin"] = user.IsAdmin
	claim["is_guest"] = false

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (a *auth) GenerateGuestToken() (string, error) {
	claim := jwt.MapClaims{}
	guestId, _ := gonanoid.New()
	claim["guest_id"] = guestId
	claim["is_guest"] = true

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
