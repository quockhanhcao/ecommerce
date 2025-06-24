package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quockhanhcao/ecommerce/internal/service"
	"github.com/quockhanhcao/ecommerce/pkg/response"
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
	// response.SuccessResponse(c, response.SuccessCode, []string{"user1", "user2", "user3"})
	response.ErrorResponse(c, response.ErrParamInvalidCode)
}
