package main

import "fmt"

func main() {
	input := "a string"

	reversed := ""

	for i := len(input) - 1; i > 0; i-- {
		reversed += string(input[i])
	}

	fmt.Println(reversed)
}
