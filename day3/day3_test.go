package main

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	a := SolvePartOne(1)
	if a != 0 {
		t.Errorf("Data from square 1 is carried 0 steps, since it's at the access port. Got: %d\n", a)
	}

	b := SolvePartOne(12)
	if b != 3 {
		t.Errorf("Data from square 12 is carried 3 steps, such as: down, left, left. Got: %d\n", b)
	}

	c := SolvePartOne(23)
	if c != 2 {
		t.Errorf("Data from square 23 is carried only 2 steps: up twice. Got: %d\n", c)
	}

	d := SolvePartOne(1024)
	if d != 31 {
		t.Errorf("Data from square 1024 must be carried 31 steps. Got: %d\n", d)
	}
}

func TestSolvePartTwo(t *testing.T) {
	a := SolvePartTwo(1)
	if a != 2 {
		t.Errorf("Square 1 starts with the value 1, Got %d\n", a)
	}

	b := SolvePartTwo(2)
	if b != 4 {
		t.Errorf("Square 2 has only one adjacent filled square (with value 1), so it also stores 1, Got %d\n", b)
	}

	c := SolvePartTwo(3)
	if c != 4 {
		t.Errorf("Square 3 has both of the above squares as neighbors and stores the sum of their values, 2, Got %d\n", c)
	}

	d := SolvePartTwo(4)
	if d != 5 {
		t.Errorf("Square 4 has all three of the aforementioned squares as neighbors and stores the sum of their values, 4, Got %d\n", d)
	}

	e := SolvePartTwo(5)
	if e != 10 {
		t.Errorf("Square 5 only has the first and fourth squares as neighbors, so it gets the value 5, Got %d\n", e)
	}
}
