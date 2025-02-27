package middleware

import (
	errors2 "gin-blog/pkg/errors"
	"gin-blog/pkg/globals"
	"gin-blog/pkg/response"
	"gin-blog/pkg/token"
	"github.com/gin-gonic/gin"
	"log"
)

func AuthMiddleware() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(context *gin.Context) {
		tokenString, err := context.Cookie("token")

		if err != nil || tokenString == "" {
			response.HTTPFail(errors2.ErrPermissions, context)
			log.Println("AuthMiddleware() err=", err)
			context.Abort()
			return
		}

		tokenString = tokenString[7:]
		err = token.ParseToken(tokenString, context, appCtx.GetDb())
		if err != nil {
			response.HTTPFail(errors2.ErrPermissions, context)
			log.Println("AuthMiddleware() err=", err)
			context.Abort()
			return
		}

		context.Next()
	}
}
