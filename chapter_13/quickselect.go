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

// NOTE side affect:  mutates the caller leaving a partially sorted slice
func (ss *SortableSlice) quickSelectMut(kthLowestVal int, leftIdx int, rightIdx int) int {
	// fmt.Println("RECURSE")
	if rightIdx-leftIdx <= 0 {
		return ss.slice[leftIdx]
	}

	pivotIdx := ss.partition(leftIdx, rightIdx)

	if kthLowestVal < pivotIdx {
		return ss.quickSelectMut(kthLowestVal, leftIdx, pivotIdx-1)
	} else if kthLowestVal > pivotIdx {
		return ss.quickSelectMut(kthLowestVal, pivotIdx+1, rightIdx)
	} else {
		return ss.slice[pivotIdx]
	}
}

func (ss SortableSlice) clone() SortableSlice {
	sliceCopy := make([]int, len(ss.slice))
	copy(sliceCopy, ss.slice)
	return SortableSlice{slice: sliceCopy}
}

// adjusted so that quickSelect doesn't mutate the original caller
func (ss SortableSlice) quickSelect(kthLowestVal int, leftIdx int, rightIdx int) int {
	// fmt.Println("RECURSE")
	copy := ss.clone()
	if rightIdx-leftIdx <= 0 {
		return copy.slice[leftIdx]
	}

	pivotIdx := copy.partition(leftIdx, rightIdx)

	if kthLowestVal < pivotIdx {
		return copy.quickSelect(kthLowestVal, leftIdx, pivotIdx-1)
	} else if kthLowestVal > pivotIdx {
		return copy.quickSelect(kthLowestVal, pivotIdx+1, rightIdx)
	} else {
		fmt.Println(copy)
		return copy.slice[pivotIdx]
	}
}

func main() {
	sortableSli := SortableSlice{slice: []int{0, 50, 20, 10, 60, 30}}
	fmt.Println(sortableSli)
	secondLowestVal := sortableSli.quickSelect(1, 0, len(sortableSli.slice)-1)
	fmt.Println(secondLowestVal)
	fmt.Println(sortableSli)
}
