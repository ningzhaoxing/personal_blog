package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type ArticleReadingsParam struct {
	ArticleID string
	Rdb       *redis.Client
	Context   context.Context
}

func articleReadingsHashKey() string {
	return "article_readings"
}

// IncreaseArticleReadings 设置文章阅读量
func IncreaseArticleReadings(a ArticleReadingsParam, increment int64) error {
	if err := a.Rdb.HIncrBy(a.Context, articleReadingsHashKey(), a.ArticleID, increment).Err(); err != nil {
		return err
	}
	return nil
}

// GetArticleReadings 获取文章阅读量
func GetArticleReadings(a ArticleReadingsParam) (int, error) {
	readingsString, err := a.Rdb.HGet(a.Context, articleReadingsHashKey(), a.ArticleID).Result()
	if err != nil {
		return -1, err
	}
	readings, err := strconv.Atoi(readingsString)
	if err != nil {
		return -1, err
	}
	return readings, nil
}

// GetArticleReadingsHash 获取hash
func GetArticleReadingsHash(al ArticleReadingsParam) (map[string]string, error) {
	res, err := al.Rdb.HGetAll(al.Context, articleReadingsHashKey()).Result()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SetArticleReadings 设置文章阅读量
func SetArticleReadings(a ArticleReadingsParam, readings int) error {
	if err := a.Rdb.HSet(a.Context, articleReadingsHashKey(), a.ArticleID, readings).Err(); err != nil {
		return err
	}
	return nil
}
