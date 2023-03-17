package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"inspur-tencentmeeting-sso/biz/model"

	"github.com/cloudwego/hertz/pkg/app"
)

// Ping .
// PingHandler 测试handler
// @Summary 测试Summary
// @Description 测试Description
// @Accept application/json
// @Produce application/json
// @Router /ping [get]
func Ping(ctx context.Context, c *app.RequestContext) {
	resp := new(model.Response)
	resp.Data = map[string]string{
		"message": "pong",
	}
	c.JSON(consts.StatusOK, resp)
}
