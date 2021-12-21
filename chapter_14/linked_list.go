package main

import (
	"fmt"
)

type Node struct {
	data     string
	nextNode *Node
}

type LinkedList struct {
	firstNode *Node
}

func (ll LinkedList) read(idx int) string {
	currentNode := ll.firstNode
	currentIdx := 0

	for currentIdx < idx {
		currentNode = currentNode.nextNode
		currentIdx += 1

		if currentNode == nil {
			return ""
		}
	}

	return currentNode.data
}

func (ll LinkedList) indexOf(value string) int {
	currentNode := ll.firstNode
	currentIdx := 0

	for currentNode != nil {
		if currentNode.data == value {
			return currentIdx
		}

		currentNode = currentNode.nextNode
		currentIdx += 1
	}

	return -1
}

func (ll LinkedList) insertAtIndex(idx int, value string) {
	newNode := Node{data: value}

	if idx == 0 {
		newNode.nextNode = ll.firstNode
		ll.firstNode = &newNode
		return
	}

	currentNode := ll.firstNode
	currentIdx := 0

	for currentIdx < (idx - 1) {
		currentNode = currentNode.nextNode
		currentIdx += 1
	}

	newNode.nextNode = currentNode.nextNode
	currentNode.nextNode = &newNode
}

func (ll LinkedList) deleteAtIndex(idx int) {
	if idx == 0 {
		ll.firstNode = ll.firstNode.nextNode
		return
	}

	currentNode := ll.firstNode
	currentIdx := 0

	for currentIdx < (idx - 1) {
		currentNode = currentNode.nextNode
		currentIdx += 1
	}

	nodeAfterDeletedNode := currentNode.nextNode.nextNode
	currentNode.nextNode = nodeAfterDeletedNode
}

func main() {
	node1 := Node{data: "once"}
	node2 := Node{data: "upon"}
	node3 := Node{data: "a"}
	node4 := Node{data: "time"}

	node1.nextNode = &node2
	node2.nextNode = &node3
	node3.nextNode = &node4
	fmt.Println(node1, node2, node3, node4)

	list := LinkedList{firstNode: &node1}
	fmt.Println(list.read(3))         // "time"
	fmt.Println(list.read(8))         // ""
	fmt.Println(list.indexOf("a"))    // 2
	fmt.Println(list.indexOf("time")) // 3
	fmt.Println(list.indexOf("in"))   // -1

	list.insertAtIndex(3, "purple")
	fmt.Println(list.indexOf("a"))      // 2
	fmt.Println(list.indexOf("time"))   // 4
	fmt.Println(list.indexOf("purple")) // 3

	list.deleteAtIndex(3)
	fmt.Println(list.indexOf("a"))    // 2
	fmt.Println(list.indexOf("time")) // 3
}
