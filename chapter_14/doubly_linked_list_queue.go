package main

import (
	"fmt"
)

type Node struct {
	data         string
	nextNode     *Node
	previousNode *Node
}

type DoublyLinkedList struct {
	firstNode *Node
	lastNode  *Node
}

func (dll *DoublyLinkedList) insertAtEnd(value string) {
	newNode := Node{data: value}

	if dll.firstNode == nil {
		dll.firstNode = &newNode
		dll.lastNode = &newNode
	} else {
		newNode.previousNode = dll.lastNode
		dll.lastNode.nextNode = &newNode
		dll.lastNode = &newNode
	}
}

func (dll *DoublyLinkedList) removeFromFront() *Node {
	if dll.firstNode == nil {
		return dll.firstNode
	}

	removedNode := dll.firstNode
	dll.firstNode = dll.firstNode.nextNode
	dll.firstNode.previousNode = nil
	return removedNode
}

func (dll DoublyLinkedList) printReverse() {
	currentNode := dll.lastNode
	resultStr := ""
	for currentNode != nil {
		resultStr += currentNode.data + ", "
		currentNode = currentNode.previousNode
	}

	if len(resultStr) > 0 {
		resultStr = resultStr[:len(resultStr)-2]
	}

	fmt.Println("R[" + resultStr + "]")
}

type Queue struct {
	DoublyLinkedList
}

func (q *Queue) enqueue(element string) {
	q.insertAtEnd(element)
}

func (q *Queue) dequeue() string {
	if removedNode := q.removeFromFront(); removedNode == nil {
		return ""
	} else {
		return removedNode.data
	}
}

func (q Queue) read() string {
	if q.firstNode == nil {
		return ""
	} else {
		return q.firstNode.data
	}
}

func main() {
	list := DoublyLinkedList{}
	fmt.Println(list.removeFromFront()) // nil
	list.insertAtEnd("once")
	list.insertAtEnd("upon")
	list.insertAtEnd("a")
	list.insertAtEnd("time")

	fmt.Println(list.firstNode, list.lastNode)
	fmt.Println(list.removeFromFront())

	queue := Queue{}
	fmt.Println(queue.dequeue()) // ""
	fmt.Println(queue.read())    // ""
	queue.enqueue("once")
	// queue.insertAtEnd("once")
	// queue.DoublyLinkedList.insertAtEnd("once")
	fmt.Println(queue.read()) // "once"
	queue.enqueue("upon")
	queue.enqueue("a")
	queue.enqueue("time")
	fmt.Println(queue.firstNode, queue.lastNode)
	fmt.Println(queue.dequeue()) // "once"
	fmt.Println(queue.read())    // "upon"
	fmt.Println(queue.firstNode, queue.lastNode)

	queue.printReverse() // R[time, a, upon]
}
