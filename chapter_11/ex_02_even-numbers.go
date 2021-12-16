// Use recursion to write a function that accepts an array of numbers and returns a new array containing just the even numbers.

package main

import (
  "fmt"
)

func evenNumbers(numbers []int) []int {
  var evens []int
  if len(numbers) == 0 {
    return evens
  }

  if numbers[0] % 2 == 0 {
    evens = append(evens, numbers[0])
  }

  evens = append(evens, evenNumbers(numbers[1:])...)
  return evens
}

func main() {
  numbers := []int{1, 2, 6, 7, 9, 3, 4, 5, 14, 201, -8}
  fmt.Println(evenNumbers(numbers))
}
