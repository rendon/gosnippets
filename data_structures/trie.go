// Basic Trie implementation
// Given a list of words, a trie allows us to look up words in O(|w|) time,
// where |w| is the length of the word to find, in other words, very fast!
package main

import (
	"fmt"
)

type Node struct {
	Char  byte
	Next  map[byte]*Node
	Value interface{}
}

type Trie struct {
	Root *Node
}

func add(node *Node, word string, p int, value interface{}) bool {
	if node == nil {
		return false
	}
	if p == len(word) {
		node.Value = value
		return false
	}
	isNew := false
	char := word[p]
	if node.Next[char] == nil {
		node.Next[char] = &Node{
			Char: char,
			Next: make(map[byte]*Node),
		}
		isNew = true
	}
	return add(node.Next[char], word, p+1, value) || isNew
}

func (t *Trie) AddWord(word string, value interface{}) bool {
	return add(t.Root, word, 0, value)
}

func find(node *Node, word string, p int) interface{} {
	if node == nil {
		return nil
	}
	if p == len(word) {
		return node.Value
	}
	char := word[p]
	if node.Next[char] == nil {
		return nil
	}
	return find(node.Next[char], word, p+1)
}

func (t *Trie) FindWord(word string) interface{} {
	return find(t.Root, word, 0)
}

func NewTrie() *Trie {
	return &Trie{
		Root: &Node{
			Next: make(map[byte]*Node),
		},
	}
}

func main() {
	words := []string{
		"abc",
		"foo",
		"bar",
		"xyz",
	}

	t := NewTrie()
	for i, w := range words {
		t.AddWord(w, i)
	}
	fmt.Printf("%v\n", t.FindWord("abc"))
	fmt.Printf("%v\n", t.FindWord("word"))
	fmt.Printf("%v\n", t.FindWord("xyz"))
	fmt.Printf("%v\n", t.FindWord("foo"))
	fmt.Printf("%v\n", t.FindWord("fo"))
}
