// Use recursion to write a function that accepts a string and returns the first index that contains the character "x". For example, the string, "abcdefghijklmnopqrstuvwxyz" has an "x" at index 23. To keep things simple, assume the string definitely has at least one "x".

package main

import (
  "fmt"
)

func firstXIndex(inputStr string, idx int) int {
  if string(inputStr[0]) == "x" {
    return idx
  }

  return firstXIndex(inputStr[1:], idx + 1)
}

func main() {
  testStr := "abcdefghijklmnopqrstuvwxyz"
  fmt.Println(firstXIndex(testStr, 0))
}
