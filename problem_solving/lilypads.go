package main

import "fmt"

type node struct {
	Value int
	Left  *node
	Right *node
}

type pondPaths struct {
	Paths [][]*node
}

type myStack struct {
	Stack []*node
}

func (s *myStack) push(pad *node) {
	s.Stack = append(s.Stack, pad)
}

func (s *myStack) pop() {
	s.Stack = s.Stack[:len(s.Stack)-1]
}

func createNodes() (node, node, node) {
	var pad1, pad2, pad3 node

	pad1.Value = 1
	pad1.Left = &pad3
	pad1.Right = &pad2

	pad2.Value = 2
	pad2.Left = &pad3
	pad2.Right = &pad1

	pad3.Value = 3
	pad3.Left = &pad1
	pad3.Right = &pad2

	return pad1, pad2, pad3
}

func viewStackByValue(stack []*node) {
	for i := 0; i < len(stack); i++ {
		fmt.Printf("Node%v, ", stack[i].Value)
	}
	fmt.Println()
}

func viewPathsByValue(paths *pondPaths) {
	for i := 0; i < len(paths.Paths); i++ {
		fmt.Printf("Path: ")
		viewStackByValue(paths.Paths[i])
	}
	fmt.Println()
}

func (paths *pondPaths) getPaths(currentPath myStack, numHops int, currentNode *node) {

	if len(currentPath.Stack) == 0 {
		currentPath.push(currentNode)
	} else {
		currentPath.push(currentNode)
		numHops--
	}

	if numHops == 0 {
		paths.Paths = append(paths.Paths, currentPath.Stack)
		currentPath.pop()
		return
	}

	left := currentNode.Left
	right := currentNode.Right

	paths.getPaths(currentPath, numHops, left)
	paths.getPaths(currentPath, numHops, right)
}

func main() {
	pad0, pad1, pad2 := createNodes()
	stack := myStack{Stack: make([]*node, 0)}

	numHops := 2

	paths := pondPaths{Paths: make([][]*node, 0)}

	paths.getPaths(stack, numHops, &pad0)

	fmt.Println(pad0)
	fmt.Println(pad1)
	fmt.Println(pad2)

	viewPathsByValue(&paths)
}
