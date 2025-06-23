package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quockhanhcao/ecommerce/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> service -> repo -> models -> dbs
func (uc *UserController) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUser(),
		"users":   []string{"a", "b", "c"},
	})
}
