package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		word := "Hello CI! " + req.URL.Path
		w.Write([]byte(word))
	})
test test
	log.Fatal(http.ListenAndServe(":8080", nil))
}
