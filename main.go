package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/spf13/viper"
	r "inspur-tencentmeeting-sso/biz"
	_ "inspur-tencentmeeting-sso/biz/pkg/lookup"
	"inspur-tencentmeeting-sso/constant"
	"inspur-tencentmeeting-sso/pkg/cfg"
	_ "inspur-tencentmeeting-sso/pkg/hlog"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// @title INSPUR-TENCENTMEETING-SSO
// @version 1.0
// @description 浪潮集团腾讯会议sso免登录获取id_token服务.

// @contact.name wangxiao05@inspur.com
// @contact.url http://10.110.55.12:8888/wangxiao
func main() {
	cfg.Init()
	fmt.Printf("%v is running...\n", fmt.Sprintf("\x1b[32m%s\x1b[0m", constant.ServerName))
	port := viper.GetString("server.port")

	h := server.Default(server.WithHostPorts(":" + port))
	h.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "api_key", "Authorization"},
		//AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	h.Use(recovery.Recovery(recovery.WithRecoveryHandler(myRecoveryHandler)))

	// register route group
	r.RegisterGroupRoute(h)

	//h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
	//	<-ctx.Done()
	//	fmt.Printf("%v\n", fmt.Sprintf("\x1b[32m%s\x1b[0m", constant.ByeBye))
	//})

	h.Spin()
}

func myRecoveryHandler(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte) {
	hlog.SystemLogger().CtxErrorf(c, "[Recovery] err=%v\nstack=%s", err, stack)
	hlog.SystemLogger().Infof("Client: %s", ctx.Request.Header.UserAgent())
	ctx.AbortWithStatus(consts.StatusInternalServerError)
}
