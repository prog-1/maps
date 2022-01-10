package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// UniqueRunes returns distinct runes and their count.
func UniqueRunes(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	return m
}

func main() {
	fmt.Print("Enter a string: ")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runeCount := UniqueRunes(s)
	for r, c := range runeCount {
		if unicode.IsPrint(r) {
			fmt.Printf("%U (%c) : %d\n", r, r, c)
		} else {
			fmt.Printf("%U (▨) : %d\n", r, c)
		}
	}
}
