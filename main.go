package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{"1", "ATA"}
	var result = constructTrie(input)
	resultAsString := strings.Join(result, "\n")
	fmt.Println("", resultAsString)
}
