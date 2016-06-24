//A child is running up a staircase with n steps and can hop either 1 step,
//2 steps, or 3 steps at a time. Implement a method to determine how many
//possible ways the child can run up the stairs.
package main

import (
	"fmt"
)

func hopSteps(numSteps int) int {
	//fmt.Printf("in hopsteps, current numSteps %d\n", numSteps)
	if numSteps < 0 {
		//fmt.Println("invalid")
		return 0
	} else if numSteps == 0 {
		//fmt.Println("end of steps")
		return 1
	} else {
		fmt.Printf("sum: %d\n", hopSteps(numSteps-1)+hopSteps(numSteps-2)+hopSteps(numSteps-3))
		return hopSteps(numSteps-1) + hopSteps(numSteps-2) + hopSteps(numSteps-3)
	}
}

func main() {
	hopSteps(5)
}
