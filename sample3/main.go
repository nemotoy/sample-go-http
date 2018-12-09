package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Exit(realMain())
}

func baseHandlerFunc(handler func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return baseHandler(http.HandlerFunc(handler))
}

func baseHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// common
		log.Println(r.URL, r.Method)
		handler.ServeHTTP(w, r)
	})
}

// handler
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func realMain() int {
	log.Println("start ...")

	http.Handle("/", baseHandlerFunc(index))
	http.ListenAndServe(":8080", nil)
	return 0
}

/*
	
	ref.
	https://qiita.com/tenntenn/items/b7bd54c7ba0ff90f1707
*/
