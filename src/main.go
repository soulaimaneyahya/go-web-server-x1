package main

import (
	"log"
	"net/http"
)

func main() {
	// Define routes and their corresponding handlers
	http.HandleFunc("/", serveFileHandler())
	http.HandleFunc("/about", serveFileHandler())
	http.HandleFunc("/posts", postsHandler)

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
