package rest

import (
	"errors"
	"fmt"
	"go-clean/src/business/entity"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (r *rest) httpRespSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	resp := entity.Response{
		Meta: entity.Meta{
			Message: message,
			Code:    code,
			IsError: false,
		},
		Data: data,
	}
	ctx.JSON(code, resp)
}

func (r *rest) httpRespError(ctx *gin.Context, code int, err error) {
	resp := entity.Response{
		Meta: entity.Meta{
			Message: err.Error(),
			Code:    code,
			IsError: true,
		},
		Data: nil,
	}
	ctx.AbortWithStatusJSON(code, resp)
}

func (r *rest) VerifyUser(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		r.httpRespError(ctx, http.StatusUnauthorized, errors.New("empty token"))
		return
	}

	var tokenString string
	_, err := fmt.Sscanf(authHeader, "Bearer %v", &tokenString)
	if err != nil {
		r.httpRespError(ctx, http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	token, err := r.ValidateToken(tokenString)
	if err != nil {
		r.httpRespError(ctx, http.StatusUnauthorized, err)
		return
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		r.httpRespError(ctx, http.StatusUnauthorized, errors.New("failed to claim token"))
		return
	}

	user := entity.User{}
	user, err = r.uc.User.GetById(uint(claim["id"].(float64)))
	if err != nil {
		r.httpRespError(ctx, http.StatusUnauthorized, errors.New("error while getting user"))
		return
	}

	c := ctx.Request.Context()
	c = r.auth.SetUserAuthInfo(c, user.ConvertToAuthUser(), tokenString)
	ctx.Request = ctx.Request.WithContext(c)

	ctx.Next()
}

func (r *rest) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("token invalid")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
