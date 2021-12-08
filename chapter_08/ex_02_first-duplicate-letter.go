// Write a function that accepts an array of strings and returns the first duplicate value it finds. For example, if the array is ["a", "b", "c", "d", "c", "e", "f"], the function should return "c", since it's duplicated within the array. (You can assume that there's one pair of duplicates within the array.) Make sure the function has an efficiency of O(N).

package main

import (
	"fmt"
)

func firstDuplicate(slice []string) string {
	hashTable := make(map[string]bool)
	for _, val := range slice {
		if hashTable[val] == true {
			return val
		}
		hashTable[val] = true
	}

	return "no duplicates"
}

func main() {
	slice := []string{"a", "b", "c", "d", "c", "e", "f"}

	fmt.Println(firstDuplicate(slice))
}
