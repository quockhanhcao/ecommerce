package inititialize

import (
	"fmt"

	"github.com/quockhanhcao/ecommerce/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Config loaded successfully", global.Config.LoggerConfig.FilePath)
	fmt.Println("Config loaded successfully", global.Config.LoggerConfig.LogLevel)

	InitLogger()
	global.Logger.Info("Config log ok", zap.String("ok", "success"))
	InitPostgres()
	InitRedis()

	router := InitRouter()
	router.Run(":8002")
}
