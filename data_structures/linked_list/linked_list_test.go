package main

import (
	"testing"
)

func TestAddToEndNilCase(t *testing.T) {
	list := list{nil, nil, 0}

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

	if list.Length != 1 {
		t.Errorf("Error: Length did not increase from 0 to 1.")
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

	list := list{&node1, &node2, 2}

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

	if list.Length != 3 {
		t.Errorf("Error: Length did not increase from 2 to 3.")
	}
}

func TestAddToBeginningNilCase(t *testing.T) {
	list := list{nil, nil, 0}

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

	if list.Length != 1 {
		t.Errorf("Error: Length did not increase from 0 to 1")
	}
}

func TestAddToBeginning(t *testing.T) {
	node1 := node{1, nil, nil}
	node2 := node{2, nil, nil}
	node3 := node{3, nil, nil}

	node1.Next = &node2
	node2.Prev = &node1

	list := list{&node1, &node2, 2}

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

	if list.Length != 3 {
		t.Errorf("Error: Length did not increase from 2 to 3")
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

	list := list{&node1, &node3, 3}

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

	list := list{&node1, &node3, 3}

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

func TestRemoveFromEnd(t *testing.T) {
	node1 := node{1, nil, nil}
	node2 := node{2, nil, nil}
	node3 := node{3, nil, nil}

	node1.Next = &node2
	node2.Prev = &node1
	node2.Next = &node3
	node3.Prev = &node2

	list := list{&node1, &node3, 3}

	t.Log("Expecting the list to be shortened by one node, with the final node being node 2.")
	t.Logf("Initial tail: %+v", list.Tail)

	list.removeFromEnd()

	t.Logf("Final tail: %+v", list.Tail)

	if list.Tail != &node2 {
		t.Errorf("Error: Tail is not equal to node 2.")
	}

	if list.Tail.Next != nil {
		t.Errorf("Error: Tail is not pointing next to nil.")
	}

	if list.Length != 2 {
		t.Errorf("Error: Length did not decrease from 3 to 2")
	}
}

func TestRemoveFromEndNoTail(t *testing.T) {
	list := list{nil, nil, 0}

	t.Log("Expecting to see an error when trying to removeFromEnd of an empty list.")

	err := list.removeFromEnd()

	t.Logf("Error produced: %+v", err)

	if err == nil {
		t.Errorf("Error: No error produced when running removeFromEnd")
	}

	if list.Tail != nil {
		t.Errorf("Error: Tail is not nil.")
	}

	if list.Length != 0 {
		t.Errorf("Error: Length is not 0.")
	}
}

func TestRemoveFromBeginning(t *testing.T) {
	node1 := node{1, nil, nil}
	node2 := node{2, nil, nil}
	node3 := node{3, nil, nil}

	node1.Next = &node2
	node2.Prev = &node1
	node2.Next = &node3
	node3.Prev = &node2

	list := list{&node1, &node3, 3}

	t.Log("Expecting the list to be shortened by one node, with the first node being node 2.")

	err := list.removeFromBeginning()

	t.Logf("Current head: %+v", list.Head)
	t.Logf("Node 2: %+v", node2)

	if err != nil {
		t.Error("Error produced when running removeFromBeginning.")
		t.Logf("Error: %+v", err)
	}

	if list.Head != &node2 {
		t.Error("Error: Head is not equal to node 2")
	}

	if list.Head.Prev != nil {
		t.Error("Error: Head is not pointing previous to nil")
	}

	if list.Length != 2 {
		t.Errorf("Error: Length did not decrease from 3 to 2")
	}
}

func TestRemoveFromBeginningNoHead(t *testing.T) {
	list := list{nil, nil, 0}

	t.Log("Expecting to see an error when trying to remove from the front of an empty list.")

	err := list.removeFromBeginning()

	if err == nil {
		t.Errorf("Error: err is nil after removeFromBeginning is run.")
	}

	if list.Head != nil {
		t.Errorf("Error: Head is not nil.")
	}

	if list.Length != 0 {
		t.Errorf("Error: Length is not 0.")
	}
}

func TestIsEmptyOnEmptyList(t *testing.T) {
	list := list{nil, nil, 0}

	t.Log("Expecting for the return from isEmpty to be true.")

	isListEmpty := list.isEmpty()

	if isListEmpty != true {
		t.Errorf("Error: isListEmpty is not true.")
	}
}

func TestIsEmptyOnNonEmptyList(t *testing.T) {
	node1 := node{1, nil, nil}
	node2 := node{2, nil, nil}
	node3 := node{3, nil, nil}

	node1.Next = &node2
	node2.Prev = &node1
	node2.Next = &node3
	node3.Prev = &node2

	list := list{&node1, &node3, 3}

	t.Log("Expecting for the return from isEmpty to be false.")

	isListEmpty := list.isEmpty()

	if isListEmpty != false {
		t.Errorf("Error: isListEmpty is true.")
	}
}

func TestRemoveNodeOfValue(t *testing.T) {
	node1 := node{1, nil, nil}
	node2 := node{2, nil, nil}
	node3 := node{3, nil, nil}

	node1.Next = &node2
	node2.Prev = &node1
	node2.Next = &node3
	node3.Prev = &node2

	list := list{&node1, &node3, 3}

	t.Log("Expecting for list length to be shortened and node 2 removed.")

	t.Logf("Initial node 1 next: %+v", node1.Next)
	t.Logf("Initial node 3 prev: %+v", node3.Prev)

	err := list.removeNodeOfValue(2)

	t.Logf("Final node 1 next: %+v", node1.Next)
	t.Logf("Final node 3 prev: %+v", node3.Prev)

	if node1.Next != &node3 {
		t.Errorf("Error: node 1 is not pointing next to node 3.")
	}

	if node3.Prev != &node1 {
		t.Errorf("Error: node 3 is not pointing prev to node 1.")
	}

	if list.Length != 2 {
		t.Errorf("Error: List length is not 2, length is %+v", list.Length)
	}

	if err != nil {
		t.Errorf("Error: Error is not nil, error is %+v", err)
	}
}

func TestRemoveNodeOfValueThatDoesNotExist(t *testing.T) {
	node1 := node{1, nil, nil}
	node2 := node{2, nil, nil}
	node3 := node{3, nil, nil}

	node1.Next = &node2
	node2.Prev = &node1
	node2.Next = &node3
	node3.Prev = &node2

	list := list{&node1, &node3, 3}

	t.Log("Expecting to see an error.")

	err := list.removeNodeOfValue(4)

	if err == nil {
		t.Error("Error is nil, it shouldn't be.")
	}

	if list.Length != 3 {
		t.Errorf("Error: Length is not 3, %+v", list.Length)
	}
}

func TestRemoveNodeOfValueFromEmptyList(t *testing.T) {
	list := list{nil, nil, 0}

	t.Log("Expecting to see an error.")

	err := list.removeNodeOfValue(1)

	if err == nil {
		t.Error("Error is nil, it shouldn't be.")
	}

	if list.Length != 0 {
		t.Errorf("Error: Length is not 0, %+v", list.Length)
	}
}

func TestRemoveGivenNode(t *testing.T) {
	node1 := node{1, nil, nil}
	node2 := node{2, nil, nil}
	node3 := node{3, nil, nil}

	node1.Next = &node2
	node2.Prev = &node1
	node2.Next = &node3
	node3.Prev = &node2

	list := list{&node1, &node3, 3}

	t.Log("Expecting to see list reduce in size because node 2 is removed.")

	t.Logf("Initial node 1 next: %+v", node1.Next)
	t.Logf("Initial node 3 prev: %+v", node3.Prev)

	err := list.removeGivenNode(&node2)

	t.Logf("Final node 1 next: %+v", node1.Next)
	t.Logf("Final node 3 prev: %+v", node3.Prev)

	if node1.Next != &node3 {
		t.Errorf("Error: node 1 is not pointing next to node 3.")
	}

	if node3.Prev != &node1 {
		t.Errorf("Error: node 3 is not pointing prev to node 1.")
	}

	if err != nil {
		t.Errorf("Error: err is not nil, %+v", err)
	}

	if list.Length != 2 {
		t.Errorf("Error: length is not 2, %+v", list.Length)
	}
}

func TestAddNodeAfter(t *testing.T) {
	node1 := node{1, nil, nil}
	node2 := node{2, nil, nil}
	node3 := node{3, nil, nil}

	node1.Next = &node2
	node2.Prev = &node1
	node2.Next = nil

	list := list{&node1, &node2, 2}

	t.Log("Expecting to grow the list by one, with the new tail being equal to node 3.")

	t.Logf("Initial List: %+v", list)
	t.Logf("Initial Tail: %+v", list.Tail)

	err := list.addNodeAfter(&node2, &node3)

	t.Logf("Final List: %+v", list)
	t.Logf("Final Tail: %+v", list.Tail)

	if list.Length != 3 {
		t.Errorf("Error: length is not 3. Length: %v", list.Length)
	}

	if list.Tail != &node3 {
		t.Errorf("Error: tail node is not equal to added node.")
	}

	if err != nil {
		t.Errorf("Error produced from addNodeAfter is not nil.")
	}
}

func TestAddNodeBefore(t *testing.T) {
	node1 := node{1, nil, nil}
	node2 := node{2, nil, nil}
	node3 := node{3, nil, nil}

	node1.Next = &node2
	node2.Prev = &node1
	node2.Next = nil

	list := list{&node1, &node2, 2}

	t.Log("Expecting to grow the list by one, with the new head being equal to node 3.")

	t.Logf("Initial List: %+v", list)
	t.Logf("Initial Tail: %+v", list.Head)

	err := list.addNodeBefore(&node1, &node3)

	t.Logf("Final List: %+v", list)
	t.Logf("Final Tail: %+v", list.Head)

	if list.Length != 3 {
		t.Errorf("Error: length is not 3. Length: %v", list.Length)
	}

	if list.Head != &node3 {
		t.Errorf("Error: Head node is not equal to added node.")
	}

	if err != nil {
		t.Errorf("Error produced from addNodeBefore is not nil.")
	}
}
