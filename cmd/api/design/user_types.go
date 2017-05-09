package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// LambdaPayload defines the data structure used in the create lambda request body.
// It is also the base type for the lambda media type used to render lambda.
var LambdaPayload = Type("LambdaPayload", func() {
	Description("LambdaPayload is the type used to create a lambda")

	Attribute("name", String, "Name of lambda", func() {
		MinLength(2)
	})

	Attribute("code", String, "Code that should be executed", func() {
		MinLength(16)
		MaxLength(500)
	})
	Attribute("method", func() {
		Enum("GET", "POST", "PUT", "DELETE")
	})
	Attribute("environment", func() {
		Enum("AWS", "GCLOUD", "AZURE")
	})
	Required("name", "code", "method", "environment")
})
