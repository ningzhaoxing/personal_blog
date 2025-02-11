package response

import (
	"gin-blog/pkg/globals"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	"net/http"
)

func HTTPSuccess(msg string, data Data, c *gin.Context) {
	c.JSON(http.StatusOK, &AppData{
		Code: globals.CodeSucess,
		Data: data,
		Msg:  msg,
	})
}

func HTTPFail(err error, c *gin.Context) {
	c.JSON(http.StatusOK, &AppData{
		Code: globals.CodeServerErr,
		Msg:  err.Error(),
	})
}

func Success(msg string, name string, data Data, c *gin.Context) {
	c.HTML(http.StatusOK, "index", &AppData{
		Code: globals.CodeSucess,
		Msg:  msg,
		Data: data,
		Name: name,
		Csrf: csrf.TemplateField(c.Request),
	})
}

func Fail(err error, name string, c *gin.Context) {
	c.HTML(http.StatusOK, "index", &AppData{
		Code: globals.CodeServerErr,
		Msg:  err.Error(),
		Name: name,
	})
}
