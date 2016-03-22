package main

import "testing"

type testArrays struct {
	EvenArray   []int
	OddArray    []int
	One         []int
	TwoUnsorted []int
	TwoSorted   []int
	EvenSorted  []int
	OddSorted   []int
}

func makeArrays() testArrays {
	Even := []int{7, 1, 8, 3}
	Odd := []int{7, 1, 4, 8, 3}
	One := []int{1}
	TwoUnsorted := []int{7, 1}
	TwoSorted := []int{1, 7}
	EvenSorted := []int{1, 7}
	OddSorted := []int{3, 4, 8}

	arrays := testArrays{EvenArray: Even, OddArray: Odd, One: One, TwoUnsorted: TwoUnsorted, TwoSorted: TwoSorted, EvenSorted: EvenSorted, OddSorted: OddSorted}

	return arrays
}

func TestIsArrayEvenEven(t *testing.T) {
	arrays := makeArrays()

	t.Log("Testing whether isArrayEven returns true if the array is even.")

	isEven := isArrayEven(arrays.EvenArray)

	if isEven != true {
		t.Errorf("Error: isEven is not true.\nisEven: %v\nLength:%v", isEven, len(arrays.EvenArray))
	}
}

func TestIsArrayEvenOdd(t *testing.T) {
	arrays := makeArrays()

	t.Log("Testing whether isArrayEven returns false if the array is odd.")

	isEven := isArrayEven(arrays.OddArray)

	if isEven == true {
		t.Errorf("Error: isEven is true.\nisEven: %v\nLength:%v", isEven, len(arrays.EvenArray))
	}
}

func getLengthsEven(t *testing.T) {
	arrays := makeArrays()

	t.Log("Testing whether getLengths returns proper lengths on an even array.")

	left, right, isEven := getLengths(arrays.EvenArray)

	if left != right {
		t.Errorf("Error: Array Lengths are not equal.\nLeft: %v\nRight: %v", left, right)
	}

	if left != 2 {
		t.Errorf("Error: Length of array is incorrect.\n Length: %v", left)
	}

	if isEven != true {
		t.Errorf("Error: isEven did not return true.\nisEven: %v", isEven)
	}
}

func getLengthsOdd(t *testing.T) {
	arrays := makeArrays()

	t.Log("Testing whether getLengths returns proper lengths on an odd array.")

	left, right, isEven := getLengths(arrays.EvenArray)

	if left == right {
		t.Errorf("Error: Array Lengths are equal.\nLeft: %v\nRight: %v", left, right)
	}

	if left != 2 && right != 3 {
		t.Errorf("Error: Length of arrays are incorrect.\n Length left: %v\nLength right: %v", left, right)
	}

	if isEven == true {
		t.Errorf("Error: isEven returned true.\nisEven: %v", isEven)
	}
}

func TestSplitArraysEven(t *testing.T) {
	arrays := makeArrays()

	t.Log("Testing whether split arrays returns two properly populated arrays when the input length is even.")

	left, right := splitArray(arrays.EvenArray)

	if len(left) != len(right) {
		t.Errorf("Error: Array Lengths are not equal.\nLeft: %v\nRight: %v", len(left), len(right))
	}

	if left[0] != 7 && left[1] != 1 {
		t.Errorf("Error: Values of left array are incorrect.\nArray: %v", left)
	}

	if right[0] != 8 && right[1] != 3 {
		t.Errorf("Error: Values of right array are incorrect.\nArray: %v", right)
	}
}

func TestSplitArraysOdd(t *testing.T) {
	arrays := makeArrays()

	t.Log("Testing whether split arrays returns two properly populated arrays when the input length is odd.")

	left, right := splitArray(arrays.OddArray)

	t.Logf("Left length: %v", len(left))
	t.Logf("Right length: %v", len(right))

	if len(left) == len(right) {
		t.Errorf("Error: Array Lengths are equal.\nLeft: %v\nRight: %v", len(left), len(right))
	}

	if left[0] != 7 && left[1] != 1 {
		t.Errorf("Error: Values of left array are incorrect.\nArray: %v", left)
	}

	if right[0] != 4 && right[1] != 8 && right[2] != 3 {
		t.Errorf("Error: Values of right array are incorrect.\nArray: %v", right)
	}
}

func TestSortArrayOfTwoOnSorted(t *testing.T) {
	arrays := makeArrays()

	t.Log("Testing whether sorting of a single array works when the array has two (sorted) values.")

	sorted := sortArrayOfTwo(arrays.TwoSorted)

	if sorted[0] != arrays.TwoSorted[0] && sorted[1] != arrays.TwoSorted[1] {
		t.Errorf("Error: Arrays are not equal.\nsorted: %v\noriginal: %v", sorted, arrays.TwoSorted)
	}

	if len(sorted) != len(arrays.TwoSorted) {
		t.Errorf("Error: Sorted array's length is different from original array's length.\nLength sorted: %v\nLength original: %v", len(sorted), len(arrays.TwoSorted))
	}
}

func TestSortArrayOfTwoOnUnsorted(t *testing.T) {
	arrays := makeArrays()

	t.Log("Testing whether sorting of a single array works when the array has two (sorted) values.")

	sorted := sortArrayOfTwo(arrays.TwoUnsorted)

	if sorted[0] != arrays.TwoSorted[0] && sorted[1] != arrays.TwoUnsorted[1] {
		t.Errorf("Error: Arrays are not equal.\nsorted: %v\noriginal: %v", sorted, arrays.TwoUnsorted)
	}

	if len(sorted) != len(arrays.TwoUnsorted) {
		t.Errorf("Error: Sorted array's length is different from original array's length.\nLength sorted: %v\nLength original: %v", len(sorted), len(arrays.TwoUnsorted))
	}
}

func TestSortSingleArrayOnOne(t *testing.T) {
	arrays := makeArrays()

	t.Log("Testing whether sorting of a single array works when the array has one value.")

	sorted, err := sortSingleArray(arrays.One)

	t.Log("Array has been sorted.")

	if sorted[0] != arrays.One[0] {
		t.Errorf("Error: Sorted array is not equal to arrays.One.\nSorted: %v\narrays.One: %v", sorted, arrays.One)
	}

	if len(sorted) != 1 {
		t.Errorf("Error: Length of sorted array is not one.\nLengthL %v", len(sorted))
	}

	if err != nil {
		t.Errorf("Error: err is not nil.\n%v", err)
	}
}

func TestSortingArrayOnMultiple(t *testing.T) {
	arrays := makeArrays()

	t.Log("Testing whether sorting of a single array works when the array has multiple values.")
	t.Logf("Starting Arrays:\narrays.One: %v\narrays.TwoSorted: %v", arrays.One, arrays.TwoSorted)

	newArray := mergeTwoArrays(arrays.One, arrays.TwoSorted)
	t.Logf("Merged Array: %v", newArray)

	for i := 1; i < len(newArray); i++ {
		if newArray[i] < newArray[i-1] {
			t.Errorf("Error: Array is not sorted.\nArray: %v", newArray)
		}
	}

	t.Logf("Starting Arrays:\narrays.EvenSorted: %v\narrays.OddSorted: %v", arrays.EvenSorted, arrays.OddSorted)

	newArray2 := mergeTwoArrays(arrays.EvenSorted, arrays.OddSorted)
	t.Logf("Merged Array: %v", newArray2)

	for i := 1; i < len(newArray2); i++ {
		if newArray2[i] < newArray2[i-1] {
			t.Errorf("Error: Array is not sorted.\nArray: %v", newArray2)
		}
	}
}
