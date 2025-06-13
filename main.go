package main

//func main() {
//	result := constructSuffixTree("ATAAATG$")
//
//	fmt.Println("", result)
//}

import (
	"fmt"
)

// SuffixTreeNode represents a node in the suffix tree.
type SuffixTreeNode struct {
	children map[rune]*SuffixTreeNode
	indexes  []int // Store the starting indexes of suffixes ending at this node
}

// NewSuffixTreeNode creates a new SuffixTreeNode.
func NewSuffixTreeNode() *SuffixTreeNode {
	return &SuffixTreeNode{
		children: make(map[rune]*SuffixTreeNode),
		indexes:  make([]int, 0),
	}
}

// SuffixTree represents the suffix tree itself.
type SuffixTree struct {
	root *SuffixTreeNode
	text string
}

// NewSuffixTree creates a suffix tree for a given string.
func NewSuffixTree(text string) *SuffixTree {
	tree := &SuffixTree{
		root: NewSuffixTreeNode(),
		text: text,
	}
	tree.build()
	return tree
}

// build constructs the suffix tree using a naive O(n^2) algorithm.
func (tree *SuffixTree) build() {
	for i := 0; i < len(tree.text); i++ {
		current := tree.root
		for _, ch := range tree.text[i:] {
			if _, ok := current.children[ch]; !ok {
				current.children[ch] = NewSuffixTreeNode()
			}
			current = current.children[ch]
			current.indexes = append(current.indexes, i)
		}
	}
}

// Print prints all the suffixes in the tree.
func (tree *SuffixTree) Print() {

	// define function variable!!!
	var dfs func(node *SuffixTreeNode, prefix string)
	// function implementation
	dfs = func(node *SuffixTreeNode, prefix string) {
		if len(node.children) == 0 {
			fmt.Println(prefix)
		}
		for ch, child := range node.children {
			dfs(child, prefix+string(ch))
		}
	}

	dfs(tree.root, "")
}

func main() {
	text := "ACA$"
	tree := NewSuffixTree(text)
	//fmt.Println("Suffixes in the Suffix Tree:")
	tree.Print()
}
