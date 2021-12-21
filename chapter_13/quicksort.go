package main

import (
	"fmt"
)

type SortableSlice []int

func (ss SortableSlice) partition(leftPtr int, rightPtr int) int {
	pivotIdx := rightPtr
	pivot := ss[pivotIdx]
	rightPtr -= 1

	for {
		for ss[leftPtr] < pivot {
			leftPtr += 1
		}

		for ss[rightPtr] > pivot {
			rightPtr -= 1
		}

		if leftPtr >= rightPtr {
			break
		} else {
			ss[leftPtr], ss[rightPtr] = ss[rightPtr], ss[leftPtr]
			leftPtr += 1
		}
	}

	ss[leftPtr], ss[pivotIdx] = ss[pivotIdx], ss[leftPtr]

	return leftPtr
}

func (ss SortableSlice) quickSort(leftIdx int, rightIdx int) {
	// fmt.Println("RECURSE")
	if rightIdx-leftIdx <= 0 {
		return
	}

	pivotIdx := ss.partition(leftIdx, rightIdx)

	ss.quickSort(leftIdx, pivotIdx-1)
	ss.quickSort(pivotIdx+1, rightIdx)
}

func main() {
	sortableSli := SortableSlice{0, 5, 2, 1, 6, 3}
	fmt.Println(sortableSli)
	sortableSli.quickSort(0, len(sortableSli)-1)
	fmt.Println(sortableSli)
}
