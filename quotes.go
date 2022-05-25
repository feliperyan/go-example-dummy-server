package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

var theQuotes []string

func init() {
	theQuotes = prepareQuotes("smart_quotes.txt")
}

func prepareQuotes(filePath string) []string {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %s", err)
		panic(1)
	}
	s := string(bytes)
	q := strings.Split(s, "\n")
	return q
}

func GetRandomQuote() string {
	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(len(theQuotes))
	s := theQuotes[num]
	return s

}

func GetQuotesForKeyword(keyword string) []string {

	fmt.Println("searching for keyword: ", keyword)

	var found []string

	for _, q := range theQuotes {
		if strings.Contains(q, keyword) {
			found = append(found, q)
		}
	}

	return found
}

func AddQuote(quote string) error {

	if quote == "" {
		return fmt.Errorf("empty quote")
	}

	theQuotes = append(theQuotes, quote)
	fmt.Println("num of quotes: ", len(theQuotes))
	return nil
}
