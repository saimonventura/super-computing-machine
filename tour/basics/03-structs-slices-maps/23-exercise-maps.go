// Exercise: Maps
// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.
// Test function runs a test suite against the provided function and prints success or failure.

// You might find strings.Fields https://go.dev/pkg/strings/#Fields helpful.

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// Allocate memory for the map
	m := make(map[string]int)

	// Split the string into a slice of strings
	words := strings.Fields(s)

	// Count the words
	for _, word := range words {
		m[word]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
