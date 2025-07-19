package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	privateUserRouter := Router.Group("/admin/user")
	// privateUserRouter.Use(rateLimit())
	// privateUserRouter.Use(Authen())
	// privateUserRouter.Use(Permission())
	{
		privateUserRouter.POST("/active_user")
	}
}
