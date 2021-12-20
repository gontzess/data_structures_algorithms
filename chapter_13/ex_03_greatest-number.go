// Write three different implementations of a function that finds the greatest number within an array. Write one function that is O(N2), one that is O(N log N), and one that is O(N).

package main

import "fmt"

// O(N^2)
func findGreatestNumber1(sli []int) int {
	if len(sli) < 1 {
		return 0
	}

	var isGreatest bool
	var greatest int
	for idxI := 0; idxI < len(sli); idxI++ {
		isGreatest = true
		for idxJ := 0; idxJ < len(sli); idxJ++ {
			if sli[idxJ] > sli[idxI] {
				isGreatest = false
			}
		}

		if isGreatest == true {
			greatest = sli[idxI]
			break
		}
	}

	return greatest
}

// O(N log N)
func findGreatestNumber2(sli []int) int {
	if len(sli) < 1 {
		return 0
	}

	sortableSli := SortableSlice{slice: sli}
	lastIdx := len(sli) - 1
	sortableSli.quickSort(0, lastIdx)

	return sortableSli.slice[lastIdx]
}

// O(N)
func findGreatestNumber3(sli []int) int {
	if len(sli) < 1 {
		return 0
	}

	greatest := sli[0]
	for idx := 1; idx < len(sli); idx++ {
		if sli[idx] > greatest {
			greatest = sli[idx]
		}
	}

	return greatest
}

func main() {
	sli := []int{3, 2, 5, 6, 7, 9, 1, 0, 4}
	fmt.Println(findGreatestNumber1(sli)) // 9
	fmt.Println(findGreatestNumber2(sli)) // 9
	fmt.Println(findGreatestNumber3(sli)) // 9
}
