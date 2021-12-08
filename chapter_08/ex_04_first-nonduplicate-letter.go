// Write a function that returns the first non-duplicated character in a string. For example, the string, "minimum" has two characters that only exist once -- the "n" and the "u" -- so your function should return the "n", since it occurs first. The function should have an efficiency of O(N).

package main

import (
	"fmt"
)

func firstNonDuplicate(phrase string) string {
	hashTable := make(map[string]int)
	for _, val := range []byte(phrase) {
		hashTable[string(val)] += 1
	}

	for _, val := range []byte(phrase) {
		if hashTable[string(val)] == 1 {
			return string(val)
		}
	}

	return "no non-duplicates"
}

func main() {
	phrase := "minimum"

	fmt.Println(firstNonDuplicate(phrase))
}
