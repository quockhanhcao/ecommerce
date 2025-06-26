package inititialize

import (
	"github.com/quockhanhcao/ecommerce/global"
	"github.com/quockhanhcao/ecommerce/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.LoggerConfig)
}
