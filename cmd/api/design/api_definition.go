package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// API defines the microservice endpoint and
// other global properties.
var _ = API("p7", func() {
	Title("p7 api component")
	Description("Microservice management service")
	Scheme("http")
	Host("127.0.0.1:8888")
	BasePath("/p7")

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
