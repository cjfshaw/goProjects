package main

import "fmt"

func isArrayEven(input []int) bool {
	var isEven bool

	length := len(input)

	remainder := length % 2

	if remainder == 0 {
		isEven = true
		return isEven
	} else {
		isEven = false
		return isEven
	}
}

func getLengths(input []int) (int, int, bool) {
	var left, right int

	isEven := isArrayEven(input)

	if isEven == true {
		left = len(input) / 2
		right = len(input) / 2
		return left, right, isEven
	} else {
		left = len(input) / 2
		right = (len(input) / 2) + 1
		return left, right, isEven
	}
}

func splitArray(input []int) ([]int, []int) {
	leftLength, rightLength, isEven := getLengths(input)

	left := make([]int, leftLength)
	right := make([]int, rightLength)

	if isEven == true {
		copy(left, input[:leftLength])
		copy(right, input[rightLength:])
	} else {
		copy(left, input[:leftLength])
		copy(right, input[rightLength-1:])
	}

	return left, right
}

func sortArrayOfTwo(input []int) []int {
	if input[0] > input[1] {
		temp := input[1]
		input[1] = input[0]
		input[0] = temp
		return input
	} else {
		return input
	}
}

func sortSingleArray(input []int) []int {
	if len(input) == 1 {
		return input
	} else {
		sorted := sortArrayOfTwo(input)
		return sorted
	}
}

func mergeTwoArrays(left []int, right []int) []int {
	newArray := make([]int, len(left)+len(right))

	lowestLeft := 0
	lowestRight := 0

	for i := 0; i < len(newArray); i++ {
		if lowestLeft < len(left) && left[lowestLeft] <= right[lowestRight] {
			newArray[i] = left[lowestLeft]
			lowestLeft++
		} else if lowestRight < len(right) {
			newArray[i] = right[lowestRight]
			lowestRight++
		}
	}

	return newArray
}

func mergeSort(input []int) []int {
	//break array, check length, if its long enough keep calling
	//once you get returns then mergeTwoArrays
	if len(input) == 1 {
		return input
	} else {
		left, right := splitArray(input)
		left = mergeSort(left)
		right = mergeSort(right)
		sorted := mergeTwoArrays(left, right)
		return sorted
		/*if len(left) > 2 {
			sortedLeft := mergeSort(left)
		} else {
			sortedLeft := sortSingleArray(left)
		}
		if len(right) > 2 {
			sortedRight := mergeSort(right)
		} else {
			sortedRight := sortSingleArray(right)
		}
		sorted := mergeTwoArrays(sortedLeft, sortedRight)
		return sorted*/
	}
}

func main() {
	fmt.Println("Main Started.")

	inputArray := []int{7, 1, 4, 8, 3}

	sortedArray := mergeSort(inputArray)

	fmt.Println("Sorted Array: ", sortedArray)
}
