package trie

import (
	"fmt"
	"sort"
)

// WordData used to return found words and their word count
type WordData struct {
	Word  string
	Count int
}

// trieNode structure used by the Trie
// Chars is a map with a character as the key and trieNode as the value.
// Count keeps count of the number of words that end in this trieNode.
// IsEnd is true when a word ends in this trieNode, false otherwise.
type trieNode struct {
	Chars map[rune]*trieNode
	Count int
	IsEnd bool
}

// NewTrie Constructor for a Trie
func newTrieNode() *trieNode {
	return &trieNode{make(map[rune]*trieNode), 0, false}
}

// addChar adds char to this trieNode
func (n *trieNode) addChar(char rune) *trieNode {
	childPtr, ok := n.Chars[char]
	if !ok {
		childPtr = newTrieNode()
		n.Chars[char] = childPtr
	}
	return childPtr
}

// setIsEnd sets the word ending in this trieNode and increments the count
func (n *trieNode) setIsEnd(end bool) {
	n.IsEnd = end
	if end {
		n.Count = n.Count + 1
	}
}

// addWord adds a word to this trieNode
func (r *trieNode) addWord(word string) {
	if len(word) > 0 {
		nodePtr := r
		for i := 0; i < len(word); i++ {
			c := word[i]
			nodePtr = nodePtr.addChar(rune(c))
		}
		nodePtr.setIsEnd(true)
	}
}

// findChar returns the child trieNode for char, if found
func (n *trieNode) findChar(char rune) (*trieNode, bool) {
	nodePtr, ok := n.Chars[char]
	return nodePtr, ok
}

// getWordsFor returns a list of words and word counts for prefix from root trieNode.
// Walks trieNodes for each prefix character to find starting trieNode for
// all child words, then calls getWords on it.
func (r *trieNode) getWordsFor(prefix string) []WordData {
	var wordDataList []WordData
	if len(prefix) > 0 {
		nodePtr := r
		for i := 0; i < len(prefix); i++ {
			char := prefix[i]
			childPtr, ok := nodePtr.findChar(rune(char))
			if !ok {
				return wordDataList
			}
			nodePtr = childPtr
		}
		wordDataList = nodePtr.getWords(prefix)
	}
	// sort results by count in decrementing order
	sort.SliceStable(wordDataList, func(i, j int) bool {
		return wordDataList[j].Count < wordDataList[i].Count
	})
	return wordDataList
}

// getWords returns a list of words and word counts for prefix from this trieNode
// Starting with this trieNode, gets all child words by recursing down all
// child trieNodes until it finds each word end.
func (n *trieNode) getWords(prefix string) []WordData {
	var wordDataList []WordData
	if n.IsEnd {
		wordDataList = append(wordDataList, WordData{prefix, n.Count})
	}
	for k, v := range n.Chars {
		moreWords := v.getWords(prefix + string(k))
		for _, wd := range moreWords {
			wordDataList = append(wordDataList, WordData{wd.Word, wd.Count})
		}
	}
	return wordDataList
}

// getAll returns all words, for testing and debugging
func (n *trieNode) getAll() []string {
	wordData := n.getWords("")
	var words []string
	for _, wd := range wordData {
		words = append(words, wd.Word)
	}
	return words
}

// printAll prints info about each word in this trieNode, for testing and debugging
func (n *trieNode) printAll(prefix string) {
	for k, v := range n.Chars {
		if v.IsEnd {
			fmt.Printf("%s%s, child: %d, count: %d, isEnd: %v\n", prefix, string(k), len(v.Chars), v.Count, v.IsEnd)
		}
		v.printAll(prefix + string(k))
	}
}
