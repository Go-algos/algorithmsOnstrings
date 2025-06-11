package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestConstructTrie(t *testing.T) {

	var inputs [][]string = [][]string{
		{"1", "ATA"},
		{"3", "AT", "AG", "AC"},
		{"3", "ATAGA", "ATC", "GAT"}}

	var outputs [][]string = [][]string{
		{"0->1:A", "2->3:A", "1->2:T"},
		{"0->1:A", "1->4:C", "1->3:G", "1->2:T"},
		{"0->1:A", "1->2:T", "2->3:A", "3->4:G", "4->5:A", "2->6:C", "0->7:G", "7->8:A", "8->9:T"}}

	for index := 0; index < len(inputs); index++ {
		input := inputs[index]
		output := outputs[index]
		var result = constructTrie(input)

		// Sorting resolve issue in the first test
		// Ref: https://gobyexample.com/sorting
		slices.Sort(result)
		slices.Sort(output)

		//if result != output {
		if !reflect.DeepEqual(result, output) {
			t.Errorf(`constructTrie() = %q %v`, result, output)
		}
	}
}
