package main

import (
	"fmt"
)

func main() {
	input := []string{"1", "ATA"}
	var result = constructTrie(input)
	//resultAsString := strings.Join(result, "\n")
	/*
		{"0->1:A", "2->3:A", "1->2:T"}
	*/
	fmt.Println("", result)
}
