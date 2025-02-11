package routers

import (
	"gin-blog/internal/modules/user/handlers"
	"gin-blog/internal/server/middleware"
	"github.com/gin-gonic/gin"
)

func articleRouter(e *gin.Engine) {
	a := e.Group("/article")

	a.GET("/get", handlers.Article())
	a.GET("/list", handlers.ArticleList())

	a.Use(middleware.AuthMiddleware())
	{
		a.GET("/edit", middleware.IsManagerMiddleware(), handlers.GetClass())
		a.POST("/publish", handlers.PublishArticle())
		a.POST("/like", handlers.LikeArticle())
		a.POST("/delete", middleware.IsManagerMiddleware(), handlers.DeleteArticle())
	}
}
