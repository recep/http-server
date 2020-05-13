package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
