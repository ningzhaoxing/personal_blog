package article

import (
	"gin-blog/internal/modules/user/models"
	db2 "gin-blog/internal/modules/user/repository/db"
	"gin-blog/internal/modules/user/requests/article"
	"gin-blog/pkg/globals"
)

type GetArticlesByClassLogic struct {
	err      error
	Acr      article.ArticleClassRequest
	appCtx   *globals.AppCtx
	articles []models.Article
}

func NewGetArticlesByClassLogic(appCtx *globals.AppCtx, acr article.ArticleClassRequest) *GetArticlesByClassLogic {
	return &GetArticlesByClassLogic{
		err:    nil,
		appCtx: appCtx,
		Acr:    acr,
	}
}

func (g *GetArticlesByClassLogic) Result() ([]models.Article, error) {
	g.GetArticlesByClass()
	return g.articles, g.err
}

func (g *GetArticlesByClassLogic) GetArticlesByClass() *GetArticlesByClassLogic {
	if g.err != nil {
		return g
	}

	g.articles, g.err = db2.QueryArticlesByClass(db2.ArticleSqlParam{
		Db:      g.appCtx.GetDb(),
		Article: models.Article{ArticleClassID: g.Acr.ArticleClassID},
	})
	return g
}
