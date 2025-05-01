package counters

import "strings"


func CountWords(words []string) map[string]int {
	wordCount := make(map[string]int)
	for _, word := range words {
		word = strings.ToLower(word)
		wordCount[word]++
	}

	return wordCount
}