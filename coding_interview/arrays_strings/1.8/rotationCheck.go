package main

import (
	"fmt"
	"sort"
	"strings"
)

//assume you have a method isSubstring which checks is one word is a substring of another
//given two strings write code to check if s2 is a rotation of s1 calling isSubstring only once
//example: erbottlewat is a rotation of waterbottle

//ideas

//sort each string, then run issubstring

//find the first match in s2, cut all the things from before it to behind it, then run issubstring

func checkSameCharacters() {
	s1 := "waterbottle"
	s2 := "erbottlewat"

	s1array := strings.Split(s1, "")
	sort.Strings(s1array)
	s2array := strings.Split(s2, "")
	sort.Strings(s2array)

	notSubstring := false

	for index, _ := range s1array {
		if s1array[index] != s2array[index] {
			notSubstring = true
		}
	}

	if notSubstring {
		fmt.Println("arrays are not equal")
	} else {
		fmt.Println("arrays are equal")
	}
}

func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"

	var startingIndex int

	for _, s1value := range s1 {
		for s2index, s2value := range s2 {
			//fmt.Printf("s1val: %s, s2val: %s\n", string(s1value), string(s2value))
			if s1value == s2value {
				startingIndex = s2index
				//fmt.Println("startingindex: ", startingIndex)
				break
			}
		}
		break
	}

	//fmt.Println("startingindex: ", startingIndex)
	var isMatch bool
	var counter, i int

	for i = startingIndex; i < len(s2); i++ {
		fmt.Printf("counter: %d, s1[counter]: %s - s2[i]: %s\n", counter, string(s1[counter]), string(s2[i]))
		for s1[counter] == s2[i] {
			fmt.Println("in if")
			isMatch = true
			fmt.Println("counter: ", counter)
			counter++
			fmt.Println("counter: ", counter)

			if i == len(s2)-1 {
				fmt.Printf("starting: %d, counter: %d, length of s2: %d\n", startingIndex, counter, len(s2))
				if (startingIndex + counter) != len(s2) {
					isMatch = false
				}
			}
		}
	}

	//fmt.Println(s2[startingIndex:])
	fmt.Println(isMatch)
	fmt.Println("s1 first sub: ", s1[:counter])
	fmt.Println("s1 second sub: ", s1[counter:])
	fmt.Println("s2 cut: ", s2[:startingIndex])

	if isMatch {
		//go version of isSubstring is contains
		isSubstring := strings.Contains(s1[counter:], s2[:startingIndex])

		if isSubstring {
			fmt.Println("is a subString")
		} else {
			fmt.Println("is not a subString")
		}
	}
}
