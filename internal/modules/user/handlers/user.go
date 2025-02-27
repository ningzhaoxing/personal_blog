package handlers

import (
	"gin-blog/internal/modules/user/logic/user"
	"gin-blog/internal/modules/user/models"
	user2 "gin-blog/internal/modules/user/requests/user"
	"gin-blog/pkg/globals"
	"gin-blog/pkg/response"
	"gin-blog/pkg/token"
	"github.com/gin-gonic/gin"
	"log"
)

func Login() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(c *gin.Context) {
		var loginRequest user2.LoginRequest
		err := c.ShouldBind(&loginRequest)
		if err != nil {
			log.Println("user.handlers.Login().ShouldBind err=", err)
			response.HTTPFail(err, c)
			return
		}

		uid, err := user.NewLogin(loginRequest, appCtx).Login()
		if err != nil {
			log.Println("user.handlers.Login().Login err=", err)
			response.HTTPFail(err, c)
			return
		}

		tokenString, err := token.GetToken(models.User{Email: loginRequest.Email})
		if err != nil {
			log.Println("user.handlers.Register().CreateToken err=", err)
			response.HTTPFail(err, c)
			return
		}
		log.Println(tokenString, "-----")
		c.SetCookie("token", tokenString, 360000, "/", "localhost", false, true)
		response.HTTPSuccess("登录成功", response.Data{Content: uid}, c)
	}
}

func Register() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(c *gin.Context) {
		var registerRequest user2.RegisterRequest
		if err := c.ShouldBind(&registerRequest); err != nil {
			log.Println("user.handlers.Register().ShouldBind err=", err)
			response.HTTPFail(err, c)
			return
		}

		if err := user.NewRegister(appCtx, registerRequest).Register(); err != nil {
			log.Println("user.handlers.Register().Register err=", err)
			response.HTTPFail(err, c)
			return
		}

		tokenString, err := token.GetToken(models.User{Email: registerRequest.Email})
		if err != nil {
			log.Println("user.handlers.Register().CreateToken err=", err)
			response.HTTPFail(err, c)
			return
		}

		c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
		response.HTTPSuccess("注册成功", response.Data{}, c)
	}
}

func LikeArticle() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(c *gin.Context) {
		var lar user2.LikeArticleRequest
		if err := c.ShouldBindJSON(&lar); err != nil {
			log.Println("user.handlers.LikeArticle().ShouldBind err=", err)
			response.Fail(err, "article", c)
			return
		}

		u, err := token.NewToken(c).GetUser()
		if err != nil {
			log.Println("user.handlers.LikeArticle().GetUser err=", err)
			response.Fail(err, "article", c)
			return
		}

		lar.UserID = int(u.ID)

		if err := user.NewLikeArticleLogic(lar, appCtx).Do(); err != nil {
			log.Println("user.handlers.LikeArticle().Do err=", err)
			response.Fail(err, "article", c)
			return
		}

		response.HTTPSuccess("点赞成功", response.Data{}, c)
	}
}

// UpdateUserHeadPhoto 修改用户头像
func UpdateUserHeadPhoto() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(c *gin.Context) {
		var uhur user2.UserHeadPhotoUpdateRequest
		err := c.ShouldBind(&uhur)
		if err != nil {
			log.Println("user.handlers.UpdateUserHeadPhoto().ShouldBind err=", err)
			response.HTTPFail(err, c)
			return
		}

		u, err := token.NewToken(c).GetUser()
		if err != nil {
			log.Println("user.handlers.UpdateUserHeadPhoto().GetUser err=", err)
			response.HTTPFail(err, c)
			return
		}

		uhur.UserID = u.ID

		err = user.NewUpdateHeadPhotoLogic(uhur, appCtx).Update()
		if err != nil {
			log.Println("user.handlers.UpdateUserHeadPhoto().Update err=", err)
			response.HTTPFail(err, c)
			return
		}
		response.HTTPSuccess("修改成功", response.Data{}, c)
	}
}
