package routers

import (
	"gin-blog/internal/modules/user/handlers"
	"gin-blog/internal/server/middleware"
	"github.com/gin-gonic/gin"
)

func commentRouter(e *gin.Engine) {
	c := e.Group("/comment")

	c.Use(middleware.AuthMiddleware())
	{
		c.POST("/publish", handlers.PublishComment())
		c.POST("/delete", middleware.IsManagerMiddleware(), handlers.DeleteComment())
	}
}
