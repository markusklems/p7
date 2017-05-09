package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// JSONError wraps a concrete Code and Message, `Code` is
// is an integer error code, `Message` is the error message.
type JSONError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// JSONProgress describes a Progress. terminalFd is the fd of the current terminal,
// Start is the initial value for the operation. Current is the current status and
// value of the progress made towards Total. Total is the end value describing when
// we made 100% progress for an operation.
type JSONProgress struct {
	terminalFd uintptr
	Current    int64 `json:"current,omitempty"`
	Total      int64 `json:"total,omitempty"`
	Start      int64 `json:"start,omitempty"`
	// If true, don't show xB/yB
	HideCounts bool `json:"hidecounts,omitempty"`
}

// JSONMessage defines a message struct. It describes
// the created time, where it from, status, ID of the
// message. It's used for docker events.
type JSONMessage struct {
	Stream          string        `json:"stream,omitempty"`
	Status          string        `json:"status,omitempty"`
	Progress        *JSONProgress `json:"progressDetail,omitempty"`
	ProgressMessage string        `json:"progress,omitempty"` //deprecated
	ID              string        `json:"id,omitempty"`
	From            string        `json:"from,omitempty"`
	Time            int64         `json:"time,omitempty"`
	TimeNano        int64         `json:"timeNano,omitempty"`
	Error           *JSONError    `json:"errorDetail,omitempty"`
	ErrorMessage    string        `json:"error,omitempty"` //deprecated
	// Aux contains out-of-band data, such as digests for push signing.
	Aux *json.RawMessage `json:"aux,omitempty"`
}

// Docker represents a docker client
// https://godoc.org/github.com/docker/docker/client
type Docker struct {
	client *client.Client
}

// Message is a docker client response message
type Message struct {
	Stream      string
	ErrorDetail string
}

func getImageBuildResponse(body []byte) (*JSONMessage, error) {
	var s = new(JSONMessage)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("Couldn't decode image build response:", err)
	}
	return s, err
}

// NewDockerClient creates a new docker client instance
func NewDockerClient() (*Docker, error) {
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.24", nil, defaultHeaders)
	//cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &Docker{
		client: cli,
	}, err
}

// ListImages prints and lists all available images
func (d *Docker) ListImages() []types.ImageSummary {
	images, err := d.client.ImageList(context.Background(), types.ImageListOptions{All: true})
	if err != nil {
		panic(err)
	}
	for _, image := range images {
		fmt.Printf("Containers: %d, Created: %d, ID: %s, ParentID: %s, Size: %d, RepoTags: %s\n", image.Containers, image.Created, image.ID[:10], image.ParentID, image.Size, image.RepoTags)
	}
	return images
}

// BuildImage builds a new docker image
func (d *Docker) BuildImage(buildContext io.Reader, codePath string, tag string, provider string) {
	tags := []string{tag}
	buildArgs := map[string]*string{
		"CODE_PATH": &codePath,
	}
	imageBuildOptions := types.ImageBuildOptions{
		BuildArgs: buildArgs,
		//CPUPeriod:  30,
		//CPUQuota:   10,
		//CPUSetCPUs: "2",
		//CPUSetMems: "12",
		//CPUShares:  20,
		//Memory:     256,
		//MemorySwap: 512,
		//ShmSize:    10,
		Tags:       tags,
		Dockerfile: setDockerfile(provider),
	}
	buildResponse, err := d.client.ImageBuild(context.Background(), buildContext, imageBuildOptions)
	if err != nil {
		fmt.Printf("Couldn't build image: %s\n", err)
	} else {
		body, err := ioutil.ReadAll(buildResponse.Body)
		if err != nil {
			fmt.Printf("Couldn't decode docker client build response: %s\n", err)
		}
		fmt.Printf(string(body))
	}
}

// CreateImage creates a new docker image.
func (d *Docker) CreateImage(parentRef string) {}

// RemoveImage tries to delete a image with given id
func (d *Docker) RemoveImage(imageID string) {
	tImageDel, err := d.client.ImageRemove(context.Background(), imageID, types.ImageRemoveOptions{})
	if err != nil {
		fmt.Printf("Couldn't remove image with id: %s, %s\n", imageID, err)
	} else {
		fmt.Printf("Removed image with id: %s, response: %s\n", imageID, tImageDel)
	}
}

// PushImage tries to push a image to registry.
func (d *Docker) PushImage(ref string) {
	ioReader, err := d.client.ImagePush(context.Background(), ref, types.ImagePushOptions{})
	if err != nil {
		fmt.Printf("Couldn't push image with ref: %s, %s\n", ref, err)
		ioReader.Close()
	} else {
		fmt.Printf("Pushed image with ref: %s\n", ref)
	}
}

// InspectImage returns the image information and its raw representation.
func (d *Docker) InspectImage(imageID string) (types.ImageInspect, []byte) {
	tImageIns, bytes, err := d.client.ImageInspectWithRaw(context.Background(), imageID)
	if err != nil {
		fmt.Printf("Couldn't retrieve image inspect information: %s\n", err)
	}
	return tImageIns, bytes
}

func setDockerfile(provider string) string {
	var dockerfile string
	switch provider {
	case "GCLOUD":
		dockerfile = "Dockerfile_gcloud"
	case "AZURE":
		dockerfile = "Dockerfile_azure"
	default:
		dockerfile = "Dockerfile_aws"
	}
	return dockerfile
}

func streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
