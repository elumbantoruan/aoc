package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n := problem1("input.txt")
	fmt.Println(n)

	n = problem2("input.txt")
	fmt.Println(n)
}

func problem1(input string) int {
	lines, err := readInput(input)
	if err != nil {
		log.Fatalf("error readInput: %q", err)
	}
	var (
		total    = 0
		maxByte1 = byte(' ')
		maxByte2 = byte(' ')
	)
	for _, line := range lines {
		length1 := len(line) - 1
		length2 := len(line)

		idxFirstRun := 0

		maxByte1 = line[0]
		for i := 1; i < length1; i++ {
			if line[i] > maxByte1 {
				maxByte1 = line[i]
				idxFirstRun = i
			}
		}
		maxByte2 = line[idxFirstRun+1]
		for i := idxFirstRun + 1; i < length2; i++ {
			if line[i] > byte(maxByte2) {
				maxByte2 = line[i]
			}
		}
		maxBytes := string(maxByte1) + string(maxByte2)
		num, _ := strconv.Atoi(maxBytes)
		total += num
	}

	return total
}

func problem2(input string) int {
	lines, err := readInput(input)
	if err != nil {
		log.Fatalf("error readInput: %q", err)
	}
	var (
		tokens   []byte
		position = -1
		result   int
	)
	for _, line := range lines {
		tokens = nil
		position = -1
		for len(tokens) < 12 {
			maxIndex := len(line) - 1 - (11 - len(tokens))
			data := getMaximizeData(line, position+1, maxIndex)
			tokens = append(tokens, data.Maximize)
			position = data.Index
		}
		tokenStr := string(tokens)
		n, _ := strconv.Atoi(tokenStr)
		result += n
	}
	return result
}

func getMaximizeData(line string, from int, to int) MaximizeIndex {
	var (
		maximize byte = ' '
		index    int  = 0
	)
	for i := from; i <= to; i++ {
		if line[i] > maximize {
			maximize = line[i]
			index = i
		}
	}
	return MaximizeIndex{Maximize: maximize, Index: index}
}

type MaximizeIndex struct {
	Maximize byte
	Index    int
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %q", filename, err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
