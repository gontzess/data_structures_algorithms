// Write a function that returns the intersection of two arrays. The intersection is a third array that contains all values contained within the first two arrays. For example, the intersection of [1, 2, 3, 4, 5] and [0, 2, 4, 6, 8] is [2, 4]. Your function should have a complexity of O(N). (If your programming language has a built-in way of doing this, don't use it. The idea is to build the algorithm yourself.)

package main

import (
	"fmt"
)

func intersection(slice1 []int, slice2 []int) []int {
	var largerSlice, smallerSlice []int

	if len(slice1) >= len(slice2) {
		largerSlice, smallerSlice = slice1, slice2
	} else {
		largerSlice, smallerSlice = slice2, slice1
	}

	hashTable := make(map[int]bool, len(largerSlice))
	for _, val := range largerSlice {
		hashTable[val] = true
	}

	var intersection []int
	for _, val := range smallerSlice {
		if hashTable[val] == true {
			intersection = append(intersection, val)
		}
	}

	return intersection
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{0, 2, 4, 6, 8}

	fmt.Println(intersection(slice1, slice2))
}
