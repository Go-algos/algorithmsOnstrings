package main

import "fmt"

func constructSuffixTree(input string) []string {
	// Construct all suffixes for the given string
	for i := 0; i < len(input); i++ {
		suffix := input[i : len(input)-1]
		fmt.Printf("Suffix: %s\n", suffix)
	}
	return []string{}
}
