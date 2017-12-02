package main

import "testing"
import "reflect"

func TestGetInput(t *testing.T) {
	expected := [][]int{
		{5, 1, 9, 5},
		{7, 5, 3},
		{2, 4, 6, 8},
	}

	actual := GetInput("day2_testinput1.txt")

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected: %v, Actual: %v\n", expected, actual)
	}
}

func TestSolvePartOne(t *testing.T) {
	input := GetInput("day2_testinput1.txt")

	checksum := SolvePartOne(input)
	if checksum != 18 {
		t.Fatalf("Expected: 18, Got: %d\n", checksum)
	}
}

func TestSolvePartTwo(t *testing.T) {
	input := GetInput("day2_testinput2.txt")

	checksum := SolvePartTwo(input)
	if checksum != 9 {
		t.Fatalf("Expected: 9, Got: %d\n", checksum)
	}
}
