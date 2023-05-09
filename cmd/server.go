package main

import (
	"fmt"
	"net/http"
	"os"
)

/**
 *	The main function.
 *	Starts the server.
 */
func main() {
	// Get port
	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":"+port, nil)
}

/**
 *	A simple handler.
 */
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
