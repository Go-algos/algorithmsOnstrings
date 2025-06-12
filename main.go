package main

import (
	"fmt"
)

func main() {
	/*
		"AATCGGGTTCAATCGGGGT",
					"2",
					"ATCG",
					"GGGT"}
	*/
	result := trieMatching("AATCGGGTTCAATCGGGGT", "GGGT")
	fmt.Println("", result)
}
