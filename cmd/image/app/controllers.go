// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/markusklems/p7/cmd/image/design
// --out=$(GOPATH)/src/github.com/markusklems/p7/cmd/image
// --version=v1.1.0-dirty
//
// API "image": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ImageController is the controller interface for the Image actions.
type ImageController interface {
	goa.Muxer
	Create(*CreateImageContext) error
	Delete(*DeleteImageContext) error
	List(*ListImageContext) error
	Show(*ShowImageContext) error
}

// MountImageController "mounts" a Image resource controller on the given service.
func MountImageController(service *goa.Service, ctrl ImageController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateImageContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateImagePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/image/images", ctrl.MuxHandler("Create", h, unmarshalCreateImagePayload))
	service.LogInfo("mount", "ctrl", "Image", "action", "Create", "route", "POST /image/images")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteImageContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/image/images/:image_id", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Image", "action", "Delete", "route", "DELETE /image/images/:image_id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListImageContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/image/images", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Image", "action", "List", "route", "GET /image/images")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowImageContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/image/images/:image_id", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Image", "action", "Show", "route", "GET /image/images/:image_id")
}

// unmarshalCreateImagePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateImagePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createImagePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// JsController is the controller interface for the Js actions.
type JsController interface {
	goa.Muxer
	goa.FileServer
}

// MountJsController "mounts" a Js resource controller on the given service.
func MountJsController(service *goa.Service, ctrl JsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/js/*filepath", ctrl.MuxHandler("preflight", handleJsOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/js/*filepath", "public/js")
	h = handleJsOrigin(h)
	service.Mux.Handle("GET", "/js/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Js", "files", "public/js", "route", "GET /js/*filepath")

	h = ctrl.FileHandler("/js/", "public/js/index.html")
	h = handleJsOrigin(h)
	service.Mux.Handle("GET", "/js/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Js", "files", "public/js/index.html", "route", "GET /js/")
}

// handleJsOrigin applies the CORS response headers corresponding to the origin.
func handleJsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// PublicController is the controller interface for the Public actions.
type PublicController interface {
	goa.Muxer
	goa.FileServer
}

// MountPublicController "mounts" a Public resource controller on the given service.
func MountPublicController(service *goa.Service, ctrl PublicController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/ui", ctrl.MuxHandler("preflight", handlePublicOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/ui", "public/html/index.html")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/ui", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "public/html/index.html", "route", "GET /ui")
}

// handlePublicOrigin applies the CORS response headers corresponding to the origin.
func handlePublicOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/swagger-ui/*filepath", "swagger-ui/")
	service.Mux.Handle("GET", "/swagger-ui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/", "route", "GET /swagger-ui/*filepath")

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger-ui/", "swagger-ui/index.html")
	service.Mux.Handle("GET", "/swagger-ui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/index.html", "route", "GET /swagger-ui/")
}
