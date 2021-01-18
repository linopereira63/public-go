package trie

import (
	"fmt"
	"strings"
)

// A Trie data type
type Trie struct {
	root *trieNode
}

// NewTrie Constructor for a Trie
func NewTrie() *Trie {
	return &Trie{newTrieNode()}
}

// AddWord adds word to this Trie
func (t *Trie) AddWord(word string) {
	t.root.addWord(strings.ToLower(word))
}

// GetWordsFor returns a list of words and word counts for prefix
func (t *Trie) GetWordsFor(prefix string) []WordData {
	return t.root.getWordsFor(strings.ToLower(strings.TrimLeft(prefix, " ")))
}

// AutoComplete returns a comma separated
func (t *Trie) AutoComplete(prefix string, wordCount int) string {
	wordDataList := t.GetWordsFor(prefix)
	result := ""
	if len(wordDataList) > 0 {
		var resultBuilder strings.Builder
		// wordCount=0 means get all
		keepCount := wordCount > 0
		count := 0
		for _, wd := range wordDataList {
			fmt.Fprintf(&resultBuilder, "%s,", wd.Word)
			count = count + 1
			if keepCount && count == wordCount {
				break
			}
		}
		result = resultBuilder.String()
		// remove the trailing ","
		result = result[:resultBuilder.Len()-1]
	}
	return result
}

// GetAllWords returns all words, for testing and debugging
func (t *Trie) GetAllWords() []string {
	return t.root.getAll()
}

// PrintAll prints info about each word in this Trie, for testing and debugging
func (t *Trie) PrintAll() {
	t.root.printAll("")
}
