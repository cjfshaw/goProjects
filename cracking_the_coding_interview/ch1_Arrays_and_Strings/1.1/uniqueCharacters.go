//Implement an algorithm to determine if a string has all unique characters.
//What if you can not use any additional data structures?

package main

import (
	"sort"
	"strings"
)

func uniqueCharactersFirstSolution(input string) bool {
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

	inputArray := strings.Fields(input)

	sort.Strings(inputArray)

	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			return false
		}
	}

	return true
}

func uniqueCharactersOptimalSolution(input string) bool {
	//don't understand this solution, link to explanations here:
	//http://stackoverflow.com/questions/9141830/explain-the-use-of-a-bit-vector-for-determining-if-all-characters-are-unique
	//concepts: bit manipulation
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
