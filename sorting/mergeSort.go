package main

// /import "fmt"

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
		if lowestLeft < len(left) && lowestRight < len(right) {
			currentLeft := left[lowestLeft]
			currentRight := right[lowestRight]

			if lowestLeft < len(left) && currentLeft <= currentRight {
				newArray[i] = currentLeft
				lowestLeft++
			} else if lowestRight < len(right) {
				newArray[i] = currentRight
				lowestRight++
			}
		} else if lowestLeft < len(left) {
			currentLeft := left[lowestLeft]
			newArray[i] = currentLeft
			lowestLeft++
		} else if lowestRight < len(right) {
			currentRight := right[lowestRight]
			newArray[i] = currentRight
			lowestRight++
		}
	}
	return newArray
}

func mergeSort(input []int) []int {
	if len(input) == 1 {
		return input
	} else {
		left, right := splitArray(input)
		left = mergeSort(left)
		right = mergeSort(right)
		sorted := mergeTwoArrays(left, right)
		return sorted
	}
}

/*func main() {
	fmt.Println("Main Started.")

	inputArray := []int{0, 7, 1, 4, 8, 3, 3, -1}

	sortedArray := mergeSort(inputArray)

	fmt.Println("Sorted Array: ", sortedArray)
}*/
