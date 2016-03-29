package main

import "testing"

type testElements struct {
	NoZeroMatrix [][]int
	ZerosMatrix  [][]int
}

func createElements() testElements {
	matrix := make([][]int, 0)
	row1 := []int{0, 1, 0}
	row2 := []int{3, 4, 5}
	matrix = append(matrix, row1)
	matrix = append(matrix, row2)

	matrix2 := make([][]int, 0)
	row3 := []int{6, 1, 2}
	row4 := []int{3, 4, 5}
	matrix2 = append(matrix2, row3)
	matrix2 = append(matrix2, row4)

	testElements := testElements{NoZeroMatrix: matrix2, ZerosMatrix: matrix}

	return testElements
}

func TestGetZeroRowAndColOnZerosMatrix(t *testing.T) {
	elements := createElements()

	row, col := getZeroRowAndCol(elements.ZerosMatrix)

	t.Log("Testing whether getZeroRowAndCol works on a matrix with zeros.")
	t.Logf("Rows and columns with zeros:\nRows: %v\nColumns: %v", row, col)

	if len(row) != 2 {
		t.Errorf("Error: More than 3 zeros found for rows.\n Rows: %v", row)
	}

	if len(col) != 2 {
		t.Errorf("Error: More than 3 zeros found for cols.\n Cols: %v", col)
	}

	if row[0] != 0 || row[1] != 0 {
		t.Errorf("Error: Incorrect values found for zero rows.\nRows: %v", row)
	}
	t.Logf("cold[1]: %v\n", col[1])
	if col[0] != 0 || col[1] != 2 {
		t.Errorf("Error: Incorrect values found for zero cols.\nCols: %v", col)
	}

}

func TestGetZeroRowAndColOnNoZerosMatrix(t *testing.T) {
	elements := createElements()

	row, col := getZeroRowAndCol(elements.NoZeroMatrix)

	t.Log("Testing whether getZeroRowAndCol works on a matrix without zeros.")

	if len(row) != 0 {
		t.Errorf("Error: More than 0 zeros found for rows.\n Rows: %v", row)
	}

	if len(col) != 0 {
		t.Errorf("Error: More than 0 zeros found for cols.\n Cols: %v", col)
	}
}

func TestZeroOutRowsOnZerosMatrix(t *testing.T) {
	elements := createElements()

	t.Log("Testing whether zeroOutRows works on a matrix with zeros.")
	t.Logf("Matrix to test: %v\n", elements.ZerosMatrix)

	row, col := getZeroRowAndCol(elements.ZerosMatrix)

	zeroRows := zeroOutRows(elements.ZerosMatrix, row)

	t.Logf("Rows and columns with zeros:\nRows: %v\nColumns: %v", row, col)
	t.Logf("Matrix after rows have been zerod: %v\n", zeroRows)

	if zeroRows[0][0] != 0 || zeroRows[0][1] != 0 || zeroRows[0][2] != 0 {
		t.Errorf("Error: Not all rows that should have been zeroed were.\nMatrix: %v", zeroRows)
	}

	if zeroRows[1][0] == 0 || zeroRows[1][1] == 0 || zeroRows[1][2] == 0 {
		t.Errorf("Error: Some values were changed to 0 that shouldn't have.\nMatrix: %v", zeroRows)
	}

}
