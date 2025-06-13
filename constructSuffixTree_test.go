package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestConstructSuffixTree(t *testing.T) {

	var inputs = []string{
		"A$",
		"ACA$",
		"ATAAATG$"}

	var expectedOutputs = [][]string{
		{
			"A$",
			"$"},
		{
			"$",
			"A",
			"$",
			"CA$",
			"CA$"},
		{
			"AAATG$",
			"G$",
			"T",
			"ATG$",
			"TG$",
			"A",
			"A",
			"AAATG$",
			"G$",
			"T",
			"G$",
			"$"}}

	for index := 0; index < len(inputs); index++ {
		input := inputs[index]

		expected := expectedOutputs[index]

		actual := constructSuffixTree(input)
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
