package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!") // this will be printed as a response when you request /
	})
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil) // this will start the HTTP server (blocking operation) and listen on port 8080
}
