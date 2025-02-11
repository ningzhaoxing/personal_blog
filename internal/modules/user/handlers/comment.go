package handlers

import (
	"gin-blog/internal/modules/user/logic/comment"
	comment2 "gin-blog/internal/modules/user/requests/comment"
	"gin-blog/pkg/globals"
	"gin-blog/pkg/response"
	"gin-blog/pkg/token"
	"github.com/gin-gonic/gin"
	"log"
)

func PublishComment() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(c *gin.Context) {
		var cpr comment2.PublishCommentRequest
		err := c.ShouldBind(&cpr)
		if err != nil {
			log.Println("user.handlers.PublishComment().ShouldBind err=", err)
			response.HTTPFail(err, c)
			return
		}

		user, err := token.NewToken(c).GetUser()
		if err != nil {
			log.Println("user.handlers.PublishComment().GetUser err=", err)
			response.HTTPFail(err, c)
			return
		}

		cpr.UserID = user.ID

		err = comment.NewPublishLogic(appCtx, cpr).Publish()
		if err != nil {
			log.Println("user.handlers.PublishComment().NewPublishLogic err=", err)
			response.HTTPFail(err, c)
			return
		}
		response.HTTPSuccess("发布成功", response.Data{}, c)
	}
}

func DeleteComment() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(c *gin.Context) {
		var cdr comment2.DeleteCommentRequest
		err := c.ShouldBind(&cdr)
		if err != nil {
			log.Println("user.handlers.DeleteComment().ShouldBind err=", err)
			response.HTTPFail(err, c)
			return
		}

		err = comment.NewDeleteLogic(appCtx, cdr).Delete()
		if err != nil {
			log.Println("user.handlers.DeleteComment().Delete err=", err)
			response.HTTPFail(err, c)
			return
		}

		response.HTTPSuccess("删除成功", response.Data{}, c)
	}
}
