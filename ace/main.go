package main

import (
	"public-go/ace/trie"
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// autoCompleteEngine pointer to the Trie instance
var autoCompleteEngine *trie.Trie
// isAlphabetic generated MatchString for this regular expression.
// Returns true only if the test string contains letters (A-Z,a-z) and spaces.
var isAlphabetic = regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString
// queryStringformat string constant
const queryStringformat = "complete,<prefix>,<max_count>"

// usage prints usage info and exits
func usage() {
	fmt.Printf("Usage of %s\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

// loadFile validate and load file content into the autoCompleteEngine
func loadFile(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if isAlphabetic(scanner.Text()) {
			autoCompleteEngine.AddWord(strings.Trim(scanner.Text(), " "))
		} else {
			errorMsg := fmt.Sprintf("Error: Invalid non-alpabetic input found in phrase: [%s]", scanner.Text())
			panic(errorMsg)
		}
	}
}

// processInput parse and process the input, and return the prefix, count and
// an ok flag that is true when the input is valid, false otherwise.
func processInput(input string) (string, int, bool) {
	inList := strings.Split(input, ",")
	switch strings.ToLower(inList[0]) {
	case "complete":
		if len(inList) != 3 {
			return "", 0, false
		}
		prefix := strings.TrimLeft(inList[1], " ")
		count, err := strconv.ParseInt(strings.Trim(inList[2], " "), 10, 64)
		if err != nil {
			return "", 0, false
		}
		return prefix, int(count), true
	case "quit":
		os.Exit(0)
	}
	return "", 0, false
}

//main parse args,start and load the autoCompleteEngine and then processes input
func main() {
	// Parse and validate the args
	filePathPtr := flag.String("f", "", "The name of the file containing the list of phrases")
	flag.Parse()
	if *filePathPtr == "" {
		usage()
	}

	// Start and load the autoCompleteEngine
	autoCompleteEngine = trie.NewTrie()
	loadFile(*filePathPtr)

	// Start processing input from sdtin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			prefix, count, ok := processInput(scanner.Text())
			if !ok {
				fmt.Printf("*** Unrecognized - Enter '%s' or 'quit'\n", queryStringformat)
			} else {
				fmt.Println(autoCompleteEngine.AutoComplete(prefix, count))
			}
		}
	}
}
