package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const unused = `unusedstr1`

func handler(w http.ResponseWriter, r *http.Request) {
	// Extract the path variable from the URL
	path := strings.TrimPrefix(r.URL.Path, "/")

	// Generate an HTML response with tags
	htmlResponse := fmt.Sprintf("<html><body><h1>Hello, It's Kubeday India 2023!!! and we are deploying image version %s</h1></body></html>", path)

	// Write the HTML response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlResponse))
}

func main() {
	// Define a handler function
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 9999
	fmt.Println("Server listening on :9999")
	log.Fatal(http.ListenAndServe(":9999", nil))
}
