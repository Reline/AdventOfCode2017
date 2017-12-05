package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input := GetInput("day4_input.txt")

	answerOne := SolvePartOne(input)
	fmt.Printf("Answer One: %d\n", answerOne)

	answerTwo := SolvePartTwo(input)
	fmt.Printf("Answer Two: %d\n", answerTwo)
}

func SolvePartOne(input []string) int {
	valid := 0
Loop:
	for _, row := range input {
		words := strings.Fields(row)
		usedWords := map[string]bool{}
		for i, word := range words {
			if _, ok := usedWords[word]; ok {
				continue Loop
			}
			usedWords[words[i]] = true
		}
		valid++
	}
	return valid
}

func SolvePartTwo(input []string) int {
	valid := 0
	Loop:
	for _, row := range input {
		words := strings.Fields(row)
		usedWords := []string{}
		for i, word := range words {
			for _, used := range usedWords {
				if isAnagram(word, used) {
					continue Loop
				}
			}
			usedWords = append(usedWords, words[i])
		}
		valid++
	}
	return valid
}

func GetInput(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(b), "\n")
}

func isAnagram(first, second string) bool {
	if len(first) != len(second) {
		return false
	}

	f := make(map[rune]int, len(first))
	for _, r := range first {
		if _, ok := f[r]; !ok {
			f[r] = 1
		} else {
			f[r]++
		}
	}

	s := make(map[rune]int, len(second))
	for _, r := range second {
		if _, ok := s[r]; !ok {
			s[r] = 1
		} else {
			s[r]++
		}
	}

	for k, v := range f {
		if s[k] != v {
			return false
		}
	}

	return true
}