package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// hello world

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello awesome world from Blorn!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listening for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
