package main

import "fmt"

func main() {
	testSli := &[]int{65, 55, 45, 35, 25, 15, 10}
	fmt.Println(bubbleSort(testSli))
}

func bubbleSort(sliPtr *[]int) *[]int {
	sli := *sliPtr
	end := len(sli) - 2
	sorted := false

	for sorted == false {
		sorted = true
		for idx := 0; idx <= end; idx++ {
			if sli[idx] > sli[idx+1] {
				sli[idx], sli[idx+1] = sli[idx+1], sli[idx]
				sorted = false
			}
		}
		end -= 1
	}

	return sliPtr
}
