package main

import (
	"fmt"
	"strings"
)

func prefixMatching(text string, trie string) bool {
	// TODO: this need to be performance optimized ???
	fmt.Printf("Current text: %s - %s\n", text, trie)
	return strings.HasPrefix(text, trie)
}

func trieMatching(text string, trie string) []int {
	result := []int{}
	for i := 0; i < len(text); i++ {
		currentText := text[i:len(text)]
		//fmt.Printf("Current text: %s\n", currentText)
		if prefixMatching(currentText, trie) {
			result = append(result, i)
		}
	}
	return result
}
