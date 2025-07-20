//go:build wireinject
package wire
import (
	"github.com/google/wire"
	"github.com/quockhanhcao/ecommerce/internal/controller"
	"github.com/quockhanhcao/ecommerce/internal/repo"
	"github.com/quockhanhcao/ecommerce/internal/service"
)

func InitUserRouterController() (*controller.UserController, error) {
	wire.Build(service.NewUserService, repo.NewUserRepo, controller.NewUserController)

	return new(controller.UserController), nil
}
