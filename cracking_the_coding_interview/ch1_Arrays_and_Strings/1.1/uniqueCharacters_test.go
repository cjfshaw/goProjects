package main

import (
	"testing"
)

const uniqueChars = "abcd"
const nonUniqueChars = "abcdd"

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

func TestUniqueCharactersMyOptimalSolution(t *testing.T) {
	myBool := true

	myBool = uniqueCharactersMyOptimalSolution(uniqueChars)

	if myBool != true {
		t.Error("Error: String with all unique characters returned false.")
	}

	myBool = uniqueCharactersMyOptimalSolution(nonUniqueChars)

	if myBool != false {
		t.Error("Error: String with non unique characters returned true.")
	}
}

func TestUniqueCharactersBookOptimalSolution(t *testing.T) {
	myBool := true

	myBool = uniqueCharactersBookOptimalSolution(uniqueChars)

	if myBool != true {
		t.Error("Error: String with all unique characters returned false.")
	}

	myBool = uniqueCharactersBookOptimalSolution(nonUniqueChars)

	if myBool != false {
		t.Error("Error: String with non unique characters returned true.")
	}
}
