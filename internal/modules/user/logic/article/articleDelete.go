package article

import (
	"gin-blog/internal/modules/user/models"
	db2 "gin-blog/internal/modules/user/repository/db"
	"gin-blog/internal/modules/user/requests/article"
	"gin-blog/pkg/globals"
)

type DeleteLogic struct {
	err    error
	Dar    article.DeleteArticleRequest
	appCtx *globals.AppCtx
}

func NewDeleteLogic(dar article.DeleteArticleRequest, appCtx *globals.AppCtx) *DeleteLogic {
	return &DeleteLogic{
		Dar:    dar,
		appCtx: appCtx,
	}
}

func (d *DeleteLogic) Delete() error {
	d.DeleteInDb()
	return d.err
}

func (d *DeleteLogic) DeleteInDb() *DeleteLogic {
	if d.err != nil {
		return d
	}
	var a models.Article
	a.ID = d.Dar.ArticleID

	d.err = db2.DeleteArticle(db2.ArticleSqlParam{
		Db:      d.appCtx.GetDb(),
		Article: a,
	})
	return d
}
