package main

import (
	"testing"
)

const uniqueChars = "abcd1234!@#$"
const nonUniqueChars = "abcd1234!@#$$"

func TestUniqueCharactersFirstSolution(t *testing.T) {
	myBool := true

	myBool = uniqueCharactersFirstSolution(uniqueChars)

	if myBool != true {
		t.Error("Error: String with all unique characters returned false.")
	}

	myBool = uniqueCharactersFirstSolution(nonUniqueChars)

	if myBool != false {
		t.Error("Error: String with non unique characters returned true.")
	}
}

func TestUniqueCharactersNoDataStructuresNaive(t *testing.T) {
	myBool := true

	myBool = uniqueCharactersNoDataStructuresNaive(uniqueChars)

	if myBool != true {
		t.Error("Error: String with all unique characters returned false.")
	}

	myBool = uniqueCharactersNoDataStructuresNaive(nonUniqueChars)

	if myBool != false {
		t.Error("Error: String with non unique characters returned true.")
	}
}

func TestUniqueCharactersNoDataStructuresNaive2(t *testing.T) {
	myBool := true

	myBool = uniqueCharactersNoDataStructuresNaive2(uniqueChars)

	if myBool != true {
		t.Error("Error: String with all unique characters returned false.")
	}

	myBool = uniqueCharactersNoDataStructuresNaive2(nonUniqueChars)

	if myBool != false {
		t.Error("Error: String with non unique characters returned true.")
	}
}

func TestUniqueCharactersOptimalSolution(t *testing.T) {
	myBool := true

	myBool = uniqueCharactersOptimalSolution(uniqueChars)

	if myBool != true {
		t.Error("Error: String with all unique characters returned false.")
	}

	myBool = uniqueCharactersOptimalSolution(nonUniqueChars)

	if myBool != false {
		t.Error("Error: String with non unique characters returned true.")
	}
}
