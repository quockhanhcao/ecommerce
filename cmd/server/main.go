package main

import (
	"github.com/quockhanhcao/ecommerce/internal/router"
)

func main() {
	router := router.NewRouter()
	router.Run(":8002")
}
