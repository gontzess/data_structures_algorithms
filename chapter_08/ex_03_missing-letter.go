// Write a function that accepts a string that contains all the letters of the alphabet except one and returns the missing letter. For example, the string, "the quick brown box jumps over a lazy dog" contains all the letters of the alphabet except the letter, "f". The function should have a time complexity of O(N).

package main

import (
	"fmt"
	"strings"
)

// assume only lowercase input for simplicity
func missingLetter(phrase string) string {
	hashTable := make(map[string]bool)

	for _, letter := range []byte(phrase) {
		hashTable[string(letter)] = true
	}

	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	for _, letter := range alphabet {
		if hashTable[letter] == false {
			return letter
		}
	}

	return "no missing"
}

func main() {
	phrase := "the quick brown box jumps over a lazy dog"

	fmt.Println(missingLetter(phrase))
}
