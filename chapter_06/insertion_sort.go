package main

import "fmt"

func main() {
	testSli := &[]int{65, 55, 45, 35, 25, 15, 10}
	fmt.Println(insertionSort(testSli))
}

func insertionSort(sliPtr *[]int) *[]int {
	sli := *sliPtr

	for idx := 1; idx <= len(sli)-1; idx++ {
		selection := sli[idx]
		position := idx - 1
		for position >= 0 {
			if sli[position] > tempVal {
				sli[position+1] = sli[position]
				position -= 1
			} else {
				break
			}
		}
		sli[position+1] = selection
	}

	return sliPtr
}
