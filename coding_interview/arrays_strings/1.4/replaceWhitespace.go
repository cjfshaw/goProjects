package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Mr John Smith    "
	//var temp string

	array := strings.Split(input, "")

	fmt.Println(array)

	for j := 0; j < len(array); j++ {
		fmt.Println(j, array[j])
		if string(array[j]) == " " {
			for i := len(array) - 1; i > j; i-- {
				array[i] = array[i-1]
				fmt.Println(i, array)
			}
		}
		j += 2
	}

	fmt.Println(array)
}
