package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Resources group related API endpoints
// together. They map to REST resources for REST
// services.

var _ = Resource("image", func() {

	DefaultMedia(Image)
	BasePath("/images")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all images.")
		Response(OK, CollectionOf(Image))
	})

	Action("show", func() {
		Routing(
			GET("/:image_id"),
		)
		Description("Retrieve image with given id. IDs 1 and 2 pre-exist in the system.")
		Params(func() {
			Param("image_id", String, "Image ID")
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new image")
		Payload(ImagePayload, func() {
			Member("codePath")
			Member("tag")
			Member("provider")
			Required("codePath", "tag")
		})
		Response(Created, "/image/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:image_id"),
		)
		Params(func() {
			Param("image_id", String, "Image ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

// Serve swagger docu and swagger-ui files
var _ = Resource("swagger", func() {
	Description("The API Swagger specification")

	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger-ui/*filepath", "swagger-ui/")
})

var _ = Resource("public", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/ui", "public/html/index.html")
})

var _ = Resource("js", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/js/*filepath", "public/js")
})
