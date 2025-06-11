package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Problem: Construct a Trie from a Collection of Patterns
func constructTrie(input []string) []string {
	// temp output

	numberOfPatterns, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Println("Failed to extract number of patterns from input")
		panic(err)
	}

	root := 0
	trie := []string{strconv.Itoa(root) + "->" + strconv.Itoa(root+1) + ":" + input[1][0:1]}
	totalNodes := 1

	// Main loop to iterate over patterns
	for index := 0; index < numberOfPatterns; index++ {
		currentNode := root
		pattern := input[index+1]

		currentEdge := pattern[currentNode : currentNode+1] // take first edge from the pattern

		for i := 0; i < len(pattern); i++ {
			currentSymb := string(pattern[i])

			if strings.HasSuffix(currentEdge, currentSymb) {
				currentNode += 1
			} else {
				a := strconv.Itoa(currentNode)
				b := strconv.Itoa(totalNodes + 1)
				trieItem := a + "->" + b + ":" + currentSymb
				currentEdge = strings.TrimSpace(pattern[i : i+1])
				totalNodes += 1
				currentNode = totalNodes
				trie = append(trie, trieItem)
			}
		}
	}

	return trie
}
