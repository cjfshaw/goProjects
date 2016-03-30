package main

import (
	"container/list"
	"fmt"
)

//write code to remove duplicates from an unsorted linked list
//how would you solve this problem if a temporary buffer was not allowed?

func main() {
	fmt.Println("main running")

	myList := list.New()
	fmt.Println("mylist.New", myList)

	/*elem1 := 1
	elem2 := 3
	elem3 := 1*/

	myList.PushBack(1)
	fmt.Println("mylist push el1", myList)
	myList.PushBack(3)
	//fmt.Println("mylist push el2", myList)
	myList.PushBack(1)
	//fmt.Println("mylist push el3", myList)
	myList.PushBack(1)
	myList.PushBack(1)
	myList.PushBack(1)
	myList.PushBack(3)
	fmt.Println("mylist push el6", myList)

	//while the currentNode.Next() is not nil
	//compare this node to every other node in the list
	//while THIS node.NExt() is not nil
	//do comparison and if it fits do action
	//increment
	//increment this currentNode

	//fmt.Println("mylist", myList)
	/*fmt.Println("el1", el1)
	fmt.Println("el2", el2)
	fmt.Println("el3", el3)*/

	/*startingNode := myList.Front()

	for startingNode.Next() != nil {
		//startingValue := startingNode.Value
		fmt.Println("in startingNode.Next() loop")
		back := myList.Back()
		currentValue := startingNode.Next()
		tempNode := currentValue.Prev()

		for startingNode.Next() != back {
			fmt.Println("in startingNode.Next != back loop")
			fmt.Println("startingNode: ", startingNode.Value)
			fmt.Println("currentnode value: ", currentValue.Value)
			if currentValue.Value == startingNode.Value {
				myList.Remove(currentValue)
				fmt.Println("directly after removal...")
				fmt.Println("mylist", myList)
				fmt.Println("currentnode prev: ", currentValue.Prev())
				fmt.Println("currentnode next: ", currentValue.Next())
				currentValue = tempNode
			} //else {
			fmt.Println("mylist", myList)
			fmt.Println("currentnode prev: ", currentValue.Prev())
			fmt.Println("currentnode next: ", currentValue.Next())
			currentValue = currentValue.Next()
			//}
		}
		startingNode = startingNode.Next()
	}

	fmt.Println("mylist push el6", myList)*/

	/*for index, value := range myList {
		for index2, value2 := range myList[index:] {
			if myList.value.Value == myList.value2.Value {
				myList.Remove(myList.value2)
			}
		}
	}*/
}
