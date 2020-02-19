package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	router := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
