package main

import (
	"net"
	"net/http"
	"time"
)

// Server is an interface that serves incoming request based on custom handler
type Server struct {
	// port to listen on
	addr string

	// handler that is responsible to manage incoming requests
	handler http.Handler

	// instance reference
	instance *http.Server
}

func newServer(host string, port string, handler http.Handler) *Server {
	addr := net.JoinHostPort(host, port)
	return &Server{
		addr:    addr,
		handler: handler,
		instance: &http.Server{
			Addr:           addr,
			Handler:        handler,
			ReadTimeout:    30 * time.Second,
			WriteTimeout:   30 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (srv Server) start() error {
	err := srv.instance.ListenAndServe()
	return err
}
