package user

import (
	"gin-blog/internal/modules/user/models"
	"gin-blog/internal/modules/user/repository/db"
	"gin-blog/internal/modules/user/requests/user"
	"gin-blog/pkg/errors"
	"gin-blog/pkg/globals"
)

type LoginLogic struct {
	err    error
	lr     user.LoginRequest
	appCtx *globals.AppCtx
}

func NewLogin(lr user.LoginRequest, appCtx *globals.AppCtx) *LoginLogic {
	return &LoginLogic{
		lr:     lr,
		appCtx: appCtx,
	}
}

func (l *LoginLogic) Login() (uint, error) {
	l.CheckEmailAndPassword()
	if l.err != nil {
		return 0, l.err
	}
	return l.lr.UID, nil
}

// CheckEmailAndPassword 检查账号密码是否正确
func (l *LoginLogic) CheckEmailAndPassword() *LoginLogic {
	if l.err != nil {
		return l
	}

	var u models.User
	u, l.err = db.QueryUserByEmail(db.UserSqlParam{
		Db:   l.appCtx.GetDb(),
		User: models.User{Email: l.lr.Email},
	})
	l.lr.UID = u.ID

	if l.err != nil {
		return l
	}

	if u.Password != l.lr.Password {
		l.err = errors.ErrEmailOrPassword
	}
	return l
}
