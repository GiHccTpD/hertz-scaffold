package biz

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/swagger"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	"inspur-tencentmeeting-sso/biz/handler"
	"inspur-tencentmeeting-sso/biz/handler/sso"
	mw "inspur-tencentmeeting-sso/biz/middleware"
	_ "inspur-tencentmeeting-sso/docs"
)

// route https://github.com/cloudwego/hertz-examples/blob/9b0d1623111c66e8bbf472b11e0ed73c83894851/route/main.go

// RegisterGroupRoute group route
func RegisterGroupRoute(h *server.Hertz) {
	port := viper.GetString("server.port")

	// ping pong
	h.GET("/ping", handler.Ping)

	// swagger https://github.com/hertz-contrib/swagger https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/middleware/swagger/
	url := swagger.URL("http://localhost:" + port + "/swagger/doc.json") // The url pointing to API definition
	h.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler, url))
	fmt.Println("swagger url is: http://localhost:" + port + "/swagger/index.html")

	// sso
	v1 := h.Group("/v1/sso", ssoGroupMiddleware()...)
	{
		v1.POST("/id_token", sso.CreateIdToken)
	}
}

func ssoGroupMiddleware() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.ReqHeaderParamValidate(),
		mw.JwtAuth(),
	}
}
