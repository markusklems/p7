package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Image is the image resource media type.
var Image = MediaType("application/vnd.image+json", func() {
	Description("A docker image")
	Attributes(func() {
		Attribute("id", String, "ID of image", func() {
			Example("ec496b4bc6e307efea8dd6e4038379e973b558ac60a0c58216a1691976c96ef6")
		})
		Attribute("fullId", String, "Complete ID of image", func() {
			Example("sha256:ec496b4bc6e307efea8dd6e4038379e973b558ac60a0c58216a1691976c96ef6")
		})
		Attribute("href", String, "API href of image", func() {
			Example("/images/1")
		})
		Attribute("containers", Integer, "Containers")
		Attribute("parentID", String, "Parent container ID")
		Attribute("repoTags", Any, "Tags of one container")
		Attribute("labels", Any, "Labels of one container")
		Attribute("created_at", DateTime, "Date of creation")
		Attribute("container", String, "Container")
		Attribute("parent", String, "Parent container")
		Attribute("repoDigests", Any, "Digests of one container")
		Attribute("comment", String, "Comment of image")
		Attribute("author", String, "Author of image")
		Attribute("architecture", String, "Architecture of image")
		Attribute("os", String, "OS of image")
		Attribute("osVersion", String, "OS version of image")
		Attribute("size", Integer, "Size of image")
		Required("id", "fullId", "href", "containers", "parentID", "repoTags", "labels", "created_at")
	})

	View("default", func() {
		Attribute("id")
		Attribute("fullId")
		Attribute("href")
		Attribute("containers")
		Attribute("parentID")
		Attribute("repoTags")
		Attribute("labels")
		Attribute("created_at")
		Attribute("container")
		Attribute("parent")
		Attribute("repoDigests")
		Attribute("comment")
		Attribute("author")
		Attribute("architecture")
		Attribute("os")
		Attribute("osVersion")
		Attribute("size")
	})

	View("tiny", func() {
		Description("tiny is the view used to list images")
		Attribute("id")
		Attribute("href")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})
