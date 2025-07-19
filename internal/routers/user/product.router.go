package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	publicProductRouter := Router.Group("/product")
	{
		publicProductRouter.GET("/search")
		publicProductRouter.GET("/detail/:id")
	}
	// private router
}
