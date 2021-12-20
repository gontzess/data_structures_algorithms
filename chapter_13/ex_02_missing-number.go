// The following function finds the "missing number" from an array of integers. That is, the array is expected to have all integers from 0 up to the array's length, but one is missing. As examples, the array, [5, 2, 4, 1, 0] is missing the number 3, and the array, [9, 3, 2, 5, 6, 7, 1, 0, 4] is missing the number 8.
// The given function runs is O(N^2). Use sorting to write a new implementation of the one given so it only takes O(N log N). (There are even faster implementations, but we're focusing on using sorting as a technique to make code faster.)

package main

import "fmt"

func findMissingNumber(sli []int) int {
	if len(sli) < 1 {
		return 0
	}

	sortableSli := SortableSlice{slice: sli}
	sortableSli.quickSort(0, len(sli)-1)

	var missing int
	for idx := 0; idx < len(sli); idx++ {
		if sortableSli.slice[idx] != idx {
			missing = idx
			break
		}
	}

	return missing
}

func main() {
	// sli1 := []int{5, 2, 4, 1, 0}
	// fmt.Println(findMissingNumber(sli1)) // 3

	sli2 := []int{9, 3, 2, 5, 6, 7, 1, 0, 4}
	fmt.Println(findMissingNumber(sli2)) // 8

	sli3 := []int{2, 7, 0, 5, 9, 6, 1, 4, 8}
	fmt.Println(findMissingNumber(sli3)) // 3

	fmt.Println(findMissingNumber([]int{5, 2, 4, 0, 1})) // 3
}
