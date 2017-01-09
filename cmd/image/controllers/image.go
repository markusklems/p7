package controllers

import (
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/goadesign/goa"
	"github.com/markusklems/p7/cmd/image/app"
	"github.com/markusklems/p7/cmd/image/docker"
)

// ImageController implements the image resource.
type ImageController struct {
	*goa.Controller
	docker *docker.Docker
}

// ToImageMedia builds an image media type from an docker image model.
func ToImageMedia(image *types.ImageSummary) *app.Image {
	return &app.Image{
		ID:         strings.TrimPrefix(image.ID, "sha256:"),
		FullID:     image.ID,
		Href:       app.ImageHref(image.ID),
		Containers: int(image.Containers), // goa design doesn't support int64 yet
		ParentID:   image.ParentID,
		RepoTags:   image.RepoTags,
		Labels:     image.Labels,
		CreatedAt:  time.Unix(image.Created, 0),
	}
}

// NewImageController creates an image controller.
func NewImageController(service *goa.Service, docker *docker.Docker) *ImageController {
	return &ImageController{
		Controller: service.NewController("ImageController"),
		docker:     docker,
	}
}

// Create runs the create action.
func (c *ImageController) Create(ctx *app.CreateImageContext) error {
	dockerBuildContext, err := os.Open("./Dockerfile.tar")
	if err != nil {
		panic(err)
	}
	defer dockerBuildContext.Close()
	payload := ctx.Payload
	c.docker.BuildImage(dockerBuildContext, payload.CodePath, payload.Tag, *payload.Provider)
	return nil
}

// Delete runs the delete action.
func (c *ImageController) Delete(ctx *app.DeleteImageContext) error {
	c.docker.RemoveImage(ctx.ImageID)
	return ctx.NoContent()
}

// List runs the list action.
func (c *ImageController) List(ctx *app.ListImageContext) error {
	dockerImages := c.docker.ListImages()
	res := make(app.ImageCollection, len(dockerImages))
	for i, image := range dockerImages {
		img := ToImageMedia(&image)
		res[i] = img
	}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *ImageController) Show(ctx *app.ShowImageContext) error {
	imageInfo, _ := c.docker.InspectImage(ctx.ImageID)
	// goa design doesn't support int64 yet
	size := int(imageInfo.Size)
	res := &app.Image{
		ID:        strings.TrimPrefix(imageInfo.ID, "sha256:"),
		FullID:    imageInfo.ID,
		Href:      app.ImageHref(imageInfo.ID),
		RepoTags:  imageInfo.RepoTags,
		Parent:    &imageInfo.Parent,
		Comment:   &imageInfo.Created,
		Container: &imageInfo.Container,
		Os:        &imageInfo.Os,
		OsVersion: &imageInfo.OsVersion,
		Size:      &size,
	}
	return ctx.OK(res)
}
