//go:generate goagen bootstrap -d github.com/markusklems/p7/cmd/kube/design

package main

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/kit"
	"github.com/goadesign/goa/middleware"
	"github.com/markusklems/p7/cmd/kube/app"
	"github.com/markusklems/p7/cmd/kube/controllers"
	"github.com/markusklems/p7/cmd/kube/k8s"
)

func main() {
	// Create service
	service := goa.New("kube")

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

	// Setup kubeContext
	kube := k8s.NewKubeContext(true, goaLogger)

	// Mount "css" controller
	cc := controllers.NewCSSController(service)
	app.MountCSSController(service, cc)
	// Mount "health" controller
	hc := controllers.NewHealthController(service)
	app.MountHealthController(service, hc)
	// Mount "img" controller
	ic := controllers.NewImgController(service)
	app.MountImgController(service, ic)
	// Mount "job" controller
	jc := controllers.NewJobController(service, kube)
	app.MountJobController(service, jc)
	// Mount "js" controller
	jsc := controllers.NewJsController(service)
	app.MountJsController(service, jsc)
	// Mount "public" controller
	pc := controllers.NewPublicController(service)
	app.MountPublicController(service, pc)
	// Mount "swagger" controller
	sc := controllers.NewSwaggerController(service)
	app.MountSwaggerController(service, sc)

	// Start service
	if err := service.ListenAndServe(":8880"); err != nil {
		service.LogError(err.Error())
	}
}
