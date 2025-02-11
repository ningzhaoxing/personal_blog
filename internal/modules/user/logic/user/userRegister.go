package user

import (
	"errors"
	"gin-blog/internal/modules/user/models"
	"gin-blog/internal/modules/user/repository/db"
	"gin-blog/internal/modules/user/requests/user"
	errors2 "gin-blog/pkg/errors"
	"gin-blog/pkg/globals"
	"gorm.io/gorm"
)

type RegisterLogic struct {
	err    error
	appCtx *globals.AppCtx
	rr     user.RegisterRequest
}

func NewRegister(appCtx *globals.AppCtx, rr user.RegisterRequest) *RegisterLogic {
	return &RegisterLogic{
		appCtx: appCtx,
		rr:     rr,
	}
}

func (r *RegisterLogic) Register() error {
	r.CheckUserIsExist().SaveUserInDb()
	if r.err != nil {
		return r.err
	}
	return nil
}

// CheckUserIsExist 查询当前用户是否已存在
func (r *RegisterLogic) CheckUserIsExist() *RegisterLogic {
	if r.err != nil {
		return r
	}

	_, err := db.QueryUserByEmail(db.UserSqlParam{
		Db:   r.appCtx.GetDb(),
		User: models.User{Email: r.rr.Email},
	})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return r
	}

	r.err = errors2.ErrUserHasExist
	return r
}

// SaveUserInDb 将用户信息存储到db中
func (r *RegisterLogic) SaveUserInDb() *RegisterLogic {
	if r.err != nil {
		return r
	}

	r.err = db.InsertUser(db.UserSqlParam{
		Db: r.appCtx.GetDb(),
		User: models.User{
			Email:    r.rr.Email,
			Password: r.rr.Password,
			Name:     r.rr.Name,
		},
	})
	return r
}
