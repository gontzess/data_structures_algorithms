package main

import (
	"fmt"
)

type TreeNode struct {
	value      int
	leftChild  *TreeNode
	rightChild *TreeNode
}

func search(searchValue int, node *TreeNode) *TreeNode {
	if node == nil || searchValue == node.value {
		return node
	} else if searchValue < node.value {
		return search(searchValue, node.leftChild)
	} else { // searchValue > node.value
		return search(searchValue, node.rightChild)
	}
}

func insert(value int, node *TreeNode) {
	if value < node.value {
		if node.leftChild == nil {
			node.leftChild = &TreeNode{value: value}
		} else {
			insert(value, node.leftChild)
		}
	} else if value > node.value {
		if node.rightChild == nil {
			node.rightChild = &TreeNode{value: value}
		} else {
			insert(value, node.rightChild)
		}
	}
}

func delete(valueToDelete int, node *TreeNode) *TreeNode {
	if node == nil {
		return node
	} else if valueToDelete < node.value {
		node.leftChild = delete(valueToDelete, node.leftChild)
		return node
	} else if valueToDelete > node.value {
		node.rightChild = delete(valueToDelete, node.rightChild)
		return node
	} else { // valueToDelete == node.value
		if node.leftChild == nil {
			return node.rightChild
		} else if node.rightChild == nil {
			return node.leftChild
		} else {
			node.rightChild = lift(node.rightChild, node)
			return node
		}
	}
}

func lift(node *TreeNode, nodeToDelete *TreeNode) *TreeNode {
	if node.leftChild != nil {
		node.leftChild = lift(node.leftChild, nodeToDelete)
		return node
	} else {
		nodeToDelete.value = node.value
		return node.rightChild
	}
}

// inorder traversal
func traverseInorder(node *TreeNode) {
	if node == nil {
		return
	}
	traverseInorder(node.leftChild)
	fmt.Println(node.value)
	traverseInorder(node.rightChild)
}

// preorder traversal
func traversePreorder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	traversePreorder(node.leftChild)
	traversePreorder(node.rightChild)
}

// postorder traversal
func traversePostorder(node *TreeNode) {
	if node == nil {
		return
	}
	traversePostorder(node.leftChild)
	traversePostorder(node.rightChild)
	fmt.Println(node.value)
}

func greatestValue(node *TreeNode) int {
	if node.rightChild != nil {
		return greatestValue(node.rightChild)
	} else {
		return node.value
	}
}

func main() {
	// node1 := TreeNode{value: 25}
	// node2 := TreeNode{value: 75}
	// root := TreeNode{50, &node1, &node2}
	// fmt.Println(root)

	// root := TreeNode{value: 3}
	// insert(2, &root)
	// insert(4, &root)
	// insert(1, &root)
	// insert(5, &root)
	// fmt.Println(search(5, &root))
	// fmt.Println(search(6, &root))

	root := TreeNode{value: 50}
	insert(25, &root)
	insert(75, &root)
	insert(11, &root)
	insert(33, &root)
	insert(30, &root)
	insert(40, &root)
	insert(56, &root)
	insert(52, &root)
	insert(61, &root)
	insert(89, &root)
	insert(82, &root)
	insert(95, &root)
	insert(55, &root)
	delete(50, &root)
	fmt.Println("-")
	traverseInorder(&root)
	fmt.Println("-")
	traversePreorder(&root)
	fmt.Println("-")
	traversePostorder(&root)
	fmt.Println("-")
	fmt.Println(greatestValue(&root))
}
