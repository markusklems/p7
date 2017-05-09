package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("health", func() {

	BasePath("/health")

	Action("check", func() {
		Routing(
			GET("/check"),
		)
		Description("Perform health check.")
		Response(OK, "text/plain")
	})
})
