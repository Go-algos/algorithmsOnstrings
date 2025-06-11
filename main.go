package main

import (
	"fmt"
)

func main() {
	input := []string{"3", "AT", "AG", "AC"}
	var result = constructTrie(input)
	//resultAsString := strings.Join(result, "\n")
	/*
		Expected:
		0->1:A
		1->4:C
		1->3:G
		1->2:T
	*/
	fmt.Println("", result)
}
