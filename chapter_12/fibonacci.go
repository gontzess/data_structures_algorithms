package main

import "fmt"

func fib1(n int) int {
	fmt.Println("RECURSION")
	if n == 0 || n == 1 {
		return n
	}

	return fib1(n-2) + fib1(n-1)
}

// memoization
func fib2(n int, memo map[int]int) int {
	fmt.Println("RECURSION")
	if n == 0 || n == 1 {
		return n
	}

	if _, present := memo[n]; present == false {
		memo[n] = fib2(n-2, memo) + fib2(n-1, memo)
	}

	return memo[n]
}

// going bottom up
func fib3(n int) int {
	fmt.Println("NO recursion")
	if n == 0 {
		return 0
	}

	a, b := 0, 1
	for idx := 1; idx < n; idx++ {
		temp := a
		a = b
		b = temp + a
	}

	return b
}

func main() {
	fmt.Println(fib1(6))
	fmt.Println(fib2(6, make(map[int]int)))
	fmt.Println(fib3(6))
}
