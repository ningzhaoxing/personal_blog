package class

import (
	db2 "gin-blog/internal/modules/user/repository/db"
	"gin-blog/internal/modules/user/requests/class"
	"gin-blog/pkg/globals"
)

type CreateClassLogic struct {
	err    error
	ccr    class.ClassCreateRequest
	appCtx *globals.AppCtx
}

func NewCreateClassLogic(ccr class.ClassCreateRequest, appCtx *globals.AppCtx) *CreateClassLogic {
	return &CreateClassLogic{
		ccr:    ccr,
		appCtx: appCtx,
	}
}

func (c *CreateClassLogic) Create() error {
	c.SaveClassInDb()

	return c.err
}

func (c *CreateClassLogic) SaveClassInDb() *CreateClassLogic {
	if c.err != nil {
		return c
	}

	c.err = db2.CreateClass(db2.ClassSqlParam{
		Db:    c.appCtx.GetDb(),
		Class: c.ccr,
	})
	return c
}
