package main

import "fmt"

//given two strings write a method to decide if one is a permutation of the other

func getChars(input string) map[string]int {
	charMap := make(map[string]int)

	for _, value := range input {
		charMap[string(value)]++
	}

	return charMap
}

func compareMaps(map1 map[string]int, map2 map[string]int) bool {
	isPermutation := true

	for key, _ := range map1 {
		//fmt.Printf("FIRST: %v : %v and %v: %v\n", key, value, key, value)
		if map1[key] != map2[key] {
			isPermutation = false
			return isPermutation
		}
	}

	for key, _ := range map2 {
		//fmt.Printf("SECOND: %v : %v and %v: %v\n", key, map2[key], key, map1[key])
		if map2[key] != map1[key] {
			isPermutation = false
			return isPermutation
		}
	}

	return isPermutation
}

func main() {
	firstPerm := "abc"
	secondPerm := "bca"

	notPerm := "cbad"

	firstMap := getChars(firstPerm)
	secondMap := getChars(secondPerm)
	notMap := getChars(notPerm)

	fmt.Println(firstMap)
	fmt.Println(secondMap)
	fmt.Println(notMap)

	firstAndSecond := compareMaps(firstMap, secondMap)

	firstAndNot := compareMaps(firstMap, notMap)

	if firstAndSecond {
		fmt.Printf("%v is a permutation of %v\n", secondPerm, firstPerm)
	} else {
		fmt.Printf("%v is not a permutation of %v\n", secondPerm, firstPerm)
	}

	if firstAndNot {
		fmt.Printf("%v is a permutation of %v\n", notPerm, firstPerm)
	} else {
		fmt.Printf("%v is not a permutation of %v\n", notPerm, firstPerm)
	}
}
