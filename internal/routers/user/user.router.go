package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	publicUserRouter := Router.Group("/user")
	{
		publicUserRouter.POST("/register")
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
