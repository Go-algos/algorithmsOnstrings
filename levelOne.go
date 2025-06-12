package main

import (
	"strconv"
)

// Ref: how to define struct
type edge struct {
	nodeFrom int
	label    string
}

// Problem: Construct a Trie from a Collection of Patterns
func constructTrie(input []string) []string {
	numberOfPatterns, _ := strconv.Atoi(input[0])
	root := 0
	trie := []string{strconv.Itoa(root) + "->" + strconv.Itoa(root+1) + ":" + input[1][0:1]}
	outgoingEdges := make(map[edge]int)
	outgoingEdges[edge{nodeFrom: 0, label: input[1][0:1]}] = root + 1 // nodeTo always predefined

	// Ref: How to work with maps in Go => https://gobyexample.com/maps
	// We are going to store edge to node index relationship

	// Main loop => to iterate over patterns
	for index := 0; index < numberOfPatterns; index++ {
		currentNode := root
		pattern := input[index+1]

		for i := 0; i < len(pattern); i++ {
			currentSymbol := string(pattern[i])

			// there is an outgoing edge from this node with the given symbol
			nodeTo, exists := outgoingEdges[edge{nodeFrom: currentNode, label: currentSymbol}]
			if exists {
				currentNode = nodeTo

				//}
			} else {

				newEdge := edge{nodeFrom: currentNode, label: currentSymbol}
				outgoingEdges[newEdge] = len(outgoingEdges) + 1
				// newNode added

				trieItem := strconv.Itoa(currentNode) + "->" + strconv.Itoa(len(outgoingEdges)) + ":" + currentSymbol
				trie = append(trie, trieItem)
				currentNode = outgoingEdges[newEdge]
			}
		}
	}
	return trie
}
