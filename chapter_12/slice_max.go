package main

import "fmt"

func maxBad(sli []int) int {
	fmt.Println("RECURSION")
	if len(sli) == 1 {
		return sli[0]
	}

	if sli[0] > maxBad(sli[1:]) {
		return sli[0]
	} else {
		return maxBad(sli[1:])
	}
}

func maxGood(sli []int) int {
	fmt.Println("RECURSION")
	if len(sli) == 1 {
		return sli[0]
	}

	maxOfRemainder := maxGood(sli[1:])

	if sli[0] > maxOfRemainder {
		return sli[0]
	} else {
		return maxOfRemainder
	}
}

func main() {
	sli := []int{1, 2, 3, 4}
	fmt.Println(maxBad(sli))
	fmt.Println(maxGood(sli))
}
