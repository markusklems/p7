package main

import "github.com/goadesign/goa"

// ImgController implements the img resource.
type ImgController struct {
	*goa.Controller
}

// NewImgController creates a img controller.
func NewImgController(service *goa.Service) *ImgController {
	return &ImgController{Controller: service.NewController("ImgController")}
}
