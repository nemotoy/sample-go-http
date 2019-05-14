package main

import (
	"fmt"
	"log"
	"net/http"
)

type helloHandler struct {
	greet string
}

func main() {
	mux := http.NewServeMux()
	helloHandler := newHelloHandler("hello")

	mux.Handle("/hello", newHeaderAuth(helloHandler))

	http.ListenAndServe(":8080", mux)
}

func newHelloHandler(greet string) *helloHandler {
	return &helloHandler{greet: greet}
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, fmt.Sprintf("%s", h.greet))
}

func newHeaderAuth(h http.Handler) http.Handler {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// common
		log.Printf("start ... ")
		id := r.Header.Get("ID")
		if id == "" {
			c := http.StatusBadRequest
			w.WriteHeader(c)
			w.Write([]byte(http.StatusText(c)))
		}
		h.ServeHTTP(w, r)
		log.Printf("end ... ")
	})
	return http.HandlerFunc(fn)
}
