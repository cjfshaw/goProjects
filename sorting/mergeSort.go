package mergesort

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

func sortSingleArray(input []int) ([]int, error) {
	var err error

	/*if len(input) > 2 {
		err = fmt.Errorf("Error: Length of array in sortSingleArray is greater than two.\nLength: %v", len(input))
		return input, err
	}*/

	if len(input) == 1 {
		return input, err
	} else {
		sorted := sortArrayOfTwo(input)
		return sorted, err
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

/*func mergeSort(input []int) []int {
	left, right := splitArray(input)


	//does the splitting but nothing else
	//need to add comparing and then merging
}*/

func main() {
	fmt.Println("Main ran.")
}
