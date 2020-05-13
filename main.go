package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8081", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Method)
	//fmt.Fprintln(w, r.Response.Status)
}
