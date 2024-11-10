package main

import (
	"fmt"
	"log"
	"net/http"
)

const serverPort = ":8080"

func main() {
	// Define routes and their corresponding handlers
	http.HandleFunc("/", serveFileHandler())
	http.HandleFunc("/about", serveFileHandler())
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/hello-world", helloWorldHandler)

	// Start the server
	fmt.Printf("Server running at http://localhost%s\n", serverPort)

	if err := http.ListenAndServe(serverPort, nil); err != nil {
		log.Fatal(err)
	}
}
