package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	seen_words := make(map[string]int)
	for i := 0; i < len(words); i++ {
		seen_words[words[i]]++
	}
	fmt.Println(seen_words)
	return seen_words
}

func main() {
	wc.Test(WordCount)
}
