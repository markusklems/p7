// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/markusklems/p7/cmd/image/design
// --out=$(GOPATH)/src/github.com/markusklems/p7/cmd/image
// --version=v1.1.0-dirty
//
// API "image": Application User Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

// imagePayload user type.
type imagePayload struct {
	CodePath *string `form:"codePath,omitempty" json:"codePath,omitempty" xml:"codePath,omitempty"`
	Provider *string `form:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
	Tag      *string `form:"tag,omitempty" json:"tag,omitempty" xml:"tag,omitempty"`
}

// Publicize creates ImagePayload from imagePayload
func (ut *imagePayload) Publicize() *ImagePayload {
	var pub ImagePayload
	if ut.CodePath != nil {
		pub.CodePath = ut.CodePath
	}
	if ut.Provider != nil {
		pub.Provider = ut.Provider
	}
	if ut.Tag != nil {
		pub.Tag = ut.Tag
	}
	return &pub
}

// ImagePayload user type.
type ImagePayload struct {
	CodePath *string `form:"codePath,omitempty" json:"codePath,omitempty" xml:"codePath,omitempty"`
	Provider *string `form:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
	Tag      *string `form:"tag,omitempty" json:"tag,omitempty" xml:"tag,omitempty"`
}
