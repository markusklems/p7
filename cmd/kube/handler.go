package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newHandler(inCluster bool) http.Handler {
	log.Printf("New handler.")
	r := mux.NewRouter()
	r.Handle("/", functionHandler{newKubeContext(inCluster), FunctionIndexHandler})

	return r
}
