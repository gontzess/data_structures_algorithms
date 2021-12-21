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

// NOTE side affect:  mutates the caller leaving a partially sorted slice
func (ss SortableSlice) quickSelectMut(kthLowestVal int, leftIdx int, rightIdx int) int {
	// fmt.Println("RECURSE")
	if rightIdx-leftIdx <= 0 {
		return ss[leftIdx]
	}

	pivotIdx := ss.partition(leftIdx, rightIdx)

	if kthLowestVal < pivotIdx {
		return ss.quickSelectMut(kthLowestVal, leftIdx, pivotIdx-1)
	} else if kthLowestVal > pivotIdx {
		return ss.quickSelectMut(kthLowestVal, pivotIdx+1, rightIdx)
	} else {
		return ss[pivotIdx]
	}
}

func (ss SortableSlice) clone() SortableSlice {
	sliceCopy := make(SortableSlice, len(ss))
	copy(sliceCopy, ss)
	return sliceCopy
}

// adjusted so that quickSelect doesn't mutate the original caller
func (ss SortableSlice) quickSelect(kthLowestVal int, leftIdx int, rightIdx int) int {
	// fmt.Println("RECURSE")
	copy := ss.clone()
	if rightIdx-leftIdx <= 0 {
		return copy[leftIdx]
	}

	pivotIdx := copy.partition(leftIdx, rightIdx)

	if kthLowestVal < pivotIdx {
		return copy.quickSelect(kthLowestVal, leftIdx, pivotIdx-1)
	} else if kthLowestVal > pivotIdx {
		return copy.quickSelect(kthLowestVal, pivotIdx+1, rightIdx)
	} else {
		// fmt.Println(copy)
		return copy[pivotIdx]
	}
}

func main() {
	sortableSli := SortableSlice{0, 50, 20, 10, 60, 30}
	fmt.Println(sortableSli)
	secondLowestVal := sortableSli.quickSelect(1, 0, len(sortableSli)-1)
	fmt.Println(secondLowestVal)
	fmt.Println(sortableSli)

	secondLowestVal2 := sortableSli.quickSelectMut(1, 0, len(sortableSli)-1)
	fmt.Println(secondLowestVal2)
	fmt.Println(sortableSli)
}
