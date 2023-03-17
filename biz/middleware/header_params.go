package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"inspur-tencentmeeting-sso/biz/model"
	"inspur-tencentmeeting-sso/constant"
)

// ReqHeaderParamValidate 请求头 Authorization 字段校验
func ReqHeaderParamValidate() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		resp := new(model.Response)
		// pre-handle
		hlog.Info("JwtAuth")

		// 参数绑定需要配合特定的go tag使用
		// https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/binding-and-validate/
		// https://github.com/bytedance/go-tagexpr
		type Header struct {
			Authorization string `header:"Authorization,required"`
		}

		// BindAndValidate
		var req Header
		err := c.BindAndValidate(&req)
		if err != nil {
			resp.Code = constant.ErrArg
			resp.Message = err.Error()
			hlog.Errorf("ReqHeaderParamValidate BindAndValidate failed, msg: %v", err)
			c.JSON(consts.StatusInternalServerError, resp)
			c.Abort()
			return
		}

		// Bind
		req = Header{}
		err = c.Bind(&req)
		if err != nil {
			resp.Code = constant.ErrArg
			resp.Message = err.Error()
			hlog.Errorf("ReqHeaderParamValidate Bind failed, msg: %v", err)
			c.JSON(consts.StatusInternalServerError, resp)
			c.Abort()
			return
		}

		// Validate，需要使用 "vd" tag
		err = c.Validate(&req)
		if err != nil {
			resp.Code = constant.ErrToken
			resp.Message = "参数错误"
			hlog.Errorf("ReqHeaderParamValidate Validate failed, msg: %v", err)
			c.JSON(consts.StatusBadRequest, resp)
			c.Abort()
			return
		}

		c.Next(ctx)
	}
}
