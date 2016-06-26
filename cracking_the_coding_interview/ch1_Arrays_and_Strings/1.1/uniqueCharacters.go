//Implement an algorithm to determine if a string has all unique characters.
//What if you can not use any additional data structures?

package main

import (
	"sort"
	"strings"
)

func uniqueCharactersFirstSolution(input string) bool {
	//space complexity: O(n)
	//time complexity: O(n)
	characterMap := make(map[string]int, 0)

	for _, value := range input {
		characterMap[string(value)]++
		if characterMap[string(value)] > 1 {
			return false
		}
	}

	return true
}

func uniqueCharactersNoDataStructuresNaive(input string) bool {
	//space complexity: O(n)
	//time complexity: O(n^2)
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}

	return true
}

func uniqueCharactersNoDataStructuresNaive2(input string) bool {
	//This solution says "I will sort the input, then compare each character to it's neighbor."
	//If I never find a character equal to it's neighbor then there are no duplicates.

	//space complexity: O(n)
	//time complexity: O(n)
	inputArray := strings.Fields(input)

	sort.Strings(inputArray)

	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return false
		}
	}

	return true
}

func uniqueCharactersBookOptimalSolution(input string) bool {
	//don't understand this solution, link to explanations here:
	//http://stackoverflow.com/questions/9141830/explain-the-use-of-a-bit-vector-for-determining-if-all-characters-are-unique
	//concepts: bit manipulation
	//use the bits in an int as a map, with 0 being you havent seen this before
	//and 1 being you have seen it
	var checker int
	var val uint

	for i := 0; i < len(input); i++ {
		val = uint(input[i] - byte('a'))
		if (checker & (1 << val)) > 0 {
			return false
		}
		checker |= (1 << val)
	}
	return true
}

func setBit(bitStore *int, position uint) {
	//shift 1 over by x bits, where x is the position variable
	//or this with bitstore to set the bit at position to 1
	*bitStore |= (1 << position)
}

func isBitSet(bitStore *int, position uint) bool {
	//and the bitstore variable with the bit at position
	value := *bitStore & (1 << position)
	return (value > 0)
}

func uniqueCharactersMyOptimalSolution(input string) bool {
	//using an int as a map here
	//if a bit is 0, you havent seen the corresponding character before
	//set it and keep iterating
	//if a bit is 1, you have seen it before, return false

	//space complexity: O(n)
	//time complexity: O(n)
	bitStore := 0

	for i := 0; i < len(input); i++ {
		switch {
		case string(input[i]) == "a":
			if isBitSet(&bitStore, 0) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 0)
			}
		case string(input[i]) == "b":
			if isBitSet(&bitStore, 1) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 1)
			}
		case string(input[i]) == "c":
			if isBitSet(&bitStore, 2) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 2)
			}
		case string(input[i]) == "d":
			if isBitSet(&bitStore, 3) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 3)
			}
		case string(input[i]) == "e":
			if isBitSet(&bitStore, 4) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 4)
			}
		case string(input[i]) == "f":
			if isBitSet(&bitStore, 5) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 5)
			}
		case string(input[i]) == "g":
			if isBitSet(&bitStore, 6) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 6)
			}
		case string(input[i]) == "h":
			if isBitSet(&bitStore, 7) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 7)
			}
		case string(input[i]) == "i":
			if isBitSet(&bitStore, 8) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 8)
			}
		case string(input[i]) == "j":
			if isBitSet(&bitStore, 9) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 9)
			}
		case string(input[i]) == "k":
			if isBitSet(&bitStore, 10) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 10)
			}
		case string(input[i]) == "l":
			if isBitSet(&bitStore, 11) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 11)
			}
		case string(input[i]) == "m":
			if isBitSet(&bitStore, 12) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 12)
			}
		case string(input[i]) == "n":
			if isBitSet(&bitStore, 13) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 13)
			}
		case string(input[i]) == "o":
			if isBitSet(&bitStore, 14) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 14)
			}
		case string(input[i]) == "p":
			if isBitSet(&bitStore, 15) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 15)
			}
		case string(input[i]) == "q":
			if isBitSet(&bitStore, 16) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 16)
			}
		case string(input[i]) == "r":
			if isBitSet(&bitStore, 17) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 17)
			}
		case string(input[i]) == "s":
			if isBitSet(&bitStore, 18) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 18)
			}
		case string(input[i]) == "t":
			if isBitSet(&bitStore, 19) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 19)
			}
		case string(input[i]) == "u":
			if isBitSet(&bitStore, 20) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 20)
			}
		case string(input[i]) == "v":
			if isBitSet(&bitStore, 21) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 21)
			}
		case string(input[i]) == "w":
			if isBitSet(&bitStore, 22) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 22)
			}
		case string(input[i]) == "x":
			if isBitSet(&bitStore, 23) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 23)
			}
		case string(input[i]) == "y":
			if isBitSet(&bitStore, 24) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 24)
			}
		case string(input[i]) == "z":
			if isBitSet(&bitStore, 25) { //corresponding bit is 0 set it to 1, if it is 1 return false
				return false
			} else {
				setBit(&bitStore, 25)
			}
		}

	}
	return true
}
