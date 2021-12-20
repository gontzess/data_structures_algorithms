package main

import (
	"fmt"
)

type SortableSlice struct {
	slice []int
}

func (ss *SortableSlice) partition(leftPtr int, rightPtr int) int {
	pivotIdx := rightPtr
	pivot := ss.slice[pivotIdx]
	rightPtr -= 1

	for {
		for ss.slice[leftPtr] < pivot {
			leftPtr += 1
		}

		for ss.slice[rightPtr] > pivot {
			rightPtr -= 1
		}

		if leftPtr >= rightPtr {
			break
		} else {
			ss.slice[leftPtr], ss.slice[rightPtr] = ss.slice[rightPtr], ss.slice[leftPtr]
			leftPtr += 1
		}
	}

	ss.slice[leftPtr], ss.slice[pivotIdx] = ss.slice[pivotIdx], ss.slice[leftPtr]

	return leftPtr
}

func (ss *SortableSlice) quickSort(leftIdx int, rightIdx int) {
	// fmt.Println("RECURSE")
	if rightIdx-leftIdx <= 0 {
		return
	}

	pivotIdx := ss.partition(leftIdx, rightIdx)

	ss.quickSort(leftIdx, pivotIdx-1)
	ss.quickSort(pivotIdx+1, rightIdx)
}

func main() {
	sortableSli := SortableSlice{slice: []int{0, 5, 2, 1, 6, 3}}
	fmt.Println(sortableSli)
	sortableSli.quickSort(0, len(sortableSli.slice)-1)
	fmt.Println(sortableSli)
}
