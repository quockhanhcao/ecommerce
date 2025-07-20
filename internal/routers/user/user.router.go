package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quockhanhcao/ecommerce/internal/wire"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	userController, _ := wire.InitUserRouterController()

	publicUserRouter := Router.Group("/user")
	{
		publicUserRouter.POST("/register", userController.Register)
		publicUserRouter.POST("/otp")
	}
	// private router
	privateUserRouter := Router.Group("/user")
	// privateUserRouter.Use(rateLimit())
	// privateUserRouter.Use(Authen())
	// privateUserRouter.Use(Permission())
	{
		privateUserRouter.GET("/info")
	}
}
