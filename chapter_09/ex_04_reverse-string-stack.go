// Write a function that uses a stack to reverse a string. (For example, "abcde" would become "edcba".) You can work with our earlier implementation of the Stack class.

package main

import (
  "fmt"
)

type Stack struct {
  data []string
}

func (s *Stack) push(element string) *Stack {
  s.data = append(s.data, element)

  return s
}

func (s *Stack) pop() string {
  lastIdx := len(s.data) - 1
  if lastIdx < 0 {
    return ""
  }

  last := s.data[lastIdx]
  s.data = s.data[:lastIdx]

  return last
}

func (s *Stack) read() string {
  lastIdx := len(s.data) - 1
  if lastIdx < 0 {
    return ""
  }

  return s.data[lastIdx]
}

func (s *Stack) isEmpty() bool {
  return len(s.data) == 0
}

func reverse(fwdString string) string {
  stack := &Stack{}
  for _, byteVal := range []byte(fwdString) {
    stack.push(string(byteVal))
  }

  revrsString := ""
  for stack.isEmpty() == false {
    revrsString += stack.pop()
  }

  return revrsString
}

func main() {
  testStr := "abcde"
  fmt.Println(reverse(testStr))
  fmt.Println(testStr)
}
