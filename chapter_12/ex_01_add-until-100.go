// The following function accepts an array of numbers and returns the sum, as long as a particular number doesn't bring the sum above 100. If adding a particular number will make the sum higher than 100, that number is ignored. However, this function makes unnecessary recursive calls. Fix the code to eliminate the unnecessary recursion:

package main

import "fmt"

func addUntil100(sli []int) int {
  if len(sli) == 0 {
    return 0
  }

  remaining := addUntil100(sli[1:])

  if sli[0] + remaining > 100 {
    return remaining
  } else {
    return sli[0] + remaining
  }
}

func main() {
  sli := []int{4, 6, 20, 50, 26, 25, 15}

  fmt.Println(addUntil100(sli))
}
