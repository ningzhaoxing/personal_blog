package routers

import (
	"gin-blog/internal/modules/user/handlers"
	"gin-blog/internal/server/middleware"
	"github.com/gin-gonic/gin"
)

func fileRouter(e *gin.Engine) {
	f := e.Group("/file")

	f.POST("/upload", middleware.AuthMiddleware(), handlers.UploadFile())
}
