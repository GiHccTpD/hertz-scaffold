package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"inspur-tencentmeeting-sso/constant"
	"strings"
)

func JwtAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//resp := new(model.Response)
		token := string(c.GetHeader("Authorization"))
		token = strings.Split(token, " ")[1]
		hlog.Info("token: ", token)

		//// check token
		//checkTokenResult, err := rpc.CheckToken(token)
		//if err != nil {
		//	hlog.Errorf("JwtAuth rpc.CheckToken failed, msg: %v", err)
		//	resp.Code = constant.ErrArg
		//	resp.Message = err.Error()
		//	c.JSON(consts.StatusInternalServerError, resp)
		//	c.Abort()
		//	return
		//}
		//hlog.Debugf("checkTokenResult: %+v", checkTokenResult)
		//
		//if checkTokenResult.Code != constant.Success {
		//	resp.Code = checkTokenResult.Code
		//	resp.Message = "token error"
		//	c.JSON(consts.StatusBadRequest, resp)
		//	c.Abort()
		//}

		//c.Set(constant.TokenKeyUserId, strings.ToLower(checkTokenResult.UserId))
		//c.Set(constant.ReqDeviceType, checkTokenResult.DeviceType)
		//c.Set(constant.TokenKeyOrgId, checkTokenResult.OrgId)
		//c.Set(constant.TokenKeyUserName, checkTokenResult.UserName)

		c.Set(constant.TokenKeyUserId, "")
		c.Set(constant.ReqDeviceType, "")
		c.Set(constant.TokenKeyOrgId, "")
		c.Set(constant.TokenKeyUserName, "")
		c.Next(ctx)
	}
}
