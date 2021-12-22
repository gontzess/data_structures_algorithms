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

func (ll *LinkedList) insertAtIndex(idx int, value string) {
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

func (ll *LinkedList) deleteAtIndex(idx int) {
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

func (ll LinkedList) print() {
	currentNode := ll.firstNode
	resultStr := ""
	for currentNode != nil {
		fmt.Print(currentNode)
		resultStr += currentNode.data + ", "
		currentNode = currentNode.nextNode
	}

	if len(resultStr) > 0 {
		resultStr = resultStr[:len(resultStr)-2]
	}

	fmt.Println("[" + resultStr + "]")
}

func (ll LinkedList) last() string {
	currentNode := ll.firstNode
	if currentNode == nil {
		return ""
	}

	for currentNode.nextNode != nil {
		currentNode = currentNode.nextNode
	}

	return currentNode.data
}

func (ll *LinkedList) reverse() {
	var previousNode, nextNode *Node
	currentNode := ll.firstNode

	for currentNode != nil {
		nextNode = currentNode.nextNode
		currentNode.nextNode = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}

	ll.firstNode = previousNode
}

// Let's say you have access to a node from somewhere in the middle of a classic linked list, but not the linked list itself. That is, you have a variable that points to an instance of Node, but you don't have access to the LinkedList instance. In this situation, if you follow this node's link, you can find all the items from this middle node until the end, but you have no way to find the nodes that precede this node in the list. Write code that will effectively delete this node from the list. The entire remaining list should remain complete, with only this node removed.

func (ll LinkedList) getNode(idx int) *Node {
	currentNode := ll.firstNode
	currentIdx := 0

	for currentIdx < idx {
		currentNode = currentNode.nextNode
		currentIdx += 1
	}

	return currentNode
}

func deleteMiddleNode(node *Node) {
	node.data = node.nextNode.data
	node.nextNode = node.nextNode.nextNode
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

	list := LinkedList{}
	fmt.Println(list.last()) // ""
	list.firstNode = &node1
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

	list.deleteAtIndex(0)
	fmt.Println(list.indexOf("a"))    // 1
	fmt.Println(list.indexOf("time")) // 2

	list.insertAtIndex(0, "ONCE")
	fmt.Println(list.firstNode.data) // "ONCE"
	fmt.Println(list.indexOf("a"))   // 2

	list.print()             // [ONCE, upon, a, time]
	fmt.Println(list.last()) // "time"
	list.reverse()
	list.print() // [time, a, upon, ONCE]
	list.reverse()
	list.print() // [ONCE, upon, a, time]

	list.insertAtIndex(4, "in")
	list.insertAtIndex(5, "Mexico")
	list.print() // [ONCE, upon, a, time, in, Mexico]

	nodeForDeletion1 := list.getNode(3)
	fmt.Println(nodeForDeletion1) // "time"
	deleteMiddleNode(nodeForDeletion1)
	list.print() // [ONCE, upon, a, in, Mexico]

	nodeForDeletion2 := list.getNode(3)
	fmt.Println(nodeForDeletion2) // "in"
	deleteMiddleNode(nodeForDeletion2)
	list.print() // [ONCE, upon, a, Mexico]
}
