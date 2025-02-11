package initialize

import "github.com/gin-gonic/gin"

func InitEngine() *gin.Engine {
	e := gin.Default()
	return e
}
