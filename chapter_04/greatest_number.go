package main

import "fmt"

func main() {
  testSli := []int{34, 93, 92, 10}
  fmt.Println(greatestNumberLinear(testSli))
}

func greatestNumberLinear(sli []int) int {
  greatest := sli[0]
  for _, i := range sli {
    if i > greatest {
      greatest = i
    }
  }
  return greatest
}
