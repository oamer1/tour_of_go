package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	words_count := make(map[string]int)
	for _, word := range strings.Fields(s) {
		words_count[word]++

	}

	return words_count
}

func main() {
	wc.Test(WordCount)
}
