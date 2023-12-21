package handler

import (
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/url-shortener/src/business/entity"
	"github.com/gin-gonic/gin"
)

// @Summary Get User List
// @Description Get list all user
// @Security BearerAuth
// @Tags Admin
// @Param limit query integer false "limit"
// @Param page query integer false "page"
// @Param disableLimit query boolean false "disable limit" Enums(true, false)
// @Produce json
// @Success 200 {object} entity.HTTPResp{data=[]entity.User{}}
// @Failure 500 {object} entity.HTTPResp{}
// @Router /v1/admin/user [GET]
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
