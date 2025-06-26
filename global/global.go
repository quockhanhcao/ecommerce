package global

import (
	"github.com/quockhanhcao/ecommerce/pkg/logger"
	"github.com/quockhanhcao/ecommerce/pkg/settings"
)

var (
	Config settings.Config
    Logger *logger.LoggerZap
)
