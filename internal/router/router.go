package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quockhanhcao/ecommerce/internal/controller"
	"github.com/quockhanhcao/ecommerce/internal/middlewares"
)

func MiddlewareA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before -> middlewareA")
		c.Next()
		fmt.Println("After -> middlewareA")
	}
}

func MiddlewareB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before -> middlewareB")
		c.Next()
		fmt.Println("After -> middlewareB")
	}
}

func MiddlewareC(c *gin.Context) {
	fmt.Println("Before -> middlewareC")
	c.Next()
	fmt.Println("After -> middlewareC")
}

func NewRouter() *gin.Engine {
	router := gin.Default()

	// use middleware
	router.Use(middlewares.AuthMiddleware(), MiddlewareB(), MiddlewareC)

	v1 := router.Group("/v1")
	{
		v1.GET("/ping/:name", controller.NewPongController().Pong)
		v1.GET("/user/:id", controller.NewUserController().GetUserById)
	}

	return router
}
