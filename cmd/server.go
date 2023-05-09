package main

import (
	"fmt"
	"net/http"
)

/**
 *	The main function.
 *	Starts the server.
 */
func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

/**
 *	A simple handler.
 */
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
