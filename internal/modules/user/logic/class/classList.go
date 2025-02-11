package class

import (
	"gin-blog/internal/modules/user/models"
	"gin-blog/internal/modules/user/repository/db"
	"gin-blog/pkg/globals"
)

type GetClassListLogic struct {
	err    error
	appCtx *globals.AppCtx
	list   []models.ArticleClass
}

func NewGetClassListLogic(appCtx *globals.AppCtx) *GetClassListLogic {
	return &GetClassListLogic{
		appCtx: appCtx,
	}
}

func (g *GetClassListLogic) Result() ([]models.ArticleClass, error) {
	g.GetClassList()
	return g.list, g.err
}

func (g *GetClassListLogic) GetClassList() *GetClassListLogic {
	if g.err != nil {
		return g
	}
	g.list, g.err = db.QueryClassList(db.ClassSqlParam{Db: g.appCtx.GetDb()})
	return g
}
