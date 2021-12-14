package main

import "fmt"

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
  return len(s.data) - 1 < 0
}

func main() {
  stack := &Stack{}
  fmt.Println(stack)
  stack.push("hi")
  stack.push("there")
  fmt.Println(stack)
  fmt.Println(stack.read())
  fmt.Println(stack.pop())
  fmt.Println(stack)
  fmt.Println(stack.pop())
  fmt.Println(stack)
  fmt.Println(stack.isEmpty())
  fmt.Println(stack.pop())
  fmt.Println(stack)
}
