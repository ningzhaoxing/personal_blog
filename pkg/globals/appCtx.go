package globals

import (
	EventBus2 "github.com/asaskevich/EventBus"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type DbType string

type AppCtx struct {
	db       *gorm.DB
	config   *Config
	rdb      *redis.Client
	eventBus EventBus2.Bus
}

func (a *AppCtx) GetEventBus() EventBus2.Bus {
	return a.eventBus
}

func (a *AppCtx) GetDb() *gorm.DB {
	return a.db
}

func (a *AppCtx) SetDb(db *gorm.DB) {
	a.db = db
}

func (a *AppCtx) GetRdb() *redis.Client {
	return a.rdb
}

func (a *AppCtx) GetConfig() *Config {
	return a.config
}

type Options func(a *AppCtx)

func NewAppCtx(ops ...Options) *AppCtx {
	appCtx := new(AppCtx)

	for _, op := range ops {
		op(appCtx)
	}
	return appCtx
}

func NewDefaultAppCtx() *AppCtx {
	return NewAppCtx(WithOptionDb(Db), WithOptionConfig(C), WithOptionRedis(R), WithOptionEventBus(EventBus))
}

func WithOptionDb(db *gorm.DB) Options {
	return func(a *AppCtx) {
		a.db = db
	}
}

func WithOptionConfig(c *Config) Options {
	return func(a *AppCtx) {
		a.config = c
	}
}

func WithOptionRedis(r *redis.Client) Options {
	return func(a *AppCtx) {
		a.rdb = r
	}
}

func WithOptionEventBus(e EventBus2.Bus) Options {
	return func(a *AppCtx) {
		a.eventBus = e
	}
}
