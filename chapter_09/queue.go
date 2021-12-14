package main

import "fmt"

type Queue struct {
  data []string
}

func (q *Queue) enqueue(element string) *Queue {
  q.data = append(q.data, element)

  return q
}

func (q *Queue) dequeue() string {
  if q.isEmpty() {
    return ""
  }

  first := q.data[0]
  q.data = q.data[1:]

  return first
}

func (q *Queue) read() string {
  if q.isEmpty() {
    return ""
  }

  return q.data[0]
}

func (q *Queue) isEmpty() bool {
  return len(q.data) - 1 < 0
}

func main() {
  queue := &Queue{}
  fmt.Println(queue)
  queue.enqueue("hi")
  queue.enqueue("there")
  fmt.Println(queue)
  fmt.Println(queue.read())
  fmt.Println(queue.dequeue())
  fmt.Println(queue)
  fmt.Println(queue.isEmpty())
  fmt.Println(queue.dequeue())
  fmt.Println(queue)
  fmt.Println(queue.isEmpty())
  fmt.Println(queue.dequeue())
  fmt.Println(queue)
}
