package main

import "testing"

func TestSolvePartOne(t *testing.T) {
	a := SolvePartOne("1122")
	if a != 3 {
		t.Errorf("1122 should produce a sum of 3, but got %d\n", a)
	}

	b := SolvePartOne("1111")
	if b != 4 {
		t.Errorf("1111 should produce a sum of 4, but got %d\n", b)
	}

	c := SolvePartOne("1234")
	if c != 0 {
		t.Errorf("1234 should produce a sum of 0, but got %d\n", c)
	}

	d := SolvePartOne("91212129")
	if d != 9 {
		t.Errorf("91212129 should produce a sum of 9, but got %d\n", d)
	}
}

func TestSolvePartTwo(t *testing.T) {
	a := SolvePartTwo("1212")
	if a != 6 {
		t.Errorf("1212 should produce a sum of 6, but got %d\n", a)
	}

	b := SolvePartTwo("1221")
	if b != 0 {
		t.Errorf("1221 should produce a sum of 0, but got %d\n", b)
	}

	c := SolvePartTwo("123425")
	if c != 4 {
		t.Errorf("123425 should produce a sum of 4, but got %d\n", c)
	}

	d := SolvePartTwo("123123")
	if d != 12 {
		t.Errorf("123123 should produce a sum of 12, but got %d\n", d)
	}

	e := SolvePartTwo("12131415")
	if e != 4 {
		t.Errorf("12131415 should produce a sum of 4, but got %d\n", e)
	}
}
