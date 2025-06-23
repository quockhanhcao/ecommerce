package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quockhanhcao/ecommerce/internal/controller"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/ping/:name", controller.NewPongController().Pong)
		v1.GET("/user/:id", controller.NewUserController().GetUserById)
	}

	return router
}
