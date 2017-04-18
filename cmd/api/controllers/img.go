package controllers

import "github.com/goadesign/goa"

// ImgController implements the img resource.
type ImgController struct {
	*goa.Controller
}

// NewImg creates a img controller.
func NewImg(service *goa.Service) *ImgController {
	return &ImgController{Controller: service.NewController("ImgController")}
}
