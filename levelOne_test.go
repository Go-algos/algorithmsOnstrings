package main

import (
	"testing"
	"reflect"
)

func TestConstructTrie(t *testing.T) {
	var input []string = []string{"1", "AA"}
	var output []string = []string{"0->1:A", "2->3:A", "1->2:T"}
	var result = constructTrie(input)
	//if result != output {
	if !reflect.DeepEqual(result, output) {
		t.Errorf(`constructTrie() = %q, %v`, result, output)
	}
}
