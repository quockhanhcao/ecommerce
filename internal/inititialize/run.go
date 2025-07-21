package inititialize

import (
	"fmt"

	"github.com/quockhanhcao/ecommerce/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	InitLogger()
	global.Logger.Info("Config log ok", zap.String("ok", "success"))
	InitPostgres()
    global.Logger.Info("Postgres init ok", zap.String("ok", "success"))
	InitRedis()

	router := InitRouter()
	router.Run(fmt.Sprintf(":%d", global.Config.ServerConfig.Port))
}
