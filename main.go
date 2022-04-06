package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type keywordObj struct {
	Word string
}

type keywordSearchResult struct {
	Result []string
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, "{\"health\": \"ok\"}")
}

func quote(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, GetRandomQuote())
}

func postAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var kw keywordObj
	err := json.NewDecoder(r.Body).Decode(&kw) // instead of unmarshal given we are reading from a stream.

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultObj := keywordSearchResult{Result: GetQuotesForKeyword(kw.Word)}
	marshalledResult, err := json.Marshal(resultObj)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(marshalledResult))

}

func main() {

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	log.Println("server starting...")

	http.HandleFunc("/health", health)
	http.HandleFunc("/quote", quote)
	http.HandleFunc("/post", postAPI)

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
