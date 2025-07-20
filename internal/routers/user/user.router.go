package user

import (
	"github.com/gin-gonic/gin"
	"github.com/quockhanhcao/ecommerce/internal/controller"
	"github.com/quockhanhcao/ecommerce/internal/repo"
	"github.com/quockhanhcao/ecommerce/internal/service"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	// non dependency injection

	userRepo := repo.NewUserRepo()

	userService := service.NewUserService(userRepo)

	userHandler := controller.NewUserController(userService)

	publicUserRouter := Router.Group("/user")
	{
		publicUserRouter.POST("/register", userHandler.Register)
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
