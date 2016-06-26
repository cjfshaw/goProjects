//A child is running up a staircase with n steps and can hop either 1 step,
//2 steps, or 3 steps at a time. Implement a method to determine how many
//possible ways the child can run up the stairs.
package main

import (
	"fmt"
)

func hopSteps(numSteps int) int {
	if numSteps < 0 {
		return 0
	} else if numSteps == 0 {
		return 1
	} else {
		return hopSteps(numSteps-1) + hopSteps(numSteps-2) + hopSteps(numSteps-3)
	}
}

func hopStepsDynamic(numSteps int, stepsMap *map[int]int) int {
	if numSteps < 0 {
		return 0
	} else if numSteps == 0 {
		return 1
	} else if (*stepsMap)[numSteps] != 0 {
		return (*stepsMap)[numSteps]
	} else {
		(*stepsMap)[numSteps] = hopStepsDynamic(numSteps-1, stepsMap) + hopStepsDynamic(numSteps-2, stepsMap) + hopStepsDynamic(numSteps-3, stepsMap)
		return (*stepsMap)[numSteps]
	}
}

func main() {
	stepsMap := make(map[int]int, 0)

	hops := hopStepsDynamic(15, &stepsMap)
	fmt.Println(hops)
}
