package controllers

import "github.com/goadesign/goa"

// CSSController implements the css resource.
type CSSController struct {
	*goa.Controller
}

// NewCSS creates a css controller.
func NewCSS(service *goa.Service) *CSSController {
	return &CSSController{Controller: service.NewController("CSSController")}
}
