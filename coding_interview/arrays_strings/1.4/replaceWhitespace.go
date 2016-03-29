package main

import (
	"fmt"
	"strings"
)

//write a method to replace all spaces in a string with %20
//you can assume that the string is of the right length

func main() {
	input := "Mr John Smith    "
	//var temp string

	array := strings.Split(input, "")

	fmt.Println(array)

	for index, value := range array {
		fmt.Println(index, value)
		if string(value) == " " {
			array[index] = "%"
			for i := len(array) - 1; i > index; i-- {
				if i == index+1 {
					array[i] = "0"
				} else {
					array[i] = array[i-1]
				}
				fmt.Println(i, array)
			}
			for i := len(array) - 1; i > index; i-- {
				if i == index+1 {
					array[i] = "2"
				} else {
					array[i] = array[i-1]
				}
				fmt.Println(i, array)
			}
		}
		index += 2
	}

	fmt.Println(array)
}
