package main

import "fmt"

func golomb(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}

	if _, present := memo[n]; present == false {
		memo[n] = 1 + golomb(n-golomb(golomb(n-1, memo), memo), memo)
	}

	return memo[n]
}

func main() {
	fmt.Println(golomb(200, make(map[int]int)))
}
