// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "p7": health Resource Client
//
// Command:
// $ goagen
// --design=github.com/markusklems/p7/cmd/api/design
// --out=$(GOPATH)/src/github.com/markusklems/p7/cmd/api
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CheckHealthPath computes a request path to the check action of health.
func CheckHealthPath() string {

	return fmt.Sprintf("/p7/health/check")
}

// Perform health check.
func (c *Client) CheckHealth(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewCheckHealthRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCheckHealthRequest create the request corresponding to the check action endpoint of the health resource.
func (c *Client) NewCheckHealthRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
