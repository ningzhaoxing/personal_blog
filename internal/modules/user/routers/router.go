package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	//e.Use(middleware.CSRF())
	userRouter(e)
	articleRouter(e)
	commentRouter(e)
	classRouter(e)
	fileRouter(e)
}
