package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
}

func main() {
	err := godotenv.Load("./local.env")
	if err != nil {
		fmt.Printf("Failed to load local.env. err: #v", err)
		os.Exit(1)
	}
	addr := os.Getenv("ADDR")

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", HelloHandler)
	mux.HandleFunc("/v1/time", CurrentTimeHandler)

	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

/*
	ref.
	https://drstearns.github.io/tutorials/gomiddleware/

*/
