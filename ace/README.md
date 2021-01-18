
Pre-requisite:
- Install golang per the instructions in https://golang.org/doc/install
- Set environment variable GOPATH to the root of the go workspace dir, i.e.
  GOPATH = $HOME/go
- mkdir $GOPATH/src

 
Install dir:
> cd $GOPATH/src
> tar xvf LinoPereira_ace.tar

To run the tests:
> cd $GOPATH/src/ace/trie
> go test

To build it:
> cd $GOPATH/src/ace
> go install

To run it:
> $GOPATH/bin/ace -f <data_file_path>
Then, enter: 'complete,<prefix>,<max_count>' or 'quit'

Example:
> cd $GOPATH/src/ace
> $GOPATH/bin/ace -f basicTestData.txt
help
*** Unrecognized - Enter 'complete,<prefix>,<max_count>' or 'quit'
complete,c,0
cat,c,car,card,cork
complete,ba,0
bathing,bat,bath,bat ting
quit

Design choices:
Then, enter: 'complete,<prefix>,<max_count>' or 'quit'
- Chose a Trie data structure because it is designed for this. It can quickly answer
  the query of which words/phrases start with a given prefix, as is the case here.
- Chose to perform sort on the result data. This approach works well where
  the result sets are typically short and varied.
  For larger result sets and more repetitive queries, a cache could be
  added to allow quick access to the n most recent query results, although such
  cache would grow the memory footprint of the program.
- Built trie_test to allow the easy addition of both more word/phrase data and
  more queries. With more time, I'd break up TestTrie into more smaller functions.
- Also, due to time constraints, tests were only written against Trie, although
  they do provide 100% coverage.
