package main

import "github.com/goadesign/goa"

// CSSController implements the css resource.
type CSSController struct {
	*goa.Controller
}

// NewCSSController creates a css controller.
func NewCSSController(service *goa.Service) *CSSController {
	return &CSSController{Controller: service.NewController("CSSController")}
}
