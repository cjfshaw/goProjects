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

func getLengths(input []int, isEven bool) (int, int) {
	var left, right int

	if isEven == true {
		left = len(input) / 2
		right = len(input) / 2
		return left, right
	} else {
		left = len(input) / 2
		right = (len(input) / 2) + 1
		return left, right
	}
}

func splitArray(input []int) ([]int, []int) {
	isEven := isArrayEven(input)
	leftLength, rightLength := getLengths(input, isEven)

	left := make([]int, leftLength)
	right := make([]int, rightLength)

	return left, right
}

func mergeSort(input []int) []int {

}

func main() {
	fmt.Println("Main ran.")
}

