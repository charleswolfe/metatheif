package main

import (
	"log"
	"net/http"
	"metatheif/cmd/metatheif/handlers"
)

// main function to start the HTTP server
func main() {
	http.HandleFunc("/fetch", handlers.FetchHandler)

	log.Println("Server started at http://localhost:8001")
	log.Fatal(http.ListenAndServe(":8080", nil))
}