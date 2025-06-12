package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestTrieMatching(t *testing.T) {
	var inputs [][]string = [][]string{
		{"AAA",
			"1",
			"AA"},
		{
			"AA",
			"1",
			"T"},
		{
			"AATCGGGTTCAATCGGGGT",
			"2",
			"ATCG",
			"GGGT"}}
	var exptedOutput [][]int = [][]int{
		{0, 1},
		{},
		{1, 4, 11, 15}}

	for index := 0; index < len(inputs); index++ {
		input := inputs[index]
		expected := exptedOutput[index]

		text := input[0]
		actual := []int{}
		for i := 2; i < len(input); i++ {
			trie := input[i]
			localActual := trieMatching(text, trie)
			// join two arrays
			for j := 0; j < len(localActual); j++ {
				actual = append(actual, localActual[j])
			}
		}

		slices.Sort(actual)
		slices.Sort(expected)

		if len(actual) > 0 && len(expected) > 0 {
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf(`Test %q trieMatching:actual = %q`, index, actual)
				t.Errorf(`Test %q trieMatching:expected = %q`, index, expected)
			}
		}
	}

}
