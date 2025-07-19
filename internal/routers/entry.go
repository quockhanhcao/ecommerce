package routers

import (
	"github.com/quockhanhcao/ecommerce/internal/routers/manager"
	"github.com/quockhanhcao/ecommerce/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
