package main

import "fmt"

func getPivotAndIndex(input []int) (int, int) {
	pivotIndex := 0
	pivotValue := input[pivotIndex]

	return pivotValue, pivotIndex
}

func getLesserAndGreaterArrays(input []int, pivotIndex int) ([]int, []int) {
	lesser := make([]int, 0)
	greater := make([]int, 0)

	for i := 1; i < len(input); i++ {
		if input[i] <= input[pivotIndex] {
			lesser = append(lesser, input[i])
		} else {
			greater = append(greater, input[i])
		}
	}

	return lesser, greater
}

func quicksort(input []int) []int {
	_, pivotIndex := getPivotAndIndex(input)
	lesser, greater := getLesserAndGreaterArrays(input, pivotIndex)

	if len(lesser) > 0 {
		lesser = quicksort(lesser)
	}

	if len(greater) > 0 {
		greater = quicksort(greater)
	}

	//greater = quicksort(greater)

	sortedArray := make([]int, 0)
	sortedArray = append(sortedArray, lesser...)
	sortedArray = append(sortedArray, input[pivotIndex])
	sortedArray = append(sortedArray, greater...)

	return sortedArray
}

func main() {
	input := []int{3, 5, 9, 1, 1, 7}

	result := quicksort(input)

	fmt.Printf("Result: %v", result)
}
