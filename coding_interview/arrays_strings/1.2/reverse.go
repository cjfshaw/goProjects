package main

import "fmt"

//implement a function void reverse (char *str) which reverses a null-terminated string

func main() {
	input := "a string"

	reversed := ""

	for i := len(input) - 1; i > 0; i-- {
		reversed += string(input[i])
	}

	fmt.Println(reversed)
}
