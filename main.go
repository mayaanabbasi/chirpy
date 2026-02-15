package main

import "net/http"

func main() {
	httpServeMux := http.NewServeMux()
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: httpServeMux,
	}

	httpServer.ListenAndServe()
}
