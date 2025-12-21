package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	n := problem1("input.txt")
	fmt.Println(n)

}

func problem1(filename string) int {
	lines, err := readInput(filename)
	if err != nil {
		log.Fatalf("error readInput: %q", err)
	}
	var total int
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[0]); c++ {
			if accessible(lines, r, c) {
				total++
			}
		}
	}
	return total
}

func accessible(lines [][]byte, row int, col int) bool {
	if lines[row][col] == '.' {
		return false
	}
	var rolls int
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}
			moveRow := r + row
			moveCol := c + col
			if moveRow >= len(lines) || moveCol >= len(lines[0]) || moveRow == -1 || moveCol == -1 {
				continue
			}
			if lines[moveRow][moveCol] == '@' {
				rolls++
			}
			if rolls > 3 {
				return false
			}
		}
	}
	return true
}

func readInput(input string) ([][]byte, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("error open file: %s: %q", input, err)
	}
	defer file.Close()

	var output [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, []byte(scanner.Text()))
	}
	return output, nil
}
