package token

import (
	"gin-blog/internal/modules/user/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims struct {
	User *models.User
	jwt.StandardClaims
}

func NewCustomClaims(user *models.User) *CustomClaims {
	user.Password = ""
	tokenExpireDuration := time.Now().Add(7 * 24 * time.Hour)
	return &CustomClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpireDuration.Unix(), //token的有效期
			IssuedAt:  time.Now().Unix(),          //token发放的时间
			Issuer:    "blog",                     //作者
			Subject:   "sql token",                //主题
		},
	}
}
