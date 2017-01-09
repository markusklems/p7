package design

import (
	_ "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// ImagePayload defines the data structure used in the create image request body.
// It is also the base type for the image media type used to render images.
var ImagePayload = Type("ImagePayload", func() {
	Attribute("codePath", func() {
		MinLength(8)
		Example("http://127.0.0.1/code")
	})
	Attribute("tag", func() {
		MinLength(3)
		Example("nginx")
	})
	Attribute("provider", func() {
		Enum("aws", "gcloud", "azure")
	})
})
