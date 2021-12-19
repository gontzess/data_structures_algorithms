// This problem is known as the "Unique Paths" problem: Let's say you have a grid of rows and columns. Write a function that accepts a number of rows and a number of columns, and calculates the number of possible "shortest" paths from the upper-leftmost square to the lower-rightmost square.

package main

import "fmt"

func countUniquePaths(cols int, rows int) int {
	if cols == 1 || rows == 1 {
		return 1
	}

	return countUniquePaths(cols-1, rows) + countUniquePaths(cols, rows-1)
}

func main() {
	fmt.Println(countUniquePaths(7, 3))
}
