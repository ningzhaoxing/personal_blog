package persistence

import (
	"context"
	"gin-blog/internal/modules/user/models"
	"gin-blog/internal/modules/user/repository/db"
	redis2 "gin-blog/internal/modules/user/repository/redis"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type ArticleLikePersistTask struct {
	err  error
	rdb  *redis.Client
	db   *gorm.DB
	data map[string]string
}

func NewArticleLikePersistTask(p PersistParam) *ArticleLikePersistTask {
	return &ArticleLikePersistTask{
		rdb: p.Rdb,
		db:  p.Db,
	}
}

func (a *ArticleLikePersistTask) Execute() error {
	a.GetDataFromRedis().PersistDataToDB()
	return a.err
}

func (a *ArticleLikePersistTask) GetDataFromRedis() *ArticleLikePersistTask {
	if a.err != nil {
		return a
	}

	a.data, a.err = redis2.GetArticleLikeHash(redis2.ArticleLikeParam{
		Rdb:     a.rdb,
		Context: context.Background(),
	})
	return a
}

func (a *ArticleLikePersistTask) PersistDataToDB() *ArticleLikePersistTask {
	if a.err != nil {
		return a
	}

	tx := a.db.Begin()
	defer tx.Rollback()
	for k, v := range a.data {
		aid, err := strconv.Atoi(k)
		if err != nil {
			a.err = err
			return a
		}

		likes, err := strconv.Atoi(v)
		if err != nil {
			a.err = err
			return a
		}

		var article models.Article
		article.Likes = likes
		article.ID = uint(aid)
		// 更新阅读量
		if err := db.UpdateArticleLikes(db.ArticleSqlParam{
			Db:      tx,
			Article: article,
		}); err != nil {
			a.err = err
			return a
		}
	}

	tx.Commit()
	log.Println("批量持久化成功... ")
	return a
}
