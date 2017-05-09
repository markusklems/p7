package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Lambda defines the media type used to render lambdas.
var Lambda = MediaType("application/vnd.lambda+json", func() {
	Description("A microservice function named lambda")
	Reference(LambdaPayload)
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique lambda ID", func() {
			Example(1)
		})
		Attribute("href", String, "API href for making requests on the lambda", func() {
			Example("/lambdas/1")
		})
		Attribute("code", String, "Source code of the user lambda")
		Attribute("created_at", DateTime, "Date of creation")
		Attribute("updated_at", DateTime, "Date of last update")
		// Attributes below inherit from the base type
		Attribute("name", String, "Name of lambda")
		Attribute("method", String, "HTTP method triggering a lambda")
		Attribute("environment", String, "Execution environment used to run a lambda")
		Required("id", "href", "name", "code", "method")
	})

	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("name")
		Attribute("method")
		Attribute("environment")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("method")
	})

	View("code", func() {
		Attribute("code")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("code")
		Attribute("method")
		Attribute("environment")
		Attribute("created_at")
		Attribute("updated_at")
		Attribute("links")
	})
})
