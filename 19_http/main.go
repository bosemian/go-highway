package main

import (
	"log"
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World\n"))
}

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: &helloHandler{},
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
