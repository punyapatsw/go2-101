package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	resp := []byte(`{"name":"Punyapat"}`)
	w.Header().Add("content-type", "application/json")
	w.Write(resp)
}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":123456", nil))
}
