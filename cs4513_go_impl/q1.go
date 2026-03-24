package cs4513_go_impl

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

/*
Find the top K most common words in a text document.
What is a word?

	A word here only consists of alphanumeric characters, e.g., catch21
	All punctuations and other characters should be removed, e.g. "don't" becomes "dont" or "end." becomes "end"; done before the charThreshold
	A word has to satifies the charThreshold, e.g., if charThreshold = 5  "apple" is a word, but neither "new" or "york" are words

Matching condition

	Matching is case insensitive

Parameters:
- path: file path
- numWords: number of words to return (i.e. k)
- charThreshold: threshold for whether a token qualifies as a word
You should use `checkError` to handle potential errors.
*/
func TopWords(path string, numWords int, charThreshold int) []WordCount {
	// TODO: implement me
	// HINT: You may find the `strings.Fields` and `strings.ToLower` functions helpful.
	// HINT: the regex "[^0-9a-zA-Z]+" can be used to spot any non-alphanumeric characters.

	// check to make sure there's no words
	if numWords <= 0 {
		return []WordCount{}
	}

	// read the file
	fileData, err := os.ReadFile(path)
	checkError(err)

	// convert file to string & lower case
	fileContent := string(fileData)
	fileContent = strings.ToLower(fileContent)

	// remove all non alphanumeric characters
	contentsRe := regexp.MustCompile("[^0-9a-zA-Z]+") //returns a regexp type
	counts := make(map[string]int)                    //map to store words

	// loop over and if it's not meeting the charThreshold, then get rid of it
	// contentsRe.ReplaceAllString(fileContent, " "): every chunk of non alphanumeric character gets replaced w a space
	// strings.Fields(...): splits the string into words
	for _, value := range strings.Fields(fileContent) {
		fileCleaned := contentsRe.ReplaceAllString(value, "")
		if len(fileCleaned) >= charThreshold {
			// look in the map with "fileCleaned" as the key
			// get the integer there & increase it by one
			counts[fileCleaned]++
		}
	}

	// now we need to sort it
	// first turn map into slice of WordCount
	var results []WordCount
	for word, count := range counts {
		results = append(results, WordCount{
			Word:  word,
			Count: count,
		})
	}

	// sort
	sortWordCounts(results)

	// avoid slicing past the end
	if numWords > len(results) {
		numWords = len(results)
	}

	return results[:numWords]
}

/*
Do NOT modify this struct!
A struct that represents how many times a word is observed in a document
*/
type WordCount struct {
	Word  string
	Count int
}

/*
Do NOT modify this function!
*/
func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

/*
Do NOT modify this function!
Helper function to sort a list of word counts in place.
This sorts by the count in decreasing order, breaking ties using the word.
*/
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}
