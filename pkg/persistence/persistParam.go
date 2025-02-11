package persistence

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type PersistParam struct {
	Db  *gorm.DB
	Rdb *redis.Client
}
