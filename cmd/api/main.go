package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloAgian)
	http.ListenAndServe(":8080", mux)
}

func helloAgian(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi there"))
}
