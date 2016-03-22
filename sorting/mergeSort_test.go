package mergesort

import "testing"

type testArrays struct {
	EvenArray []int
	OddArray  []int
}

func makeArrays() testArrays {
	Even := []int{7, 1, 8, 3}
	Odd := []int{7, 1, 4, 8, 3}

	arrays := testArrays{EvenArray: Even, OddArray: Odd}

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
