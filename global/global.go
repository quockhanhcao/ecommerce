package global

import (
	"github.com/quockhanhcao/ecommerce/pkg/logger"
	"github.com/quockhanhcao/ecommerce/pkg/settings"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	DB     *gorm.DB
	Rdb    *redis.Client
)
