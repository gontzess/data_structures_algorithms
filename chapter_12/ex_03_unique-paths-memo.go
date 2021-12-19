package main

import (
	"fmt"
	"strconv"
)

func uniquePaths(cols int, rows int, memo map[string]int) int {
	if cols == 1 || rows == 1 {
		return 1
	}

	pathKey := strconv.Itoa(cols) + "," + strconv.Itoa(rows)
	if _, present := memo[pathKey]; present == false {
		memo[pathKey] = uniquePaths(cols-1, rows, memo) + uniquePaths(cols, rows-1, memo)
	}

	return memo[pathKey]
}

func main() {
	fmt.Println(uniquePaths(7, 3, make(map[string]int)))
}
