package main

import "fmt"
import "io/ioutil"
import "log"
import "strings"
import "strconv"

func main() {
	input := GetInput("day2_input.txt")

	answerOne := SolvePartOne(input)
	fmt.Printf("Answer One: %d\n", answerOne)

	answerTwo := SolvePartTwo(input)
	fmt.Printf("Answer Two: %d\n", answerTwo)
}

func GetInput(filename string) [][]int {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	input := [][]int{}
	rows := strings.Split(string(b), "\n")
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		cols := strings.Fields(row)
		val := []int{}
		for _, col := range cols {
			v, _ := strconv.Atoi(col)
			val = append(val, v)
		}
		input = append(input, val)
	}
	return input
}

func SolvePartOne(input [][]int) int {
	checksum := 0
	for _, r := range input {
		low := 1000000
		high := 0
		for i := range r {
			if r[i] < low {
				low = r[i]
			}
			if r[i] > high {
				high = r[i]
			}
		}
		checksum += high - low
	}
	return checksum
}

func SolvePartTwo(input [][]int) int {
	checksum := 0
	for _, row := range input {
	Loop:
		for _, col := range row {
			for _, c := range row {
				if col == c {
					continue
				}
				if col%c == 0 {
					checksum += col / c
					break Loop
				}
			}
		}
	}
	return checksum
}
