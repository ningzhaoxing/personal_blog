package token

import (
	"gin-blog/internal/modules/user/models"
	db2 "gin-blog/internal/modules/user/repository/db"
	"gin-blog/pkg/globals"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetToken 生成token
func GetToken(user models.User) (string, error) {
	claims := NewCustomClaims(&user)
	var CustomSecret = []byte(globals.C.Jwt.Secret)
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(CustomSecret) //根据前面自定义的Jwt秘钥生成token
	if err != nil {
		//返回生成的错误
		return "", err
	}
	tokenString = "Bearer " + tokenString
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return tokenString, nil
}

func ParseToken(t string, c *gin.Context, db *gorm.DB) error {
	claims := CustomClaims{}
	var CustomSecret = []byte(globals.C.Jwt.Secret)
	_, err := jwt.ParseWithClaims(t, &claims, func(token *jwt.Token) (interface{}, error) {
		return CustomSecret, nil
	})
	if err != nil {
		return err
	}
	user := claims.User
	u, err := db2.QueryUserByEmail(db2.UserSqlParam{
		Db:   db,
		User: *user,
	})

	if err != nil {
		return err
	}

	c.Set("user", u)
	return nil
}
