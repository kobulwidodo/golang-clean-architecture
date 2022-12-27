package auth

import (
	"context"
	"encoding/json"
	"go-clean/src/business/entity"
	"go-clean/src/lib/codes"
	"go-clean/src/lib/errors"
	"go-clean/src/lib/log"

	firebase "firebase.google.com/go"
	firebase_auth "firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type key string

const (
	userAuthInfo key = "userAuthInfo"
)

type Interface interface {
	RegisterUser(ctx context.Context, user entity.FirebaseUser) (entity.FirebaseUser, error)
	VerifyToken(ctx context.Context, bearertoken string) (*firebase_auth.Token, error)
	GetUser(ctx context.Context, uid string) (entity.FirebaseUser, error)
	SetUserAuthInfo(ctx context.Context, user entity.AuthUser, fbasetoken *firebase_auth.Token) context.Context
}

type auth struct {
	log      log.Interface
	firebase *firebase_auth.Client
}

type Config struct {
	Firebase FirebaseConf
}

type FirebaseConf struct {
	AccountKey FirebaseAccountKey
}

type FirebaseAccountKey struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderx509CertURL string `json:"auth_provider_x509_cert_url"`
	Clientx509CertURL       string `json:"client_x509_cert_url"`
}

func Init(cfg Config, log log.Interface) Interface {
	ctx := context.Background()

	accountKey, err := json.Marshal(cfg.Firebase.AccountKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsJSON(accountKey))
	if err != nil {
		log.Fatal(err.Error())
	}

	firebaseAuth, err := app.Auth(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &auth{
		log:      log,
		firebase: firebaseAuth,
	}
}

func (a *auth) RegisterUser(ctx context.Context, user entity.FirebaseUser) (entity.FirebaseUser, error) {
	params := (&firebase_auth.UserToCreate{}).
		UID(user.ID).
		Email(user.Email).
		Password(user.Password).
		DisplayName(user.DisplayName)

	u, err := a.firebase.CreateUser(ctx, params)
	if err != nil {
		return user, err
	}

	newUser := entity.FirebaseUser{
		ID:          u.UID,
		Email:       u.Email,
		DisplayName: u.DisplayName,
		PhotoURL:    u.PhotoURL,
		IsDisabled:  u.Disabled,
	}

	return newUser, nil
}

func (a *auth) VerifyToken(ctx context.Context, bearertoken string) (*firebase_auth.Token, error) {
	token, err := a.firebase.VerifyIDToken(ctx, bearertoken)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeUnauthorized, "token invalid")
	}
	return token, nil
}

func (a *auth) GetUser(ctx context.Context, uid string) (entity.FirebaseUser, error) {
	firebaseUser := entity.FirebaseUser{}

	fbUserRecord, err := a.firebase.GetUser(ctx, uid)
	if err != nil {
		return firebaseUser, err
	}

	firebaseUser.ID = fbUserRecord.UID
	firebaseUser.DisplayName = fbUserRecord.DisplayName
	firebaseUser.Email = fbUserRecord.Email
	firebaseUser.IsDisabled = fbUserRecord.Disabled
	firebaseUser.PhotoURL = fbUserRecord.PhotoURL

	return firebaseUser, nil
}

func (a *auth) SetUserAuthInfo(ctx context.Context, user entity.AuthUser, fbasetoken *firebase_auth.Token) context.Context {
	userAuth := entity.AuthUserInfo{
		AuthUser:      user,
		FirebaseToken: *fbasetoken,
	}

	return context.WithValue(ctx, userAuthInfo, userAuth)
}
