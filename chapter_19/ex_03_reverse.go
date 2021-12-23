// Create a new function to reverse an array that takes up just O(1) extra space.

package main

import "fmt"

// space complexity of O(N)
func reverse1(sli []int) []int {
	var newSli []int
	for idx := len(sli) - 1; idx >= 0; idx-- {
		newSli = append(newSli, sli[idx])
	}

	return newSli
}

// space complexity of O(1)
func reverse2(sli *[]int) {
	lastIdx := len(*sli) - 1
	for idx := 0; idx <= lastIdx/2; idx++ {
		(*sli)[idx], (*sli)[lastIdx-idx] = (*sli)[lastIdx-idx], (*sli)[idx]
	}
}

func main() {
	testSli := []int{65, 55, 45, 35, 25, 15, 10}
	// testSli := []int{65, 55, 45, 35, 25, 15, 10, 5}
	fmt.Println("rev1", reverse1(testSli))
	fmt.Println("orig not mutated", testSli)
	reverse2(&testSli)
	fmt.Println("rev 2, orig mutated", testSli)
}
