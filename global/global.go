package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Db      *gorm.DB
	RedisDB *redis.Client
	Logger  *zap.Logger
)
