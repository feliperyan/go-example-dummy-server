package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

var theQuotes []string

func prepareQuotes(filePath string) []string {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %s", err)
		return nil
	}
	s := string(bytes)
	q := strings.Split(s, "\n")
	return q
}

func GetRandomQuote() string {
	if theQuotes == nil {
		theQuotes = prepareQuotes("smart_quotes.txt")
	}

	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(len(theQuotes))
	s := theQuotes[num]
	return fmt.Sprintf("A smart man once said: \n\n\t%s", s)

}

func GetQuotesForKeyword(keyword string) []string {

	fmt.Println("searching for keyword: ", keyword)

	if theQuotes == nil {
		theQuotes = prepareQuotes("smart_quotes.txt")
	}

	found := []string{""}

	for _, q := range theQuotes {
		if strings.Contains(q, keyword) {
			found = append(found, q)
		}
	}

	return found
}
