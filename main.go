package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	httpServeMux := http.NewServeMux()

	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: httpServeMux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(httpServer.ListenAndServe())
}
