package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// private router
	privateAdminRouter := Router.Group("/admin/user")
	// privateUserRouter.Use(rateLimit())
	// privateUserRouter.Use(Authen())
	// privateUserRouter.Use(Permission())
	{
		privateAdminRouter.GET("/active_user")
	}
}
