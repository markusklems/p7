package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Resources group related API endpoints
// together. They map to REST resources for REST
// services.

var _ = Resource("lambda", func() {
	Description("A microservice function named lambda")
	BasePath("/lambdas")
	DefaultMedia(Lambda)

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("List all avilable lambdas")
		Response(OK, func() {
			Media(CollectionOf(Lambda, func() {
				View("default")
				View("tiny")
			}))
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("show", func() {
		Routing(
			GET("/:lambda_id"),
		)
		Description("Retrieve lambda with given id")
		Params(func() {
			Param("lambda_id", Integer, "Lambda ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("watch", func() {
		Routing(
			GET("/:lambda_id/watch"),
		)
		Scheme("ws")
		Description("Retrieve lambda with given id")
		Params(func() {
			Param("lambda_id", Integer)
		})
		Response(SwitchingProtocols)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Creates a lambda")
		Payload(LambdaPayload)
		Response(Created, "^/lambdas/[0-9]+$")
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PATCH("/:lambda_id"),
		)
		Params(func() {
			Param("lambda_id", Integer)
		})
		Payload(LambdaPayload)
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("code", func() {
		Routing(
			GET("/:lambda_id/actions/code"),
		)
		Description("Retrieve lambda code with given id")
		Params(func() {
			Param("lambda_id", Integer, "Lambda ID")
		})
		Response(OK)
		Response(NotFound)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:lambda_id"),
		)
		Params(func() {
			Param("lambda_id", Integer)
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
	Files("/", "public/html/index.html")
	Files("/img/*filepath", "public/img")
	Files("/css/*filepath", "public/css")
})

var _ = Resource("js", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/js/*filepath", "public/js")
})
