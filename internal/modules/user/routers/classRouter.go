package routers

import (
	"gin-blog/internal/modules/user/handlers"
	"gin-blog/internal/server/middleware"
	"github.com/gin-gonic/gin"
)

func classRouter(e *gin.Engine) {
	c := e.Group("/class")

	c.Use(middleware.AuthMiddleware())
	{
		c.POST("/create", middleware.IsManagerMiddleware(), handlers.CreateClass())
	}

	c.GET("/articles", handlers.GetArticlesByClass())
}
