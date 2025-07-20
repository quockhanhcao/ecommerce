package service

import (
	"github.com/quockhanhcao/ecommerce/internal/repo"
	"github.com/quockhanhcao/ecommerce/pkg/response"
)

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepo.GetInfoUser()
// }

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepo
}

// Register implements IUserService.
func (u *userService) Register(email string, purpose string) int {
    // check email existence
    if u.userRepo.GetUserByEmail(email) {
        return response.ErrUserAlreadyExists
    }
    return response.SuccessCode
}

func NewUserService(userRepo repo.IUserRepo) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
