package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8081", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, r.Method)
	fmt.Fprintf(w, http.StatusText(http.StatusOK))
}
