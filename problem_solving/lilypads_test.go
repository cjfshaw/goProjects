package main

import "testing"

type testPond struct {
	Stack myStack
	Node0 node
	Node1 node
	Node2 node
}

func createElements() testPond {
	var pad1, pad2, pad3 node

	pad1.Value = 1
	pad1.Left = &pad3
	pad1.Right = &pad2

	pad2.Value = 2
	pad2.Left = &pad1
	pad2.Right = &pad3

	pad3.Value = 3
	pad3.Left = &pad2
	pad3.Right = &pad1

	stack := myStack{Stack: make([]*node, 0)}

	elements := testPond{Stack: stack, Node0: pad1, Node1: pad2, Node2: pad3}

	return elements
}

func TestPush(t *testing.T) {
	elements := createElements()

	t.Log("Testing push on stack.  Expecting an empty stack to gain one entry.")

	t.Logf("Starting Stack: %v", elements.Stack)
	t.Logf("Pointer to Node 0: %v", &elements.Node0)

	elements.Stack.push(&elements.Node0)

	if len(elements.Stack.Stack) != 1 {
		t.Errorf("Error: Length of stack did not increase from 0 to 1.\n Length: %v", len(elements.Stack.Stack))
	}

	if elements.Stack.Stack[0] != &elements.Node0 {
		t.Errorf("Error: Top of stack is not node0.\nStack: %v, Node0: %v", elements.Stack.Stack, elements.Node0)
	}
}

func TestPop(t *testing.T) {
	elements := createElements()

	t.Log("Testing pop on stack.  Expecting a populated stack to lose one entry.")

	elements.Stack.push(&elements.Node0)
	elements.Stack.push(&elements.Node1)
	elements.Stack.push(&elements.Node2)

	t.Logf("Starting stack: %v", elements.Stack)

	t.Logf("-2: %v", elements.Stack.Stack[0:len(elements.Stack.Stack)-2])
	t.Logf("-1: %v", elements.Stack.Stack[0:len(elements.Stack.Stack)-1])
	t.Logf("-0: %v", elements.Stack.Stack[0:len(elements.Stack.Stack)])

	elements.Stack.pop()

	if len(elements.Stack.Stack) != 2 {
		t.Errorf("Error: Length of stack did not decrease from 3 to 2.\n Length: %v", len(elements.Stack.Stack))
	}

	if elements.Stack.Stack[1] != &elements.Node1 {
		t.Errorf("Error: Incorrect node was removed, top of stack is not Node1.\nStack: %v\n, Node1: %v\n", elements.Stack.Stack, elements.Node1)
	}
}
