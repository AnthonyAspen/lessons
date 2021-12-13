package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// Exercise: Maps
// Implement WordCount.
// It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.

// You might find strings.Fields helpful.

func WordCount(s string) map[string]int {
	strArr := strings.Fields(s)
	strMap := make(map[string]int)

	// loop throw array of strings
	for _,word := range strArr{
		// check if the key is presented
       _,ok := strMap[word]
		// if yes - increase the count, if not, then add it 
		if ok {
			strMap[word] += 1
		} else{
			strMap[word] = 1
		}
	}
	return strMap
}

func main() {
	wc.Test(WordCount)
}
