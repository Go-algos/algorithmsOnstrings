package main

import "strings"

func prefixMatching(text string, trie string) bool {
	//symbol := text[0:1]
	v := 0 // root of the trie
	leafDetector := trie[len(trie)-2 : 1]
	for i := 0; i < len(text); i++ {
		triePart := trie[v : v+1]
		if triePart == leafDetector { // we detected leaf + there are prev matches
			return true
		} else if strings.HasSuffix(trie, text[v:i]) {
			v = i
		} else {
			return false
		}
	}
	return false
}

func trieMatching(text string, trie string) []int {
	result := []int{}
	for i := 0; i < len(text); i++ {
		currentText := text[i : len(text)-1]
		//fmt.Printf("Current text: %s\n", currentText)
		if prefixMatching(currentText, trie) {
			result = append(result, i)
		}
	}
	return result
}
