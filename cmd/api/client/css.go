// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/markusklems/p7/cmd/api/design
// --out=$(GOPATH)/src/github.com/markusklems/p7/cmd/api
// --version=v1.1.0-dirty
//
// API "p7": css Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
)

// DownloadCSS downloads /files with the given filename and writes it to the file dest.
// It returns the number of bytes downloaded in case of success.
func (c *Client) DownloadCSS(ctx context.Context, filename, dest string) (int64, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	p := path.Join("/css/", filename)
	u := url.URL{Host: c.Host, Scheme: scheme, Path: p}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.Client.Do(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			if len(b) > 0 {
				body = ": " + string(b)
			}
		}
		return 0, fmt.Errorf("%s%s", resp.Status, body)
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	return io.Copy(out, resp.Body)
}
