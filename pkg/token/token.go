package token

import (
	"gin-blog/internal/modules/user/models"
	"gin-blog/pkg/errors"
	"github.com/gin-gonic/gin"
)

type Token struct {
	c *gin.Context
}

func NewToken(c *gin.Context) *Token {
	return &Token{c: c}
}

// GetUser 获取user
func (t *Token) GetUser() (*models.User, error) {
	u, exist := t.c.Get("user")

	if !exist {
		return nil, errors.ErrTokenAuth
	}
	us := u.(models.User)

	return &us, nil
}

func (t *Token) IsVisitor() bool {
	tokenString, err := t.c.Cookie("token")

	if err != nil || tokenString == "" {
		return true
	}
	return false
}
