package main

import (
	"net/http"
	"testing"
)

var host = "127.0.0.1"
var port = "8080"
var handler http.Handler

func TestNewServer(t *testing.T) {
	server := newServer(host, port, handler)
	if server == nil {
		t.Error("Return from newServer should not be nil")
	}
}

//func TestServerStart(t *testing.T) {
//	server := newServer(host, port, handler)
//	err := server.start()
//	if err != nil {
//		t.Error("newServer should listen on defined port")
//	}
//}
