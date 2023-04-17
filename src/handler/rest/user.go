package rest

import (
	"go-clean/src/business/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Register User
// @Description Register New User
// @Tags Auth
// @Param user body entity.CreateUserParam true "user info"
// @Produce json
// @Success 200 {object} entity.Response{data=entity.User{}}
// @Failure 400 {object} entity.Response{}
// @Failure 401 {object} entity.Response{}
// @Failure 404 {object} entity.Response{}
// @Failure 500 {object} entity.Response{}
// @Router /api/v1/auth/register [POST]
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

// @Summary Login User
// @Description Login User
// @Tags Auth
// @Param user body entity.LoginUserParam true "user info"
// @Produce json
// @Success 200 {object} entity.Response{}
// @Failure 400 {object} entity.Response{}
// @Failure 401 {object} entity.Response{}
// @Failure 404 {object} entity.Response{}
// @Failure 500 {object} entity.Response{}
// @Router /api/v1/auth/login [POST]
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
