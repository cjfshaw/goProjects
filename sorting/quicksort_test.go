package main

import "testing"

func TestGetPivotAndIndex(t *testing.T) {
	arrays := makeArrays()

	pivotValue, pivotIndex := getPivotAndIndex(arrays.OddArray)

	if pivotValue != arrays.OddArray[0] {
		t.Errorf("Error: Pivot value is incorrect. PivotValue: %v\n", pivotValue)
	}

	if pivotIndex != 0 {
		t.Errorf("Error: Pivot index is incorrect. PivotIndex: %v\n", pivotIndex)
	}
}

func TestGetArrays(t *testing.T) {
	arrays := makeArrays()

	pivotValue, pivotIndex := getPivotAndIndex(arrays.OddArray)
	lesser, greater := getLesserAndGreaterArrays(arrays.OddArray, pivotIndex)

	for i := 0; i < len(lesser); i++ {
		if lesser[i] > pivotValue {
			t.Errorf("Error: Lesser array contains a value greater than pivot value.\nLesser: %v\nPivot: %v", lesser, pivotValue)
		}
	}

	for j := 0; j < len(greater); j++ {
		if greater[j] < pivotValue {
			t.Errorf("Error: Greater array contains a value lesser than pivot value.\nGreater: %v\nPivot: %v", greater, pivotValue)
		}
	}
}

func TestQuicksort(t *testing.T) {
	arrays := makeArrays()

	sorted := arrays.OddArray

	if sorted[0] != 1 && sorted[1] != 3 && sorted[2] != 4 && sorted[3] != 7 && sorted[4] != 8 {
		t.Errorf("Error: Array is not sorted.\nArray: %v", sorted)
	}
}
