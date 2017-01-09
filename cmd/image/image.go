package main

import (
	"github.com/goadesign/goa"
	"github.com/markusklems/p7/cmd/image/app"
)

// ImageController implements the image resource.
type ImageController struct {
	*goa.Controller
}

// NewImageController creates a image controller.
func NewImageController(service *goa.Service) *ImageController {
	return &ImageController{Controller: service.NewController("ImageController")}
}

// Create runs the create action.
func (c *ImageController) Create(ctx *app.CreateImageContext) error {
	// ImageController_Create: start_implement

	// Put your logic here

	// ImageController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *ImageController) Delete(ctx *app.DeleteImageContext) error {
	// ImageController_Delete: start_implement

	// Put your logic here

	// ImageController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *ImageController) List(ctx *app.ListImageContext) error {
	// ImageController_List: start_implement

	// Put your logic here

	// ImageController_List: end_implement
	res := app.ImageCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *ImageController) Show(ctx *app.ShowImageContext) error {
	// ImageController_Show: start_implement

	// Put your logic here

	// ImageController_Show: end_implement
	res := &app.Image{}
	return ctx.OK(res)
}
