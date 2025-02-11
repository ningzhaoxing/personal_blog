package user

import (
	"gin-blog/internal/modules/user/models"
	"gin-blog/internal/modules/user/repository/db"
	"gin-blog/internal/modules/user/requests/user"
	"gin-blog/pkg/globals"
)

type UpdateHeadPhotoLogic struct {
	err    error
	uhur   user.UserHeadPhotoUpdateRequest
	appCtx *globals.AppCtx
}

func NewUpdateHeadPhotoLogic(uhur user.UserHeadPhotoUpdateRequest, appCtx *globals.AppCtx) *UpdateHeadPhotoLogic {
	return &UpdateHeadPhotoLogic{
		uhur:   uhur,
		appCtx: appCtx,
	}
}

func (u *UpdateHeadPhotoLogic) Update() error {
	u.SaveInDb()
	return u.err
}

func (u *UpdateHeadPhotoLogic) SaveInDb() *UpdateHeadPhotoLogic {
	if u.err != nil {
		return u
	}

	var user models.User
	user.ID = u.uhur.UserID
	user.HeadPhoto = u.uhur.Url

	u.err = db.UpdateHeadPhoto(db.UserSqlParam{
		Db:   u.appCtx.GetDb(),
		User: user,
	})
	return u
}
