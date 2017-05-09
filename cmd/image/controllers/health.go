package main

import (
	"github.com/goadesign/goa"
	"github.com/markusklems/p7/cmd/image/controllers/app"
)

// HealthController implements the health resource.
type HealthController struct {
	*goa.Controller
}

// NewHealthController creates a health controller.
func NewHealthController(service *goa.Service) *HealthController {
	return &HealthController{Controller: service.NewController("HealthController")}
}

// Check runs the check action.
func (c *HealthController) Check(ctx *app.CheckHealthContext) error {
	// HealthController_Check: start_implement

	// Put your logic here

	// HealthController_Check: end_implement
	return nil
}
