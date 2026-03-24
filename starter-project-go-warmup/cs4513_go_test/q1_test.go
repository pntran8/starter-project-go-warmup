// Package cs4513_go_test includes test cases for the top-k word count and parallel sum
package cs4513_go_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	// this import path denotes the directory containing Go source files for the package cs4513_go_impl
	// further, the import declaration binds a short name (by default package name cs4513_go_impl) to the imported package; the short name will be used to refer to package content
	"starter-project-go-warmup/cs4513_go_impl"
)

func equal(counts1, counts2 []cs4513_go_impl.WordCount) bool {
	if len(counts1) != len(counts2) {
		return false
	}
	for i := range counts1 {
		if counts1[i] != counts2[i] {
			return false
		}
	}
	return true
}

func assertEqual(t *testing.T, answer, expected []cs4513_go_impl.WordCount) {
	if !equal(answer, expected) {
		t.Fatal(fmt.Sprintf(
			"Word counts did not match...\nExpected: %v\nActual: %v",
			expected,
			answer))
	}
}

func writeTempFile(t *testing.T, content string) string {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "input.txt")

	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	return path
}

func TestTopWordsNegativeNumWords(t *testing.T) {
	path := writeTempFile(t, "apple banana apple")

	got := cs4513_go_impl.TopWords(path, -1, 1)
	want := []cs4513_go_impl.WordCount{}

	assertEqual(t, got, want)
}

func TestTopWordsBasicCounting(t *testing.T) {
	path := writeTempFile(t, "apple banana apple carrot banana apple")

	got := cs4513_go_impl.TopWords(path, 2, 1)
	// keyed struct literal
	want := []cs4513_go_impl.WordCount{
		{Word: "apple", Count: 3},
		{Word: "banana", Count: 2},
	}

	assertEqual(t, got, want)
}

func TestTopWordsCaseInsensitive(t *testing.T) {
	path := writeTempFile(t, "Apple apple APPLE BaNaNa banana")

	got := cs4513_go_impl.TopWords(path, 2, 1)
	want := []cs4513_go_impl.WordCount{
		{Word: "apple", Count: 3},
		{Word: "banana", Count: 2},
	}

	assertEqual(t, got, want)
}

func TestTopWordsRemovesPunctuation(t *testing.T) {
	path := writeTempFile(t, "don't go. dont, GO!!!")

	got := cs4513_go_impl.TopWords(path, 2, 1)
	want := []cs4513_go_impl.WordCount{
		{Word: "dont", Count: 2},
		{Word: "go", Count: 2},
	}

	assertEqual(t, got, want)
}

func TestTopWordsNumWordsLargerThanUniqueWords(t *testing.T) {
	path := writeTempFile(t, "apple banana apple")

	got := cs4513_go_impl.TopWords(path, 10, 1)
	want := []cs4513_go_impl.WordCount{
		{Word: "apple", Count: 2},
		{Word: "banana", Count: 1},
	}
	assertEqual(t, got, want)

}

func TestTopWordsAppliesThresholdAfterCleaning(t *testing.T) {
	path := writeTempFile(t, "apple end. kiwi new york")

	got := cs4513_go_impl.TopWords(path, 3, 5)
	want := []cs4513_go_impl.WordCount{
		{Word: "apple", Count: 1},
	}

	assertEqual(t, got, want)
}

func TestTopWordsTieBreaksAlphabetically(t *testing.T) {
	path := writeTempFile(t, "banana apple carrot")

	got := cs4513_go_impl.TopWords(path, 3, 1)
	want := []cs4513_go_impl.WordCount{
		{Word: "apple", Count: 1},
		{Word: "banana", Count: 1},
		{Word: "carrot", Count: 1},
	}

	assertEqual(t, got, want)
}

func TestSimple(t *testing.T) {
	// top 4; any words
	answer1 := cs4513_go_impl.TopWords("simple.txt", 4, 0)
	// top 5; len(word) >=4
	answer2 := cs4513_go_impl.TopWords("simple.txt", 5, 4)
	// top 5; len(word) >= 5
	answer3 := cs4513_go_impl.TopWords("simple.txt", 2, 5)

	expected1 := []cs4513_go_impl.WordCount{
		{"hello", 5},
		{"you", 3},
		{"and", 2},
		{"dont", 2},
	}
	expected2 := []cs4513_go_impl.WordCount{
		{"hello", 5},
		{"dont", 2},
		{"everyone", 2},
		{"look", 2},
		{"again", 1},
	}
	expected3 := []cs4513_go_impl.WordCount{
		{"hello", 5},
		{"everyone", 2},
	}
	assertEqual(t, answer1, expected1)
	assertEqual(t, answer2, expected2)
	assertEqual(t, answer3, expected3)
}

func TestDeclarationOfIndependence(t *testing.T) {
	answer1 := cs4513_go_impl.TopWords("declaration_of_independence.txt", 5, 6)

	answer2 := cs4513_go_impl.TopWords("declaration_of_independence.txt", 4, 2)

	expected1 := []cs4513_go_impl.WordCount{
		{"people", 10},
		{"states", 8},
		{"government", 6},
		{"powers", 5},
		{"assent", 4},
	}

	expected2 := []cs4513_go_impl.WordCount{
		{"of", 80},
		{"the", 78},
		{"to", 65},
		{"and", 57},
	}

	assertEqual(t, answer1, expected1)
	assertEqual(t, answer2, expected2)
}

func TestDracula(t *testing.T) {

	answer := cs4513_go_impl.TopWords("pg-dracula.txt", 5, 8)

	expected := []cs4513_go_impl.WordCount{
		{"jonathan", 187},
		{"professor", 156},
		{"something", 139},
		{"anything", 99},
		{"terrible", 97},
	}

	assertEqual(t, answer, expected)
}
