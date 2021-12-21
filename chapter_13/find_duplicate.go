package main

import (
	"fmt"
)

func hasDuplicateValue(sortableSli SortableSlice) bool {
	sortableSli.quickSort(0, len(sortableSli)-1)

	for idx := 0; idx < len(sortableSli); idx++ {
		if sortableSli[idx] == sortableSli[idx+1] {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(hasDuplicateValue(SortableSlice{5, 9, 3, 2, 4, 5, 6})) // true
}
