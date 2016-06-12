package main

import (
	"fmt"
)

type node struct {
	Value int
	Left  *node
	Right *node
}

type tree struct {
	Root *node
}

type testElements struct {
	myTree    *tree
	node0     *node
	node1     *node
	node2     *node
	node3     *node
	node4     *node
	node5     *node
	node6     *node
	node7     *node
	nodeArray []*node
}

func createTestElements() testElements {
	node0 := node{Value: 0, Left: nil, Right: nil}
	node1 := node{Value: 1, Left: nil, Right: nil}
	node2 := node{Value: 2, Left: nil, Right: nil}
	node3 := node{Value: 3, Left: nil, Right: nil}
	node4 := node{Value: 4, Left: nil, Right: nil}
	node5 := node{Value: 5, Left: nil, Right: nil}
	node6 := node{Value: 6, Left: nil, Right: nil}
	node7 := node{Value: 7, Left: nil, Right: nil}

	node0.Left = &node1
	node0.Right = &node2

	node1.Left = &node3

	node2.Left = &node4
	node2.Right = &node5

	node3.Left = &node6
	node3.Right = &node7

	myTree := tree{Root: &node0}

	var nodeArray []*node

	testElements := testElements{
		myTree: &myTree,
		node0:  &node0, node1: &node1, node2: &node2, node3: &node3,
		node4: &node4, node5: &node5, node6: &node6, node7: &node7,
		nodeArray: nodeArray,
	}

	return testElements
}

func printNodeValues(stack []*node) {
	fmt.Println("")
	for _, value := range stack {
		fmt.Printf("%d  ", value.Value)
	}
}

func depthFirstSearchNoStack(currentNode *node) {
	fmt.Printf("%d  ", currentNode.Value)
	if currentNode.Left != nil {
		depthFirstSearchNoStack(currentNode.Left)
	}
	if currentNode.Right != nil {
		depthFirstSearchNoStack(currentNode.Right)
	}
}

func depthFirstSearchWithStack(stack []*node) {
	fmt.Println("\nCurrent Node: ", stack[len(stack)-1])

	if stack[len(stack)-1].Left == nil && stack[len(stack)-1].Right == nil {
		stack = stack[:len(stack)-1]
		return
	}

	if stack[len(stack)-1].Left != nil {
		fmt.Printf("\nLeft node of %v is not nil, so adding %v to the stack which currently equals:", stack[len(stack)-1], stack[len(stack)-1].Left)
		printNodeValues(stack)
		stack = append(stack, stack[len(stack)-1].Left)
		depthFirstSearchWithStack(stack)
	}

	if stack[len(stack)-1].Right != nil {
		fmt.Printf("\nRight node of %v is not nil, so adding %v to the stack which currently equals:", stack[len(stack)-1], stack[len(stack)-1].Right)
		printNodeValues(stack)
		stack = append(stack, stack[len(stack)-1].Right)
		depthFirstSearchWithStack(stack)
	}
}

func main() {
	fmt.Println("Running main")

	testElements := createTestElements()

	testElements.nodeArray = append(testElements.nodeArray, testElements.myTree.Root)

	fmt.Println("DFS with no stack")
	depthFirstSearchNoStack(testElements.nodeArray[0])

	var stack []*node
	stack = append(stack, testElements.myTree.Root)

	fmt.Println("DFS with a stack")
	depthFirstSearchWithStack(stack)
}
