package main

import "testing"

func TestSolvePartOne(t *testing.T) {
	if 1 != SolvePartOne([]string{"aa bb cc dd ee"}) {
		t.Error()
	}

	if 0 != SolvePartOne([]string{"aa bb cc dd aa"}) {
		t.Error()
	}

	if 1 != SolvePartOne([]string{"aa bb cc dd aaa"}) {
		t.Error()
	}
}

func TestSolvePartTwo(t *testing.T) {
	if 1 != SolvePartTwo([]string {"abcde fghij"}) {
		t.Error()
	}

	if 0 != SolvePartTwo([]string {"abcde xyz ecdab"}) {
		t.Error()
	}

	if 1 != SolvePartTwo([]string {"a ab abc abd abf abj"}) {
		t.Error()
	}

	if 1 != SolvePartTwo([]string {"iiii oiii ooii oooi oooo"}) {
		t.Error()
	}

	if 0 != SolvePartTwo([]string {"oiii ioii iioi iiio"}) {
		t.Error()
	}
}