package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body := `<h1>Pablo es el mejor aunque no funcione esto</h1>`
	fmt.Fprintf(w, body)
}

func main() {
	log.Print("Hello world webapp started.")

	http.HandleFunc("/", handler)

	port := "8080"

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}