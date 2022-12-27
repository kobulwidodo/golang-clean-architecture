package rest

import (
	"errors"
	"fmt"
	"go-clean/src/business/entity"
	"go-clean/src/lib/codes"
	customErrors "go-clean/src/lib/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *rest) BodyLogger(ctx *gin.Context) {
	if r.conf.LogRequest {
		r.log.Info(fmt.Sprintf(infoRequest, ctx.Request.RequestURI, ctx.Request.Method))
	}

	ctx.Next()
	if r.conf.LogResponse {
		if ctx.Writer.Status() < 300 {
			r.log.Info(
				fmt.Sprintf(infoResponse, ctx.Request.RequestURI, ctx.Request.Method, ctx.Writer.Status()))
		} else {
			r.log.Error(
				fmt.Sprintf(infoResponse, ctx.Request.RequestURI, ctx.Request.Method, ctx.Writer.Status()))
		}
	}
}

func (r *rest) VerifyToken(ctx *gin.Context) {
	token := ctx.Request.Header.Get("authorization")
	if token == "" {
		r.httpRespError(ctx, customErrors.NewWithCode(codes.CodeUnauthorized, "empty token"))
		return
	}

	var tokenID string
	_, err := fmt.Sscanf(token, "Bearer %v", &tokenID)
	if err != nil {
		r.httpRespError(ctx, customErrors.NewWithCode(codes.CodeUnauthorized, "invalid token"))
		return
	}

	fbaseToken, err := r.auth.VerifyToken(ctx.Request.Context(), tokenID)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	fbaseUser, err := r.auth.GetUser(ctx.Request.Context(), fbaseToken.UID)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	user, err := r.uc.User.GetByUID(ctx, fbaseToken.UID)
	var ce customErrors.CustomError
	if errors.As(err, &ce) {
		if ce.Code == codes.CodeSQLRecordDoesNotExist {
			params := entity.CreateUserParams{
				Name:     fbaseUser.DisplayName,
				UID:      fbaseUser.ID,
				Email:    fbaseUser.Email,
				ImageUrl: fbaseUser.PhotoURL,
			}
			_, err := r.uc.User.Create(ctx, params)
			if err != nil {
				r.httpRespError(ctx, err)
				return
			}
		} else {
			r.httpRespError(ctx, err)
			return
		}
	}

	c := ctx.Request.Context()
	c = r.auth.SetUserAuthInfo(c, user.ConvertToAuthUser(), fbaseToken)
	ctx.Request = ctx.Request.WithContext(c)

	ctx.Next()
}

func (r *rest) httpRespSuccess(ctx *gin.Context, code codes.Message, data interface{}) {
	resp := entity.Response{
		Meta: entity.Meta{
			Message: code.Message,
			Code:    code.HttpCode,
		},
		Data: data,
	}

	ctx.JSON(code.HttpCode, resp)
}

func (r *rest) httpRespError(ctx *gin.Context, err error) {
	var ce customErrors.CustomError
	if errors.As(err, &ce) {
		resp := entity.Response{
			Meta: entity.Meta{
				Message: ce.Cause,
				Code:    ce.Code.HttpCode,
			},
			Data: nil,
		}
		ctx.JSON(ce.Code.HttpCode, resp)
	}
	ctx.JSON(http.StatusInternalServerError, err.Error())
}
