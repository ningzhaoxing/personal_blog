package routers

import (
	"gin-blog/internal/modules/user/handlers"
	"gin-blog/internal/server/middleware"
	"gin-blog/pkg/response"
	"github.com/gin-gonic/gin"
)

func userRouter(e *gin.Engine) {
	e.GET("/login", func(context *gin.Context) {
		response.Success("", "login", response.Data{}, context)
	})
	e.GET("/register", func(context *gin.Context) {
		response.Success("", "register", response.Data{}, context)
	})

	auth := e.Group("/auth", middleware.RateLimitMiddleware(1, 1))
	auth.POST("/login", handlers.Login())
	auth.POST("/register", handlers.Register())

	e.POST("/headPhoto/update", middleware.AuthMiddleware(), middleware.IsManagerMiddleware(), handlers.UpdateUserHeadPhoto())
}
