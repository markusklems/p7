//go:generate goagen bootstrap -d github.com/markusklems/p7/cmd/image/design

package main

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/kit"
	"github.com/goadesign/goa/middleware"
	"github.com/markusklems/p7/cmd/image/app"
	"github.com/markusklems/p7/cmd/image/controllers"
	"github.com/markusklems/p7/cmd/image/docker"
)

func main() {
	// Create service
	service := goa.New("image")

	// Setup logger
	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)
	goaLogger := goakit.New(logger)
	service.WithLogger(goaLogger)

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
	ic := controllers.NewImageController(service, docker)
	app.MountImageController(service, ic)

	// Mount "js" controller onto service
	jc := controllers.NewJsController(service)
	app.MountJsController(service, jc)

	// Mount "public" controller onto service
	pc := controllers.NewPublicController(service)
	app.MountPublicController(service, pc)

	// Mount "swagger" controller onto service
	sc := controllers.NewSwaggerController(service)
	app.MountSwaggerController(service, sc)

	// Start service
	if err := service.ListenAndServe(":8890"); err != nil {
		service.LogError("startup", "err", err)
	}
}
