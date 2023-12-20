package handler

import (
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/url-shortener/src/business/entity"
	"github.com/gin-gonic/gin"
)

func (r *rest) GetListUser(ctx *gin.Context) {
	var param entity.UserParam
	if err := r.BindParams(ctx, &param); err != nil {
		r.httpRespError(ctx, err)
		return
	}

	user, pg, err := r.uc.User.GetListAsAdmin(ctx.Request.Context(), param)
	if err != nil {
		r.httpRespError(ctx, err)
		return
	}

	r.httpRespSuccess(ctx, codes.CodeSuccess, user, pg)
}
