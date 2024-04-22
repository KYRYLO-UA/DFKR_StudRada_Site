package main

import (
	"DFKR_STUDRADA_SITE/api"
	"net/http"
)

// import "github.com/gorilla/mux"
// import "github.com/google/uuid"

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
