package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type helloHandler struct {
	greet string
}

func main() {
	mux := http.NewServeMux()
	helloHandler := newHelloHandler("hello")

	mux.Handle("/hello", baser(helloHandler))

	http.ListenAndServe(":8080", mux)
}

func newHelloHandler(greet string) *helloHandler {
	return &helloHandler{greet: greet}
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, fmt.Sprintf("%s", h.greet))
}

func baser(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// common
		log.Printf("start ... ")
		time.Sleep(1 * time.Second)
		h.ServeHTTP(w, r)
		log.Printf("end ... ")
	})
}
