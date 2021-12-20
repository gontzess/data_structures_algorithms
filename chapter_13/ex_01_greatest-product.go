// Given an array of positive numbers, write a function that returns the greatest product of any three numbers. The approach of using three nested loops would clock in at O(N3), which is very slow. Use sorting to implement the function in a way that it computes at O(N log N) speed. (There are even faster implementations, but we're focusing on using sorting as a technique to make code faster.)

package main

import (
	"fmt"
)

func greatestProduct(sli []int) int {
	if len(sli) < 3 {
		return 0
	}

	sortableSli := SortableSlice{slice: sli}
	lastIdx := len(sortableSli.slice) - 1
	sortableSli.quickSort(0, lastIdx)

	return sortableSli.slice[lastIdx] * sortableSli.slice[lastIdx-1] * sortableSli.slice[lastIdx-2]
}

func main() {
	sli := []int{5, 9, 3, 2, 4, 5, 6}
	fmt.Println(greatestProduct(sli))
}
