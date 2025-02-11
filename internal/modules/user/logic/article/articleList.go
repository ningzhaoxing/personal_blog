package article

import (
	"gin-blog/internal/modules/user/models"
	db2 "gin-blog/internal/modules/user/repository/db"
	"gin-blog/internal/modules/user/requests/article"
	"gin-blog/pkg/globals"
	"gin-blog/utils/time"
)

type GetArticleListLogic struct {
	err      error
	appCtx   *globals.AppCtx
	alr      article.ArticleListRequest
	Articles []models.Article
}

func NewArticleListLogic(appCtx *globals.AppCtx, alr article.ArticleListRequest) *GetArticleListLogic {
	return &GetArticleListLogic{
		appCtx: appCtx,
		alr:    alr,
	}
}

func (a *GetArticleListLogic) ArticleList() ([]models.Article, error) {
	a.GetArticleListByPage().GetArticleLikes().GetArticleReadings().GetArticleCommentNum().TimeTrans()
	if a.err != nil {
		return nil, a.err
	}
	return a.Articles, a.err
}

// GetArticleListByPage 从数据库分页查找文章列表
func (a *GetArticleListLogic) GetArticleListByPage() *GetArticleListLogic {
	if a.err != nil {
		return a
	}
	a.Articles, a.err = db2.QueryArticleListByPage(db2.ArticleSqlParam{Db: a.appCtx.GetDb(), Alr: a.alr})
	return a
}

// GetArticleReadings 从redis更新阅读量
func (a *GetArticleListLogic) GetArticleReadings() *GetArticleListLogic {
	articleLogic := NewGetArticleLogic(a.appCtx, article.ArticleRequest{})
	for i := range a.Articles {
		articleLogic.GetArticleReadings(a.Articles[i].ID)
		a.err = articleLogic.err
		if a.err != nil {
			return a
		}
		a.Articles[i].Readings = articleLogic.Ar.Article.Readings
	}
	return a
}

// GetArticleLikes 从redis更新点赞量
func (a *GetArticleListLogic) GetArticleLikes() *GetArticleListLogic {
	articleLogic := NewGetArticleLogic(a.appCtx, article.ArticleRequest{})
	for i := range a.Articles {
		articleLogic.GetArticleLikeNum(a.Articles[i].ID)
		a.err = articleLogic.err
		if a.err != nil {
			return a
		}
		a.Articles[i].Likes = articleLogic.Ar.Article.Likes
	}
	return a
}

// GetArticleCommentNum 获取评论数量
func (a *GetArticleListLogic) GetArticleCommentNum() *GetArticleListLogic {
	articleLogic := NewGetArticleLogic(a.appCtx, article.ArticleRequest{})
	for i := range a.Articles {
		articleLogic.GetArticleCommentNum(a.Articles[i].ID)
		a.err = articleLogic.err
		if a.err != nil {
			return a
		}
		a.Articles[i].CommentNum = articleLogic.Ar.Article.CommentNum
	}
	return a
}

// TimeTrans 时间换算
func (a *GetArticleListLogic) TimeTrans() *GetArticleListLogic {
	for i := range a.Articles {
		a.Articles[i].PublishAt, a.err = time.TimeAgo(a.Articles[i].CreatedAt)
		if a.err != nil {
			return a
		}
	}
	return a
}
