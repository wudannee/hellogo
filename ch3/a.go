// Package stuffs
package ch3

import (
	"fmt"
	"unicode/utf8"
)

// Greet returns a greeting message for the given name.
func Greet(name string) string {
	return "Hello " + name
}

// CountTotalWords returns the number of words in a given string.
func CountTotalWords(s string) int {
	fmt.Println(s)
	total := utf8.RuneCountInString(s)
	return total
}

// CountWordsInDetails returns a map of words and their respective counts in the given string.
// The keys of the map are the words, and the values are the counts.
func CountWordsInDetails(s string) (wordsMap *map[string]int) {
	wordsMap = &map[string]int{}
	for _, word := range s {
		(*wordsMap)[string(word)]++
	}
	return
}

func DemoWords() {
	msg := "你好golang, 你好世界"
	n := CountTotalWords(msg)
	fmt.Println("total number of words:", n)

	mapWords := CountWordsInDetails(msg)
	for k, v := range *mapWords {
		fmt.Printf("%s: %d\n", k, v)
	}
}
