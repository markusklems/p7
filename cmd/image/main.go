//go:generate goagen bootstrap -d github.com/markusklems/p7/cmd/image/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/markusklems/p7/cmd/image/app"
	"github.com/markusklems/p7/cmd/image/controllers"
	"github.com/markusklems/p7/cmd/image/docker"
)

func main() {
	// Create service
	service := goa.New("image")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Setup docker client
	var err error
	docker, err := docker.NewDockerClient()
	if err != nil {
		panic(err)
	}

	// Mount "image" controller onto service
	c := controllers.NewImageController(service, docker)
	app.MountImageController(service, c)

	// Mount "js" controller onto service
	c2 := controllers.NewJsController(service)
	app.MountJsController(service, c2)

	// Mount "public" controller onto service
	c3 := controllers.NewPublicController(service)
	app.MountPublicController(service, c3)

	// Mount "swagger" controller onto service
	c4 := controllers.NewSwaggerController(service)
	app.MountSwaggerController(service, c4)

	// Start service
	if err := service.ListenAndServe(":8890"); err != nil {
		service.LogError("startup", "err", err)
	}
}
