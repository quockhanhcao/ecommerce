package inititialize

import (
	"github.com/gin-gonic/gin"
	"github.com/quockhanhcao/ecommerce/global"
	"github.com/quockhanhcao/ecommerce/internal/routers"
)

func InitRouter() *gin.Engine {
	var router *gin.Engine
	if global.Config.ServerConfig.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		router = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	}

	// middlewares
	// router.Use() // logging
	// router.Use() // cors
	// router.Use() // rate limit
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := router.Group("/v1/2025")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		managerRouter.InitAdminRouter(MainGroup)
		managerRouter.InitUserRouter(MainGroup)
	}
	return router
}
