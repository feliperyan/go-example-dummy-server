package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type keywordObj struct {
	Word string
}

type newQuote struct {
	Quote string `json:"quote"`
}

type keywordSearchResult struct {
	Result []string
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, "{\"health\": \"ok\"}")
}

func quote(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var nq newQuote
		err := json.NewDecoder(r.Body).Decode(&nq) // instead of unmarshal, given we are reading from a stream.
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = AddQuote(nq.Quote)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "quote added")
		return
	}

	q := newQuote{Quote: GetRandomQuote()}
	fmt.Printf("Get quote: %v \n", q)

	qjson, err := json.Marshal(q)
	if err != nil {
		fmt.Printf("error json marshalling quote: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(qjson))
}

func search(w http.ResponseWriter, r *http.Request) {
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

func request_quote_from_api(api string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		resp, err := http.Get(api)
		if err != nil {
			fmt.Printf("error on GET from backend api at %s. Err: %v\n", api, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("error reading body from api at %s. Err: %v\n", api, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		q := &newQuote{}
		err = json.Unmarshal(body, q)
		if err != nil {
			fmt.Printf("error marshalling body from api at %s. Err: %v\n", api, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "<h1>Quote: %v</h1>", q.Quote)
	}
}

func main() {
	log.Println("Initialising...")

	port, ok := os.LookupEnv("PORT")
	apiEndpoint, ok2 := os.LookupEnv("QUOTEAPIENDPOINT")

	if !ok {
		port = "8080"
	}
	log.Println("Listening on: ", port)

	http.HandleFunc("/", health)

	// Server is in backend mode
	if !ok2 {
		log.Println("BackEnd mode")
		http.HandleFunc("/quote", quote)
		http.HandleFunc("/search", search)

	} else { // Server is in frontend mode. Expects a backend it can reach to get quotes from
		log.Println("FrontEnd mode. API backend set to: ", apiEndpoint)
		http.HandleFunc("/fetch_quote", request_quote_from_api(apiEndpoint))
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
