package main

//list struct contains head and tail pointers
//head and tail are node objects
//node struct contain value, pointer to prev node, pointer to next node
//you should be able to
//add a node (to the end)
//remove a node (with value)
//remove a node (given the node that comes before it or after it)
//find a node in a list
//determine if list is empty

import (
	"errors"
	"fmt"
)

type node struct {
	Value int
	Prev  *node
	Next  *node
}

type list struct {
	Head *node
	Tail *node
}

func (list *list) addToEnd(newNode *node) {
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		list.Tail.Next = newNode
		newNode.Prev = list.Tail
		list.Tail = newNode
	}
}

func (list *list) addToBeginning(newNode *node) {
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Next = list.Head
		list.Head.Prev = newNode
		list.Head = newNode
	}
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

func main() {
	fmt.Println("linked_list.go ran")
}
