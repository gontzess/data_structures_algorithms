// Use recursion to write a function that accepts an array of strings and returns the total number of characters across all the strings. For example, if the input array is ["ab", "c", "def", "ghij"], the output should be 10 since there are 10 characters in total.

package main

import (
  "fmt"
)

func totalCharacters(input []string) int {
  if len(input) == 1 {
    return len(input[0])
  }

  return len(input[0]) + totalCharacters(input[1:])
}

func main(){
  testSli := []string{"ab", "c", "def", "ghij"}
  fmt.Println(totalCharacters(testSli))
}
