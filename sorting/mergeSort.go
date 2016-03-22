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

/*func mergeSort(input []int) []int {
	left, right := splitArray(input)


	//does the splitting but nothing else
	//need to add comparing and then merging
}*/

func main() {
	fmt.Println("Main ran.")
}
