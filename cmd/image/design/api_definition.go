package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// API defines the microservice endpoint and
// other global properties.
var _ = API("image", func() {
	Description("Docker image creation server")
	Scheme("http")
	Host("localhost:8890")
	BasePath("/image")

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})
