package rest

import (
	"go-clean/src/business/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *rest) RegisterUser(ctx *gin.Context) {
	var userParam entity.CreateUserParam
	if err := ctx.ShouldBindJSON(&userParam); err != nil {
		r.httpRespError(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := r.uc.User.Create(userParam)
	if err != nil {
		r.httpRespError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.httpRespSuccess(ctx, http.StatusCreated, "successfully registered new user", user)
}

func (r *rest) LoginUser(ctx *gin.Context) {
	var userParam entity.LoginUserParam
	if err := ctx.ShouldBindJSON(&userParam); err != nil {
		r.httpRespError(ctx, http.StatusBadRequest, err)
		return
	}

	token, err := r.uc.User.Login(userParam)
	if err != nil {
		r.httpRespError(ctx, http.StatusInternalServerError, err)
		return
	}

	r.httpRespSuccess(ctx, http.StatusOK, "successfully login", gin.H{"token": token})
}
