package server

import (
	"gin-blog/internal/modules/user/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(e *gin.Engine) {

	e.StaticFS("/static/css", http.Dir("view/static/css"))
	e.StaticFS("/static/img", http.Dir("view/static/img"))
	e.StaticFS("/static/js", http.Dir("view/static/js"))
	e.StaticFS("/uploads/", http.Dir("./uploads"))
	e.LoadHTMLGlob("view/pages/*")

	routers.InitRouter(e)
}
