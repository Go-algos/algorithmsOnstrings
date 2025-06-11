package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestConstructTrie(t *testing.T) {

	var inputs [][]string = [][]string{{"1", "ATA"}, {"3", "AT", "AG", "AC"}}

	var outputs [][]string = [][]string{[]string{"0->1:A", "2->3:A", "1->2:T"}, []string{"0->1:A", "1->4:C", "1->3:G", "1->2:T"}}

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
