package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Helper function to serve HTML files
func serveFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		filename, found := routes[r.URL.Path]
		if !found {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, filename)
	}
}

var posts = []Post{
	{
		1,
		"lorem ipsum dolor",
		"Res hosti dixerit nulla habeatur scipio venisset obcaecati. Tueri bonum tali damna principio erga diligi necessarius.",
	},
	{
		2,
		"lorem ipsum door ait amet",
		"Laetitia gessisse ferri fastidii phaedrum consequentium cernantur otiosum. Summis vindicet conspectum dicat affecta animadversionem inportuno praeclara. Fugiendus amori appetere requirere robustus primorum ipsis difficilius debilitatem.",
	},
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}
