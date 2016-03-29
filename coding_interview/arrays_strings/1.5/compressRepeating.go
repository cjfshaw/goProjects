package main

import (
	"fmt"
	"strconv"
)

//compress a string to the following format: char numTimesRepeated
//if there are no repeats, return the original string
//example: "aabcccccaaa" -> a2b1c5a3
//example: "abcd" -> abcd

func main() {
	var currentCount, uniqueCount int
	input := "aabcccccaaa"
	//input := "abcd"
	var newString, countString string

	//fmt.Println("length: ", len(input))

	for i := 0; i < len(input)-1; i++ {
		/*fmt.Println("in if block")
		fmt.Printf("i: %d\n", i)
		fmt.Printf("input(i): %s\n", string(input[i]))*/
		if input[i] != input[i+1] {
			uniqueCount++
		}
		if input[i] == input[i+1] {
			//fmt.Printf("i: %s, i+1: %s\n", string(input[i]), string(input[i+1]))
			currentCount++
		} else {
			//fmt.Println("in else block")
			currentCount++
			countString = strconv.Itoa(currentCount)
			newString += string(input[i]) + countString
			currentCount = 0
		}
	}

	currentCount++
	countString = strconv.Itoa(currentCount)
	newString += string(input[len(input)-1]) + countString

	//fmt.Println(uniqueCount)

	if uniqueCount == len(input)-1 {
		fmt.Println(input)
	} else {
		fmt.Printf("%s\n", newString)
	}
}
