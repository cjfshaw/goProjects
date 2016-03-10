package linked_list

//list struct contains head and tail pointers
//head and tail are node objects
//node struct contain value, pointer to prev node, pointer to next node
//you should be able to
//add a node (to the end) done
//add a node (to the beginning) done
//remove a node (from the end) done
//remove a node (from the beginning) done
//determine is list is empty done
//determine length of list done
//find a node in a list done
//remove a node (with value) done
//add a node after a given node done
//add a node before a given node done

import (
	"errors"
)

type node struct {
	Value int
	Prev  *node
	Next  *node
}

type list struct {
	Head   *node
	Tail   *node
	Length int
}

func (list *list) isEmpty() bool {
	var listIsEmpty bool

	if list.Length == 0 && list.Head == nil && list.Tail == nil {
		listIsEmpty = true
	}

	return listIsEmpty
}

func (list *list) addToEnd(newNode *node) {
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		list.Length++
	} else {
		list.Tail.Next = newNode
		newNode.Prev = list.Tail
		list.Tail = newNode
		list.Length++
	}
}

func (list *list) addToBeginning(newNode *node) {
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		list.Length++
	} else {
		newNode.Next = list.Head
		list.Head.Prev = newNode
		list.Head = newNode
		list.Length++
	}
}

func (list *list) removeFromEnd() error {
	var err error

	if list.Tail == nil {
		err = errors.New("Error in removeFromEnd, tail is nil.")
		return err
	} else {
		list.Tail.Prev.Next = nil
		list.Tail = list.Tail.Prev
		if list.Length != 0 {
			list.Length--
		}
	}

	return err
}

func (list *list) removeFromBeginning() error {
	var err error

	if list.Head == nil {
		err = errors.New("Error in removeFromBeginning, head is nil.")
		return err
	} else {
		list.Head.Next.Prev = nil
		list.Head = list.Head.Next
		if list.Length != 0 {
			list.Length--
		}
	}

	return err
}

func (list *list) findNodeByValue(nodeValue int) (*node, error) {
	var err error

	currentNode := list.Head

	for currentNode.Value != nodeValue {
		currentNode = currentNode.Next

		if currentNode == nil {
			err = errors.New("Error: No node found in this list with that value.")
			return currentNode, err
		}
	}

	return currentNode, err
}

func (list *list) removeNodeOfValue(nodeValue int) error {
	var err error

	if list.isEmpty() == true {
		err = errors.New("Error: List is empty, cannot remove a node from an empty list.")
		return err
	}

	foundNode, err := list.findNodeByValue(nodeValue)

	if err != nil {
		return err
	}

	foundNode.Prev.Next = foundNode.Next
	foundNode.Next.Prev = foundNode.Prev
	list.Length--

	return err
}

func (list *list) removeGivenNode(nodeToRemove *node) error {
	var err error

	if list.isEmpty() == true {
		err = errors.New("Error: can not removeGivenNode on empty list.")
		return err
	}

	//how can I verify that this node is in the list without iterating through the entire list?
	//I read how go does it, they include a field in their element struct that references a list
	//that is so fucking clever

	nodeToRemove.Prev.Next = nodeToRemove.Next
	nodeToRemove.Next.Prev = nodeToRemove.Prev
	list.Length--

	return err
}

func (list *list) addNodeAfter(nodeToFind *node, nodeToAdd *node) error {
	var err error

	if nodeToFind == nil || nodeToAdd == nil {
		err = errors.New("Error: An input node is nil.")
		return err
	}

	nodeToAdd.Prev = nodeToFind
	nodeToAdd.Next = nodeToFind.Next
	if nodeToFind == list.Tail {
		list.Tail = nodeToAdd
	} else {
		nodeToFind.Next.Prev = nodeToAdd
	}
	nodeToFind.Next = nodeToAdd
	list.Length++

	return err
}

func (list *list) addNodeBefore(nodeToFind *node, nodeToAdd *node) error {
	var err error

	if nodeToFind == nil || nodeToAdd == nil {
		err = errors.New("Error: An input node is nil.")
		return err
	}

	nodeToAdd.Prev = nodeToFind.Prev
	nodeToAdd.Next = nodeToFind
	if nodeToFind == list.Head {
		list.Head = nodeToAdd
	} else {
		nodeToFind.Prev.Next = nodeToAdd
	}
	nodeToFind.Prev = nodeToAdd
	list.Length++

	return err
}
