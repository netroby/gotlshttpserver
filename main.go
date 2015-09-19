package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("About to listen on 443. Go to https://0.0.0.1:10443/")
	err := http.ListenAndServeTLS("0.0.0.0:443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
