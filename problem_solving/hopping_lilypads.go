package main

import (
	"fmt"
)

type lilypad struct {
	Value int
	Left  *lilypad
	Right *lilypad
}

type stack struct {
	Stack []*lilypad
}

func (s *stack) push(pad *lilypad) {
	s.Stack = append(s.Stack, pad)
}

func (s *stack) pop() {
	s.Stack = s.Stack[:len(s.Stack)-2]
}

func (s *stack) lastIn() *lilypad {
	return s.Stack[len(s.Stack)-1]
}

func createLilypads() (lilypad, lilypad, lilypad) {
	var pad1, pad2, pad3 lilypad

	pad1.Value = 1
	pad1.Left = &pad3
	pad1.Right = &pad2

	pad2.Value = 2
	pad2.Left = &pad1
	pad2.Right = &pad3

	pad3.Value = 3
	pad3.Left = &pad2
	pad3.Right = &pad1

	return pad1, pad2, pad3
}

func (pad *lilypad) getHoppablePads() []*lilypad {
	hoppablePads := make([]*lilypad, 2)

	hoppablePads[0] = pad.Left
	hoppablePads[1] = pad.Right

	return hoppablePads
}

func findPaths(pathArray [][]int, currentPad *lilypad, totalHops int, stack *stack) [][]int {
	currentArray := make([]int, 0)
	remainingHops := totalHops - len(stack.Stack) + 1

	fmt.Printf("Starting findPaths with the following info: \n")
	for _, currentPad := range stack.Stack {
		fmt.Printf("stack value: %v,", currentPad.Value)
	}
	fmt.Println()
	fmt.Printf("remainingHops: %v\n", remainingHops)

	if remainingHops <= 0 {
		pathArray = append(pathArray, currentArray)
		return pathArray
	}

	fmt.Printf("pathArray: %v", pathArray)

	for _, currentPad := range stack.Stack {
		fmt.Printf("appending")
		currentArray = append(currentArray, currentPad.Value)
	}

	hoppablePads := stack.lastIn().getHoppablePads()

	for _, currentPad := range hoppablePads {
		stack.push(currentPad)
		remainingHops--
		pathArray = findPaths(pathArray, currentPad, totalHops, stack)
	}

	fmt.Printf("pathArray: %v", pathArray)

	return pathArray
}

func main() {
	//you have x lilypads
	//the lilypads are arranged in a circle
	//you have y number of hops
	//you can only hop to your left or to your right
	//you always start on the same pad (1)
	//determine all of the possible paths you can take
	//example for 3 lilypads and 2 hops you would return something like...
	//[[1,2,3],[1,2,1],[1,3,2],[1,3,1]]
	//example for 4 lilypads and 3 hops
	//[[1,2,3,4][1,2,3,2][1,2,1,2][1,2,1,4][1,4,3,2][1,4,3,4][1,4,1,4][1,4,1,2]]

	//so I need to
	//get current pad
	//find what I can touch
	//choose one of the options
	//add that option to my current list
	//decrement hops available
	//repeat this until hops are 0
	//when hops are 0 return the list
	fmt.Println("main started")

	pad1, pad2, pad3 := createLilypads()

	fmt.Println("to avoid errors...", pad2, pad3)

	startingPad := pad1

	pathArray := make([][]int, 0)

	//numPads := 3
	numHops := 2

	var myStack stack

	myStack.push(&pad1)

	pathArray = findPaths(pathArray, &startingPad, numHops, &myStack)

	fmt.Println(pathArray)

}
