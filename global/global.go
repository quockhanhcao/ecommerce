package global

import (
	"github.com/quockhanhcao/ecommerce/pkg/logger"
	"github.com/quockhanhcao/ecommerce/pkg/settings"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	DB     *gorm.DB
)
