/*A magic index in an array A[0..n-1] is defined to be an index such that
A[i] = i.  Given a sorted array of distinct integers, write a method to find a
magic index, if one exists, in an array A.
FOLLOW UP
What if the values are not distinct?
*/

package main

import (
	"fmt"
)

func findMagicIndicesIterative(array []int) {

	for i := 0; i < len(array); i++ {
		if array[i] == i {
			fmt.Printf("%d, ", i)
		}
	}
	fmt.Println()
}

func findMagicIndicesRecursive(index int, array []int) int {
	if array[index] > index {
		if len(array)-1 >= array[index] {
			return findMagicIndicesRecursive(array[index], array)
		} else {
			return -1
		}
	} else if array[index] == index {
		return index
	} else {
		return findMagicIndicesRecursive(index+1, array)
	}
}

func main() {
	noMagicIndices := []int{1, 2, 3, 4}
	magicIndices := []int{0, 2, 2, 4, 5}

	indices := findMagicIndicesRecursive(0, noMagicIndices)
	fmt.Println("no indices: ", indices)

	moreIndices := findMagicIndicesRecursive(0, magicIndices)
	fmt.Println("indices: ", moreIndices)
}
