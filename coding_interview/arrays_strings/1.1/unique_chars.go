package main

import "fmt"

func countCharsWithMap(input string) map[string]int {
	charCount := make(map[string]int)

	for _, value := range input {
		charCount[string(value)]++
	}

	return charCount
}

func areCharsUniqueWithMap(input map[string]int) bool {
	isUnique := true

	for _, value := range input {
		if value != 1 {
			isUnique = false
			return isUnique
		}
	}

	return isUnique
}

func countCharsWithoutMap(input string) bool {
	isUnique := true

	for firstIndex, firstValue := range input {
		for secondIndex, secondValue := range input {
			//fmt.Printf("first: %v, second: %v\n", string(firstValue), string(secondValue))
			if firstValue == secondValue && firstIndex != secondIndex {
				isUnique = false
			}
		}
	}

	return isUnique
}

func main() {
	uniqueInput := "this one"

	repeatedInput := "aa"

	mapMap := countCharsWithMap(uniqueInput)
	mapIsUnique := areCharsUniqueWithMap(mapMap)
	if mapIsUnique {
		fmt.Println("All characters are unique.")
	} else {
		fmt.Println("All characters are not unique.")
	}

	mapMap2 := countCharsWithMap(repeatedInput)
	mapIsUnique2 := areCharsUniqueWithMap(mapMap2)
	if mapIsUnique2 {
		fmt.Println("All characters are unique.")
	} else {
		fmt.Println("All characters are not unique.")
	}

	mapMap3 := countCharsWithoutMap(uniqueInput)
	if mapMap3 {
		fmt.Println("All characters are unique.")
	} else {
		fmt.Println("All characters are not unique.")
	}

	mapMap4 := countCharsWithoutMap(repeatedInput)
	if mapMap4 {
		fmt.Println("All characters are unique.")
	} else {
		fmt.Println("All characters are not unique.")
	}
}
