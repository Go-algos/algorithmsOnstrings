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

	var expectedOutputs [][]string = [][]string{
		{"0->1:A", "2->3:A", "1->2:T"},
		{"0->1:A", "1->4:C", "1->3:G", "1->2:T"},
		{"0->1:A", "1->2:T", "2->3:A", "3->4:G", "4->5:A", "2->6:C", "0->7:G", "7->8:A", "8->9:T"}}

	for index := 0; index < len(inputs); index++ {
		input := inputs[index]

		expected := expectedOutputs[index]
		var actual = constructTrie(input)

		// Sorting resolve issue in the first test
		// Ref: https://gobyexample.com/sorting
		slices.Sort(actual)
		slices.Sort(expected)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf(`Test %q structTrie() = %q`, index, actual)
			t.Errorf(`Test %q constructTrie() = %q`, index, expected)
		}
	}
}
