package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
	"time"
)

// https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/middleware/cors/

func CORS() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Content-Type", "api_key", "Authorization"},
			//AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			//AllowOriginFunc: func(origin string) bool {
			//	return origin == "*"
			//},
			MaxAge: 12 * time.Hour,
		})
	}
}
