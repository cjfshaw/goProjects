package main

import (
	"fmt"
	"sort"
)

//write an algorithm such that if an element in a MxN matrix is 0, that you set its row and column to 0

func getZeroRowAndCol(matrix [][]int) ([]int, []int) {
	row := make([]int, 0)
	col := make([]int, 0)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row = append(row, i)
				col = append(col, j)
			}
		}
	}

	sort.Ints(row)
	sort.Ints(col)

	return row, col
}

func zeroOutRows(matrix [][]int, rows []int) [][]int {
	for _, value := range rows {
		for i := 0; i < len(matrix); i++ {
			if i == value {
				for j := 0; j < len(matrix[i]); j++ {
					matrix[i][j] = 0
				}
			}
		}
	}

	return matrix
}

func zeroOutCols(matrix [][]int, cols []int) [][]int {
	for _, value := range cols {
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				if j == value {
					matrix[i][j] = 0
				}
			}
		}
	}

	return matrix
}

func main() {
	fmt.Println("main running")
	matrix := make([][]int, 0)
	row1 := []int{0, 1, 2}
	row2 := []int{3, 4, 5}
	matrix = append(matrix, row1)
	matrix = append(matrix, row2)

	rows, cols := getZeroRowAndCol(matrix)

	matrix = zeroOutRows(matrix, rows)
	matrix = zeroOutCols(matrix, cols)

	fmt.Printf("Matrix: %v\n", matrix)
}
