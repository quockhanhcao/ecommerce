package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quockhanhcao/ecommerce/internal/service"
	"github.com/quockhanhcao/ecommerce/pkg/response"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	email := c.PostForm("email")
	purpose := c.PostForm("purpose")

	// Call the service to register the user
	resultCode := uc.userService.Register(email, purpose)

	if resultCode == response.SuccessCode {
		response.SuccessResponse(c, resultCode, "User registered successfully")
	} else {
		response.ErrorResponse(c, resultCode)
	}
}

// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// // controller -> service -> repo -> models -> dbs
// func (uc *UserController) GetUserById(c *gin.Context) {
// 	// response.SuccessResponse(c, response.SuccessCode, []string{"user1", "user2", "user3"})
// 	response.ErrorResponse(c, response.ErrParamInvalidCode)
// }
