package main

import "fmt"

func main() {
	testSli := &[]int{65, 55, 45, 35, 25, 15, 10}
	fmt.Println(selectionSort(testSli))
}

func selectionSort(sliPtr *[]int) *[]int {
	sli := *sliPtr

	for start := 0; start <= len(sli)-2; start++ {
		lowestIdx := start
		for idx := start + 1; idx <= len(sli)-1; idx++ {
			if sli[idx] < sli[lowestIdx] {
				lowestIdx = idx
			}
		}
		if lowestIdx != start {
			sli[start], sli[lowestIdx] = sli[lowestIdx], sli[start]
		}
	}

	return sliPtr
}
