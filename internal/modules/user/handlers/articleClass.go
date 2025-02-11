package handlers

import (
	"gin-blog/internal/modules/user/logic/class"
	class2 "gin-blog/internal/modules/user/requests/class"
	"gin-blog/pkg/globals"
	"gin-blog/pkg/response"
	"gin-blog/pkg/token"
	"github.com/gin-gonic/gin"
	"log"
)

func GetClass() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(c *gin.Context) {
		list, err := class.NewGetClassListLogic(appCtx).Result()
		if err != nil {
			log.Println("user.handlers.GetClass().Result err=", err)
			response.Fail(err, "articleEdit", c)
			return
		}

		u, err := token.NewToken(c).GetUser()
		if err != nil {
			log.Println("user.handlers.GetClass().GetUser err=", err)
			response.Fail(err, "articleEdit", c)
			return
		}

		response.Success("", "articleEdit", response.Data{
			Content: map[string]any{
				"Class": list,
				"Uid":   u.ID,
			},
		}, c)
	}
}

func CreateClass() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(c *gin.Context) {
		var ccr class2.ClassCreateRequest
		err := c.ShouldBind(&ccr)
		if err != nil {
			log.Println("user.handlers.CreateClass().ShouldBind err=", err)
			response.HTTPFail(err, c)
			return
		}

		err = class.NewCreateClassLogic(ccr, appCtx).Create()
		if err != nil {
			log.Println("user.handlers.CreateClass().NewCreateClassLogic err=", err)
			response.HTTPFail(err, c)
			return
		}
		response.HTTPSuccess("创建成功", response.Data{}, c)
	}
}
