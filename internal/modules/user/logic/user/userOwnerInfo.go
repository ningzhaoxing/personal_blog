package user

import (
	"gin-blog/internal/modules/user/models"
	"gin-blog/internal/modules/user/repository/db"
	"gin-blog/pkg/globals"
)

type OwnerInfoLogic struct {
	err    error
	appCtx *globals.AppCtx
	user   models.User
}

func NewOwnerInfoLogic(appCtx *globals.AppCtx) *OwnerInfoLogic {
	return &OwnerInfoLogic{
		appCtx: appCtx,
	}
}

func (o *OwnerInfoLogic) Get() (models.User, error) {
	o.GetOwnerInfo()
	return o.user, o.err
}

func (o *OwnerInfoLogic) GetOwnerInfo() *OwnerInfoLogic {
	if o.err != nil {
		return o
	}

	o.user, o.err = db.QueryOwnerUser(db.UserSqlParam{
		Db: o.appCtx.GetDb(),
	})
	return o
}
