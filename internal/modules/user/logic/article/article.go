package article

import (
	"context"
	"gin-blog/internal/modules/user/logic/user"
	"gin-blog/internal/modules/user/models"
	db2 "gin-blog/internal/modules/user/repository/db"
	"gin-blog/internal/modules/user/repository/redis"
	"gin-blog/internal/modules/user/requests/article"
	user2 "gin-blog/internal/modules/user/requests/user"
	"gin-blog/pkg/globals"
	"gin-blog/pkg/subscribes"
	"gin-blog/utils/time"
	"strconv"
)

type GetArticleLogic struct {
	err    error
	appCtx *globals.AppCtx
	Ar     article.ArticleRequest
}

func NewGetArticleLogic(appCtx *globals.AppCtx, ar article.ArticleRequest) *GetArticleLogic {
	return &GetArticleLogic{
		appCtx: appCtx,
		Ar:     ar,
	}
}

func (g *GetArticleLogic) Result() (models.Article, error) {
	g.Ar.Article.IsLiked, g.err = g.GetArticleInDb().
		GetArticleLikeNum(g.Ar.ID).
		GetArticleReadings(g.Ar.ID).
		IncreaseArticleReadings(1).
		GetArticleCommentNum(g.Ar.ID).
		TimeTrans().
		UserLikeOrNot()

	return g.Ar.Article, g.err
}

func (g *GetArticleLogic) GetArticleInDb() *GetArticleLogic {
	if g.err != nil {
		return g
	}

	g.Ar.Article.ID = g.Ar.ID
	g.Ar.Article, g.err = db2.QueryArticle(db2.ArticleSqlParam{
		Db:      g.appCtx.GetDb(),
		Article: g.Ar.Article,
	})
	return g
}

// UserLikeOrNot 检查用户是否点赞
func (g *GetArticleLogic) UserLikeOrNot() (bool, error) {
	// 游客
	if g.Ar.UserID == 0 {
		return false, nil
	}
	// 用户
	exist, err := user.NewLikeArticleLogic(user2.LikeArticleRequest{
		ArticleID: int(g.Ar.ID),
		UserID:    int(g.Ar.UserID),
	}, g.appCtx).UidExistInArticleSetOrNot()

	if err != nil {
		return false, err
	}
	return exist, nil
}

// GetArticleLikeNum 获取文章点赞数量
func (g *GetArticleLogic) GetArticleLikeNum(aid uint) *GetArticleLogic {
	if g.err != nil {
		return g
	}
	g.Ar.Article.Likes, g.err = redis.GetArticleLikeNum(redis.ArticleLikeParam{
		ArticleID: strconv.Itoa(int(aid)),
		Rdb:       g.appCtx.GetRdb(),
		Context:   context.Background(),
	})

	return g
}

// GetArticleReadings 获取文章阅读量
func (g *GetArticleLogic) GetArticleReadings(aid uint) *GetArticleLogic {
	if g.err != nil {
		return g
	}
	g.Ar.Article.Readings, g.err = redis.GetArticleReadings(redis.ArticleReadingsParam{
		ArticleID: strconv.Itoa(int(aid)),
		Rdb:       g.appCtx.GetRdb(),
		Context:   context.Background(),
	})
	return g
}

// IncreaseArticleReadings 增加文章阅读量
func (g *GetArticleLogic) IncreaseArticleReadings(increment int) *GetArticleLogic {
	if g.err != nil {
		return g
	}
	g.appCtx.GetEventBus().Publish(globals.ArticleReadingsTopic, subscribes.ArticleReadingsTopic{
		ArticleID: g.Ar.ID,
		Increment: increment,
		AppCtx:    g.appCtx,
	})
	return g
}

// GetArticleCommentNum 获取评论数量
func (g *GetArticleLogic) GetArticleCommentNum(aid uint) *GetArticleLogic {
	if g.err != nil {
		return g
	}

	g.Ar.Article.CommentNum, g.err = db2.QueryArticleCommentNum(db2.CommentSqlParam{Db: g.appCtx.GetDb()}, aid)
	return g
}

// TimeTrans 时间换算
func (g *GetArticleLogic) TimeTrans() *GetArticleLogic {
	if g.err != nil {
		return g
	}

	g.Ar.Article.PublishAt, g.err = time.TimeAgo(g.Ar.Article.CreatedAt)
	return g
}
