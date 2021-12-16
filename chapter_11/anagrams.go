package main

import "fmt"

func anagrams(inputStr string) []string {
  var collection []string
  if len(inputStr) == 1 {
    collection = append(collection, inputStr)
    return collection
  }

  substringAnagrams := anagrams(inputStr[1:])
  // fmt.Println(substringAnagrams)

  for _, substring := range substringAnagrams {
    for idx := 0; idx <= len(substring); idx++ {
      copy := substring[:idx] + string(inputStr[0]) + substring[idx:]
      collection = append(collection, copy)
    }
  }

  return collection
}

func main() {
  testString := "abcd"
  fmt.Println(anagrams(testString))
}
