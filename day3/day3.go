package main

import "fmt"
import "math"

const input = 312051

func main() {
	answerOne := SolvePartOne(input)
	fmt.Printf("Answer One: %d\n", answerOne)

	answerTwo := SolvePartTwo(input)
	fmt.Printf("Answer Two: %d\n", answerTwo)
}

func SolvePartOne(square int) int {
	sqrt := math.Sqrt(float64(square))
	floatLayer := (sqrt - 1.0) / 2.0
	layer := int(math.Ceil(floatLayer))

	segmentSize := int(float64(layer)*2.0 + 1.0)
	largestNumber := int(math.Pow(float64(segmentSize), 2.0))

	fourth := largestNumber
	third := fourth - segmentSize + 1
	second := third - segmentSize + 1
	first := second - segmentSize + 1

	halfway := (segmentSize - 1) / 2
	if square == first || square == second || square == third || square == fourth {
		return layer + halfway
	}

	var corner int
	switch {
	case square < first:
		corner = first
	case square > first && square < second:
		corner = second
	case square > second && square < third:
		corner = third
	case square > third && square < fourth:
		corner = fourth
	}

	middle := corner - halfway
	distance := math.Abs(float64(middle - square))
	return layer + int(distance)
}

type Direction int

const (
	Down  Direction = 0
	Up    Direction = 1
	Left  Direction = 2
	Right Direction = 3
)

type Grid struct {
	G map[int]map[int]int
}

func (grid Grid) get(x, y int) int {
	if _, okay := grid.G[x]; okay {
		val, _ := grid.G[x][y]
		return val
	}
	return 0
}

func (grid Grid) set(x, y, val int) {
	if _, okay := grid.G[x]; !okay {
		grid.G[x] = map[int]int{}
	}
	grid.G[x][y] = val
}

func SolvePartTwo(input int) int {
	grid := Grid{map[int]map[int]int{}}
	grid.set(0, 0, 1)

	currentDirection := Right
	currentX := 0
	currentY := 0

	for {
		switch currentDirection {
		case Down:
			currentY--
		case Up:
			currentY++
		case Left:
			currentX--
		case Right:
			currentX++
		}

		value := grid.get(currentX-1, currentY) +
			grid.get(currentX-1, currentY+1) +
			grid.get(currentX-1, currentY-1) +
			grid.get(currentX+1, currentY) +
			grid.get(currentX+1, currentY+1) +
			grid.get(currentX+1, currentY-1) +
			grid.get(currentX, currentY+1) +
			grid.get(currentX, currentY-1)


		if value > input {
			return value
		}

		grid.set(currentX, currentY, value)

		switch currentDirection {
		case Down:
			if grid.get(currentX + 1, currentY) == 0 {
				currentDirection = Right
			}
		case Up:
			if grid.get(currentX - 1, currentY) == 0 {
				currentDirection = Left
			}
		case Left:
			if grid.get(currentX, currentY - 1) == 0 {
				currentDirection = Down
			}
		case Right:
			if grid.get(currentX, currentY + 1) == 0 {
				currentDirection = Up
			}
		}

	}
}
