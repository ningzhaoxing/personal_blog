package handlers

import (
	"gin-blog/internal/modules/user/logic/article"
	"gin-blog/internal/modules/user/logic/class"
	"gin-blog/internal/modules/user/logic/comment"
	"gin-blog/internal/modules/user/logic/user"
	article2 "gin-blog/internal/modules/user/requests/article"
	comment2 "gin-blog/internal/modules/user/requests/comment"
	"gin-blog/pkg/globals"
	"gin-blog/pkg/response"
	"gin-blog/pkg/token"
	"github.com/gin-gonic/gin"
	"log"
)

// ..

func ArticleList() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(c *gin.Context) {
		var alr article2.ArticleListRequest
		if err := c.ShouldBind(&alr); err != nil {
			log.Println("user.handlers.ArticleList().ShouldBind err=", err)
			response.Fail(err, "home", c)
			return
		}

		list, err := article.NewArticleListLogic(appCtx, alr).ArticleList()
		if err != nil {
			log.Println("user.handlers.ArticleList().ArticleList err=", err)
			response.Fail(err, "home", c)
			return
		}

		u, err := user.NewOwnerInfoLogic(appCtx).Get()
		if err != nil {
			log.Println("user.handlers.ArticleList().Get err=", err)
			response.Fail(err, "home", c)
			return
		}

		response.Success("", "home", response.Data{
			Header: u,
			Content: map[string]any{
				"User":     u,
				"Articles": list,
			},
		}, c)
	}
}

func PublishArticle() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(c *gin.Context) {
		var par article2.PublishArticleRequest
		if err := c.ShouldBindJSON(&par); err != nil {
			log.Println("user.handlers.PublishArticle().GetUser err=", err)
			response.HTTPFail(err, c)
			return
		}

		user, err := token.NewToken(c).GetUser()
		if err != nil {
			log.Println("user.handlers.ArticleList().GetUser err=", err)
			response.Fail(err, "home", c)
			return
		}
		par.User = *user

		aid, err := article.NewPublishLogic(appCtx, par).Publish()
		if err != nil {
			log.Println("user.handlers.PublishArticle().GetUser err=", err)
			response.HTTPFail(err, c)
			return
		}

		response.HTTPSuccess("发布成功", response.Data{
			Content: aid,
		}, c)
	}
}

func Article() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(c *gin.Context) {
		var ar article2.ArticleRequest
		if err := c.ShouldBind(&ar); err != nil {
			log.Println("user.handlers.Article().ShouldBind err=", err)
			response.Fail(err, "article", c)
			return
		}

		art, err := article.NewGetArticleLogic(appCtx, ar).Result()
		if err != nil {
			log.Println("user.handlers.Article().ShouldBind err=", err)
			response.Fail(err, "article", c)
			return
		}

		comments, err := comment.NewGetCommentListLogic(appCtx, comment2.CommentListRequest{ArticleID: ar.ID}).CommentList()
		if err != nil {
			log.Println("user.handlers.Article().CommentList err=", err)
			response.Fail(err, "article", c)
			return
		}

		response.Success("", "article", response.Data{
			Content: map[string]any{
				"User":     art.User,
				"Article":  art,
				"Comments": comments,
			},
		}, c)
	}
}

func GetArticlesByClass() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(c *gin.Context) {
		var acr article2.ArticleClassRequest
		err := c.ShouldBind(&acr)
		if err != nil {
			log.Println("user.handlers.GetArticlesByClass().ShouldBind err=", err)
			response.Fail(err, "articleClass", c)
			return
		}

		classes, err := class.NewGetClassListLogic(appCtx).Result()
		if err != nil {
			log.Println("user.handlers.GetArticlesByClass().NewGetClassListLogic err=", err)
			response.Fail(err, "articleClass", c)
			return
		}

		articles, err := article.NewGetArticlesByClassLogic(appCtx, acr).Result()
		if err != nil {
			log.Println("user.handlers.GetArticlesByClass().NewGetArticlesByClassLogic err=", err)
			response.Fail(err, "articleClass", c)
			return
		}

		response.Success("", "articleClass", response.Data{
			Content: map[string]any{
				"Classes":  classes,
				"Articles": articles,
			},
		}, c)
	}
}

func DeleteArticle() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(c *gin.Context) {
		var dar article2.DeleteArticleRequest
		err := c.ShouldBindJSON(&dar)
		if err != nil {
			log.Println("user.handlers.DeleteArticle().ShouldBindJSON err=", err)
			response.HTTPFail(err, c)
			return
		}

		err = article.NewDeleteLogic(dar, appCtx).Delete()
		if err != nil {
			log.Println("user.handlers.DeleteArticle().NewDeleteLogic err=", err)
			response.HTTPFail(err, c)
			return
		}

		response.HTTPSuccess("删除成功", response.Data{}, c)
	}
}
