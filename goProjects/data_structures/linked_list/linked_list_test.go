package main

import (
	"testing"
)

func TestAddToEndNilCase(t *testing.T) {
	list := list{nil, nil}

	node := node{1, nil, nil}

	t.Log("Expecting list to be populated, head and tail, with the added node.")
	t.Logf("Initial list: %+v.\n", list)
	list.addToEnd(&node)
	t.Logf("Final list: %+v.\n", list)
	t.Logf("Final list head: %+v.\n", list.Head)
	t.Logf("Final list tail: %+v.\n", list.Tail)

	if list.Head != &node {
		t.Errorf("Error: Head is not equal to added node.")
	}

	if list.Tail != &node {
		t.Errorf("Error: Tail is not equal to added node.")
	}

	if list.Head != list.Tail {
		t.Errorf("Error: Head is not equal to tail.")
	}
}

func TestAddToEnd(t *testing.T) {
	//makes 3 nodes, two of which are in a list on creation
	//adds the third and verifies that the list is now 3, with the correct order

	node1 := node{1, nil, nil}
	node2 := node{2, nil, nil}
	node3 := node{3, nil, nil}

	node1.Next = &node2
	node2.Prev = &node1

	list := list{&node1, &node2}

	t.Log("Expecting final node to be node 3, pointing to previous node 2 and next node nil.")
	t.Logf("Initial list: %+v\n", list)
	t.Logf("Initial tail: %+v\n", list.Tail)

	list.addToEnd(&node3)

	t.Logf("Final list: %+v\n", list)
	t.Logf("Final tail: %+v\n", list.Tail)

	if list.Tail != &node3 {
		t.Errorf("Error: Final node is not node 3.\n")
	}

	if list.Tail.Prev != &node2 {
		t.Errorf("Error: Final node is not pointing to previous tail.\n")
	}

	if list.Tail.Next != nil {
		t.Errorf("Error: Final node's next is not nil.\n")
	}
}

func TestAddToBeginningNilCase(t *testing.T) {
	list := list{nil, nil}

	node := node{1, nil, nil}

	t.Log("Expecting list to be populated, head and tail, with the added node.")
	t.Logf("Initial list: %+v.\n", list)
	list.addToBeginning(&node)
	t.Logf("Final list: %+v.\n", list)
	t.Logf("Final list head: %+v.\n", list.Head)
	t.Logf("Final list tail: %+v.\n", list.Tail)

	if list.Head != &node {
		t.Errorf("Error: Head is not equal to added node.")
	}

	if list.Tail != &node {
		t.Errorf("Error: Tail is not equal to added node.")
	}

	if list.Head != list.Tail {
		t.Errorf("Error: Head is not equal to tail.")
	}
}

func TestAddToBeginning(t *testing.T) {
	node1 := node{1, nil, nil}
	node2 := node{2, nil, nil}
	node3 := node{3, nil, nil}

	node1.Next = &node2
	node2.Prev = &node1

	list := list{&node1, &node2}

	t.Log("Expecting first node to node 3, pointing Next to node 1 and previous to nil.")
	t.Logf("Initial list: %+v", list)
	t.Logf("Initial head: %+v", list.Head)

	list.addToBeginning(&node3)

	t.Logf("Initial list: %+v", list)
	t.Logf("Initial head: %+v", list.Head)

	if list.Head != &node3 {
		t.Errorf("Error: Head is not equal to added node.")
	}

	if list.Head.Next != &node1 {
		t.Errorf("Error: Head is not pointing next to node 1.")
	}

	if list.Head.Prev != nil {
		t.Errorf("Error: Head is not pointin previous to nil.")
	}
}

func TestFindNodeByValue(t *testing.T) {
	node1 := node{1, nil, nil}
	node2 := node{2, nil, nil}
	node3 := node{3, nil, nil}

	node1.Next = &node2
	node2.Prev = &node1
	node2.Next = &node3
	node3.Prev = &node2

	list := list{&node1, &node3}

	foundNode, err := list.findNodeByValue(2)

	if err != nil {
		t.Errorf("Error when running findNodeByValue: \n%+v", err)
	}

	t.Log("Expecting to find node 2.")
	t.Logf("Value of node 2: %+v", node2)
	t.Logf("Value of found node: %+v", foundNode)

	if node2 != *foundNode {
		t.Errorf("Error: Node found was incorrect.")
	}
}

func TestFindNodeByValueNoNodePresent(t *testing.T) {
	node1 := node{1, nil, nil}
	node2 := node{2, nil, nil}
	node3 := node{3, nil, nil}

	node1.Next = &node2
	node2.Prev = &node1
	node2.Next = &node3
	node3.Prev = &node2

	list := list{&node1, &node3}

	t.Log("Expecting to find no matching node.")

	foundNode, err := list.findNodeByValue(4)

	t.Logf("foundNode: %+v", foundNode)
	t.Logf("Error produced: %+v", err)

	if err == nil {
		t.Errorf("Error: No error produced when running findNodeByValue.")
	}

	if foundNode != nil {
		t.Errorf("Error: foundNode is not nil even though no node with that value is present in the list.")
	}
}
