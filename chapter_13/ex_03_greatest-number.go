// Write three different implementations of a function that finds the greatest number within an array. Write one function that is O(N2), one that is O(N log N), and one that is O(N).

package main

import "fmt"

// O(N^2)
func findGreatestNumber1(sortableSli SortableSlice) int {
	if len(sortableSli) < 1 {
		return 0
	}

	var isGreatest bool
	var greatest int
	for idxI := 0; idxI < len(sortableSli); idxI++ {
		isGreatest = true
		for idxJ := 0; idxJ < len(sortableSli); idxJ++ {
			if sortableSli[idxJ] > sortableSli[idxI] {
				isGreatest = false
			}
		}

		if isGreatest == true {
			greatest = sortableSli[idxI]
			break
		}
	}

	return greatest
}

// O(N log N)
func findGreatestNumber2(sortableSli SortableSlice) int {
	if len(sortableSli) < 1 {
		return 0
	}
	lastIdx := len(sortableSli) - 1
	sortableSli.quickSort(0, lastIdx)

	return sortableSli[lastIdx]
}

// O(N)
func findGreatestNumber3(sortableSli SortableSlice) int {
	if len(sortableSli) < 1 {
		return 0
	}

	greatest := sortableSli[0]
	for idx := 1; idx < len(sortableSli); idx++ {
		if sortableSli[idx] > greatest {
			greatest = sortableSli[idx]
		}
	}

	return greatest
}

func main() {
	fmt.Println(findGreatestNumber1(SortableSlice{3, 2, 5, 6, 7, 9, 1, 0, 4})) // 9
	fmt.Println(findGreatestNumber2(SortableSlice{3, 2, 5, 6, 7, 9, 1, 0, 4})) // 9
	fmt.Println(findGreatestNumber3(SortableSlice{3, 2, 5, 6, 7, 9, 1, 0, 4})) // 9
}
