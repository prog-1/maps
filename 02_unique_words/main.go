package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

// UniqueWords returns distinct words and their count. Words consist of letters
// only (unicode.IsLetter must return true).
func UniqueWords(s string) map[string]int {
	m := make(map[string]int)
	start := -1 // First byte of a word. Negative while skipping separators.
	for i, r := range s {
		if unicode.IsLetter(r) {
			if start < 0 { // First letter of a word.
				start = i
			}
		} else if start >= 0 { // First separator.
			m[s[start:i]]++
			start = -1
		}
	}
	if start >= 0 {
		m[s[start:]]++
	}
	return m
}

func main() {
	fmt.Print("Enter a string: ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	wordCount := UniqueWords(s)
	// Output in alphabetical order.
	words := make([]string, 0, len(wordCount))
	for w := range wordCount {
		words = append(words, w)
	}
	sort.Strings(words)
	for _, w := range words {
		fmt.Println(w, wordCount[w])
	}
}
