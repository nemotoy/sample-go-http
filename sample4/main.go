package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct {
	greet string
}

func main() {
	mux := http.NewServeMux()
	helloHandler := newHelloHandler("hello")

	mux.Handle("/hello", helloHandler)

	http.ListenAndServe(":8080", mux)
}

func newHelloHandler(greet string) *helloHandler {
	return &helloHandler{greet: greet}
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, fmt.Sprintf("%s", h.greet))
}
