package initialize

import (
	"context"
	"fmt"
	"gin-blog/pkg/globals"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func InitRedis(c *globals.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.Redis.Host, c.Redis.Port),
		Password: "",
		DB:       c.Redis.Db,
		PoolSize: c.Redis.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := rdb.Ping(ctx).Result() // 检查连接redis是否成功
	if err != nil {
		log.Printf("Connect HttpFail: %v \n\n", err)
		log.Println("redis_addr: ", c.Redis.Host)
		panic(err)
	}
	fmt.Println("redis连接成功")
	return rdb
}
