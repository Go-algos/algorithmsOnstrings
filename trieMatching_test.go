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
			"AA"}}
	var exptedOutput [][]int = [][]int{
		{0, 1}}

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

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf(`Test %q structTrie() = %q`, index, actual)
			t.Errorf(`Test %q constructTrie() = %q`, index, expected)
		}
	}

}
