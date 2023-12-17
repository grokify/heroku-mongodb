package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if len(port) == 0 {
		port = "8080"
	}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		if _, err := io.WriteString(w, "Hello, world!\n"); err != nil {
			log.Print(err.Error())
		}
	}

	timeout := time.Second

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	svr := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		IdleTimeout:       timeout,
		ReadHeaderTimeout: timeout,
		ReadTimeout:       timeout,
		WriteTimeout:      timeout,
		MaxHeaderBytes:    1 << 20,
	}

	log.Fatal(svr.ListenAndServe())
}
