package sso

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"inspur-tencentmeeting-sso/biz/model"
	"inspur-tencentmeeting-sso/biz/model/sso"
	"inspur-tencentmeeting-sso/constant"
)

// swagger 注释在：https://github.com/swaggo/swag/blob/master/README_zh-CN.md

// CreateIdToken 生成id_token
// @Summary 生成id_token
// @Description 生成id_token
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer"
// @Success 200 {object} model.OKResponse "正确返回"
// @Success 400 {object} model.TokenErrorResponse "参数（token错误）错误返回"
// @Success 500 {object} model.ServerErrorResponse "服务器错误返回"
// @Router /v1/sso/id_token [post]
func CreateIdToken(ctx context.Context, c *app.RequestContext) {
	var (
		resp      = new(model.Response)
		userId, _ = c.Get(constant.TokenKeyUserId)
	)

	hlog.Info("userId: ", userId)
	resp.Data = &sso.Response{
		IDToken: "",
	}
	c.JSON(consts.StatusOK, resp)
}
