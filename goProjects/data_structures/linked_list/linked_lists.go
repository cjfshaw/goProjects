package main

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
//remove a node (with value)
//remove a node (given the node that comes before it or after it)
//add a node after a given node
//add a node before a given node
//remove a node after a given node
//remove a node before a given node
//determine if list is empty

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
