package trie

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

// test input data
var testWordList = []WordData{
	{"cat", 1},
	{"car", 1},
	{"Bath", 4},
	{"BAT", 3},
	{"rat", 1},
	{"c", 2},
	{"cork", 1},
	{"", 1},
	{"BATHING", 5},
	{"CARD", 1},
	{"Cat chased the dog", 1},
}
// test prefixes
var testPrefixes = []string{
	"ca", "ba", "Bat", "c", "cor", "cork", "be", "t", "", "bad bat", " car", "cat "}
// expectedResults the computed expected results for the testPrefixes, given the
// testWordList as input.
var expectedResults map[string][]WordData


//==============================================================

// TestTrie tests loading and querying the Trie
func TestTrie(t *testing.T) {

	//  Create the trie and add all test words
	tPtr := NewTrie()
	for _, wd := range testWordList {
		for i := 0; i < wd.Count; i++ {
			tPtr.AddWord(wd.Word)
		}
	}
	tPtr.PrintAll()
	fmt.Println()

	// setup test data to compare results against
	for i, w := range testWordList {
		testWordList[i].Word = strings.ToLower(w.Word)
	}
	sort.SliceStable(testWordList, func(i, j int) bool {
		return testWordList[j].Count < testWordList[i].Count
	})
	dedupedTestWordCount := getExpectedCount(testWordList)
	expectedResults = buildExpectedResults()

	// verify that all expected words were added
	// Could also compare each list item against the other
	allWords := tPtr.GetAllWords()
	if len(allWords) != dedupedTestWordCount {
		t.Error("Expected len of allWords", len(allWords),
			"and dedupped testWordList", dedupedTestWordCount, "to be equal")
	}

	// test FindWordsFor and FindWordsForAsString for eachtest data prefix
	for _, prefix := range testPrefixes {
		foundWords := tPtr.GetWordsFor(prefix)
		foundWordsAsString := tPtr.AutoComplete(prefix, 0)
		foundWordsAsStringList := strings.Split(foundWordsAsString, ",")
		if !(len(foundWords) == 0 && len(foundWordsAsStringList) == 1) &&
			len(foundWords) != len(foundWordsAsStringList) {
			t.Error("Expected len of foundWords", len(foundWords), "and foundWordsAsStringList", len(foundWordsAsStringList), "to be equal")
		}
		fmt.Println("FOUND:", prefix, ":", foundWords)
		expectedWords, ok := expectedResults[prefix]
		if !ok {
			// Filter out empty prefix values
			continue
		}
		// verify that counts match between found and expected
		if len(foundWords) != len(expectedWords) {
			t.Error("Expected len of foundWords", len(foundWords), "and expectedWords", len(expectedWords), "to be equal")
			t.Error("foundWords:", foundWords)
			t.Error("expectedWords:", expectedWords)
		}
		// Verify that word list is in expected order
		for i := 0; i < len(foundWords); i++ {
			// Note that found and expected words may not match, but counts must be the same
			if foundWords[i].Count != expectedWords[i].Count {
				t.Error("Expected found", foundWords[i], "and expected word", expectedWords[i], "at i:", i, "to have same count")
			}
		}
	}

	// test FindWordsForAsString with count=1 for eachtest data prefix
	for _, prefix := range testPrefixes {
		foundWords := tPtr.GetWordsFor(prefix)
		oneWordAsString := tPtr.AutoComplete(prefix, 1)
		oneWordAsStringList := strings.Split(oneWordAsString, ",")
		if (len(foundWords) > 0 && len(oneWordAsStringList) != 1) &&
			!(len(foundWords) == 0 && len(oneWordAsStringList) == 1) {
			t.Error("Expected len of foundWords", len(foundWords), "and oneWordAsStringList", len(oneWordAsStringList), "to be equal")
		}
	}
}

//==============================================================

// buildExpectedResults build the expected results
func buildExpectedResults() map[string][]WordData {
	expectedResults := make(map[string][]WordData)
	for _, prefix := range testPrefixes {
		if len(prefix) > 0 {
			expectedResults[prefix] = getExpectedWordsFor(strings.ToLower(strings.TrimLeft(prefix, " ")))
		}
	}
	return expectedResults
}

// printExpectedResults for debug only
func printExpectedResults() {
	for p, wdList := range expectedResults {
		fmt.Println("Expected:", p, "-", wdList)
	}
}

// getExpectedWordsFor computes the expected test data for prefix
func getExpectedWordsFor(prefix string) []WordData {
	var words []WordData
	if len(prefix) > 0 {
		// dedupe
		prevWord := ""
		for _, wd := range testWordList {
			word := strings.ToLower(wd.Word)
			if word != prevWord && strings.HasPrefix(word, prefix) {
				words = append(words, WordData{word, wd.Count})
			}
			prevWord = word
		}
	}
	return words
}

// getExpectedCount filters out empty test data, since it won't be added to Trie
func getExpectedCount(words []WordData) int {
	count := 0
	for _, wd := range words {
		if len(wd.Word) > 0 {
			count = count + 1
		}
	}
	return count
}
