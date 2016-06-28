/*Imagine a robot sitting on the upper left corner of an X by Y grid. The
robot can only move in two directions: right and down.  How many possible paths
are there for the robot to go from (0,0) to (X,Y)?

Follow up: Imagine certain spots are off-limits, such that the robot cannot
step on them.  Design an algorithm for the robot to find a path from the top
left to the bottom right.*/

package main

import (
	"fmt"
)

func robotPaths(currentX int, currentY int, maxX int, maxY int) int {
	//space complexity
	//time complexity
	currentPaths := 0
	if currentX == maxX && currentY == maxY {
		//fmt.Printf("in base case, x: %d, y= %d\n", currentX, currentY)
		return 1
	}
	if currentX < maxX {
		//fmt.Printf("in X, x: %d, y= %d\n", currentX, currentY)
		//currentX++
		currentPaths = currentPaths + robotPaths(currentX+1, currentY, maxX, maxY)
	}
	if currentY < maxY {
		//fmt.Printf("in Y, x: %d, y= %d\n", currentX, currentY)
		//currentY++
		currentPaths = currentPaths + robotPaths(currentX, currentY+1, maxX, maxY)
	}
	return currentPaths
}

func robotPathsBlockedSpaces(currentX int, currentY int, maxX int, maxY int) int {
	//space complexity
	//time complexity
	//fmt.Printf("in root, x: %d, y= %d\n", currentX, currentY)
	currentPaths := 0
	if currentX == maxX && currentY == maxY {
		//fmt.Printf("in base case, x: %d, y= %d\n", currentX, currentY)
		return 1
	}
	if currentX < maxX {
		//fmt.Printf("in X, x: %d, y= %d\n", currentX, currentY)
		//currentX++
		if currentX+1 == 1 && currentY == 1 {
			return 1
		} else {
			currentPaths = currentPaths + robotPathsBlockedSpaces(currentX+1, currentY, maxX, maxY)
		}
	}
	if currentY < maxY {
		//fmt.Printf("in Y, x: %d, y= %d\n", currentX, currentY)
		//currentY++
		if currentY+1 == 1 && currentX == 1 {
			return 1
		} else {
			currentPaths = currentPaths + robotPathsBlockedSpaces(currentX, currentY+1, maxX, maxY)
		}
	}
	return currentPaths
}

func main() {
	answer := robotPaths(0, 0, 2, 2)
	fmt.Println(answer)

	blockedSpacesAnswer := robotPathsBlockedSpaces(0, 0, 2, 2)
	fmt.Println(blockedSpacesAnswer)
}
