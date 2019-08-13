package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if len(port) == 0 {
		port = "8080"
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
