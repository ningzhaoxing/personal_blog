package middleware

import (
	errors2 "gin-blog/pkg/errors"
	"gin-blog/pkg/response"
	"gin-blog/pkg/token"
	"github.com/gin-gonic/gin"
)

func IsManagerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, err := token.NewToken(context).GetUser()
		if err != nil {
			response.HTTPFail(errors2.ErrPermissions, context)
			context.Abort()
			return
		}

		if !user.IsOwner {
			response.HTTPFail(errors2.ErrPermissions, context)
			context.Abort()
			return
		}
		context.Next()
	}
}
