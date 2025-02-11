package initialize

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func InitLogger() {
	file, err := os.OpenFile("web_app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("日志文件打开错误")
	}
	gin.DefaultWriter = io.MultiWriter(file)
}
