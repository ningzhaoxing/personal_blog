package subscribes

import (
	"context"
	"gin-blog/internal/modules/user/repository/redis"
	"gin-blog/pkg/globals"
	"strconv"
)

type ArticleReadingsTopic struct {
	ArticleID uint
	Increment int
	AppCtx    *globals.AppCtx
}

// UpdateArticleReadingsSubscribe 修改文章阅读量事件
func UpdateArticleReadingsSubscribe(a ArticleReadingsTopic) error {
	err := redis.IncreaseArticleReadings(redis.ArticleReadingsParam{
		ArticleID: strconv.Itoa(int(a.ArticleID)),
		Rdb:       a.AppCtx.GetRdb(),
		Context:   context.Background(),
	}, int64(a.Increment))
	return err
}
