package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const PORT = 8080
const VERSION = "3.0"


func main() {
	startServer(handler)
}

func startServer(handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc("/", handler)
	log.Printf("starting server...")
	log.Printf("Version: %s", VERSION)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request from %s", r.Header.Get("User-Agent"))
	host, err := os.Hostname()
	if err != nil {
		host = "unknown host"
	}
	resp := fmt.Sprintf("Hello from %s (Version: %s)", host, VERSION)
	_, err = w.Write([]byte(resp))
	if err != nil {
		log.Panicf("not able to write http output: %s", err)
	}
}
