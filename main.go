package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, "{\"health\": \"ok\"}")
}

func quote(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, GetRandomQuote())
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hit /quote or /health")
}

func main() {

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	log.Println("server starting...")

	http.HandleFunc("/health", health)
	http.HandleFunc("/quote", quote)

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
