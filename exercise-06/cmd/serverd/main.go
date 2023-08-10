package main

import (
	"github.com/xuanvan229/go23/exercise-06/pkg/repo"
	"github.com/xuanvan229/go23/exercise-06/pkg/routes"
)

func main() {
	repo.SetupRepo()
	routers := routes.InitRouters()
	err := routers.Run(":8080")
	if err != nil {
		return
	}
}
