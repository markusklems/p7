package controllers

import (
	"github.com/goadesign/goa"
	"github.com/markusklems/p7/cmd/api/app"
)

// HealthController implements the health resource.
type HealthController struct {
	*goa.Controller
}

// NewHealth creates a health controller.
func NewHealth(service *goa.Service) *HealthController {
	return &HealthController{Controller: service.NewController("HealthController")}
}

// Check runs the check action.
func (c *HealthController) Check(ctx *app.CheckHealthContext) error {
	return ctx.OK([]byte("OK"))
}
