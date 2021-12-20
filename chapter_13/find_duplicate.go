package main

import (
	"fmt"
)

func hasDuplicateValue(sli []int) bool {
	sortableSli := SortableSlice{slice: sli}
	sortableSli.quickSort(0, len(sortableSli.slice)-1)

	for idx := 0; idx < len(sortableSli.slice); idx++ {
		if sortableSli.slice[idx] == sortableSli.slice[idx+1] {
			return true
		}
	}

	return false
}

func main() {
	sli := []int{5, 9, 3, 2, 4, 5, 6}
	fmt.Println(hasDuplicateValue(sli))
}
