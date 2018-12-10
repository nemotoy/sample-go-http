package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct{}

func main() {
	mux := http.NewServeMux()
	helloHandler := newHelloHandler()

	mux.Handle("/hello", helloHandler)

	http.ListenAndServe(":8080", mux)
}

func newHelloHandler() *helloHandler {
	return &helloHandler{}
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World")
}
