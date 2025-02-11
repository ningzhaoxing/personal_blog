package globals

import (
	EventBus2 "github.com/asaskevich/EventBus"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	E        *gin.Engine
	Db       *gorm.DB
	C        *Config
	R        *redis.Client
	EventBus EventBus2.Bus
)
