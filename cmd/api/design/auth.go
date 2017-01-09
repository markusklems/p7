package design

import (
	_ "github.com/goadesign/goa/design"
	_ "github.com/goadesign/goa/design/apidsl"
)

//var BasicAuth = BasicAuthSecurity("BasicAuth", func() {
//	Description("Use client ID and client secret to authenticate")
//})
//
//var APIKey = APIKeySecurity("api_key", func() {
//	Query("key")
//})

//var _ = Resource("auth", func() {
//	DefaultMedia(Auth)
//	BasePath("/auth")
//
//	Action("basic", func() {
//		Security("api_key")
//		Routing(GET("/basic"))
//		Response(OK)
//	})
//})

// Auth is the auth info media type.
//var Auth = MediaType("application/vnd.goa-cellar.auth+json", func() {
//	Description("User info extracted from security token")
//	TypeName("Auth")
//
//	Attributes(func() {
//		Attribute("issuer", String, "Token issuer")
//		Attribute("id", String, "User ID")
//		Attribute("email", String, "User email", func() {
//			Format("email")
//		})
//		Required("id")
//	})
//
//	View("default", func() {
//		Attribute("issuer")
//		Attribute("id")
//		Attribute("email")
//	})
//})
