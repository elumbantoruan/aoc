package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	num := problem1()
	fmt.Println(num)

	num = problem2()
	fmt.Println(num)
}

func problem1() int {
	lines, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("error readInput %q", err)
	}
	numbers, err := generateRangeNumbers(lines)
	if err != nil {
		log.Fatalf("error generateRangeNumbers %q", err)
	}
	total := countTotalInvalidIDs(numbers)
	return total
}

func problem2() int {
	lines, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("error readInput %q", err)
	}
	total := countTotalInvalidIDsTwo(lines)
	return total
}

func countTotalInvalidIDs(numbers []string) int {
	var total = 0
	for _, nums := range numbers {
		items := strings.Split(nums, ",")
		for _, num := range items {
			length := len(num)
			if length%2 == 1 {
				continue
			}
			firstHalf := num[:length/2]
			secondHalf := num[length/2:]
			if firstHalf == secondHalf {
				n, err := strconv.Atoi(num)
				if err != nil {
					continue
				}
				total += n
			}
		}
	}
	return total
}

func countTotalInvalidIDsTwo(numbers []string) int {
	var total = 0
	for _, nums := range numbers {
		items := strings.Split(nums, "-")
		begin, _ := strconv.Atoi(items[0])
		end, _ := strconv.Atoi(items[1])

		for number := begin; number <= end; number++ {
			token := strconv.Itoa(number)
			tokenLength := len(token)

			lenDigit := tokenLength / 2

			for i := 1; i <= lenDigit; i++ {
				if checkNumber(token, tokenLength, i) {
					total += number
					break
				}
			}
		}
	}
	return total
}

func checkNumber(number string, tokenLength int, lenDigit int) bool {
	if tokenLength == lenDigit {
		return false
	}
	if tokenLength%lenDigit != 0 {
		return false
	}
	pattern := number[:lenDigit]
	pointer := 0
	for {
		pointer += lenDigit
		if pointer >= tokenLength {
			return true
		}
		target := number[pointer : pointer+lenDigit]
		if pattern != target {
			break
		}
	}
	return false
}

func readInput(input string) ([]string, error) {
	file, err := os.Open(input)
	if err != nil {
		return nil, fmt.Errorf("error reading input file:%s, %q", input, err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		lines = append(lines, items...)
	}

	return lines, nil
}

func generateRangeNumbers(input []string) ([]string, error) {
	var lines []string

	for _, line := range input {
		items := strings.Split(line, "-")
		left, err := strconv.Atoi(items[0])
		if err != nil {
			return nil, fmt.Errorf("error converting val to number: %q", err)
		}
		right, err := strconv.Atoi(items[1])
		if err != nil {
			return nil, fmt.Errorf("error converting val to number: %q", err)
		}
		var (
			numbers []string
		)
		for left <= right {
			numbers = append(numbers, strconv.Itoa(left))
			left++
		}
		line := strings.Join(numbers, ",")
		lines = append(lines, line)
	}

	return lines, nil
}
