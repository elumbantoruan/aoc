package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	v := problem1("input.txt")
	fmt.Println(v)

	v = problem2("input.txt")
	fmt.Println(v)
}

func problem1(filename string) int {
	rotations, err := readInput(filename)
	if err != nil {
		log.Fatalf("err from reading input file: %v", err)
	}
	var (
		point   = 50
		counter = 0
	)
	for _, rotation := range rotations {
		dir := rotation[:1]
		val, err := strconv.Atoi(rotation[1:])
		if err != nil {
			log.Fatalf("value cannot be formatted: %v", err)
		}
		val = val % 100
		switch dir {
		case "L":
			point -= val
			if point < 0 {
				point += 100
			}
		case "R":
			point += val
			point = point % 100
		}
		if point == 0 {
			counter++
		}
	}
	return counter
}

func problem2(filename string) int {
	rotations, err := readInput(filename)
	if err != nil {
		log.Fatalf("err from reading input file: %v", err)
	}
	var (
		pointed = 50
		counter = 0
	)
	for _, rotation := range rotations {
		dir := rotation[:1]
		number, err := strconv.Atoi(rotation[1:])
		if err != nil {
			log.Fatalf("value cannot be formatted: %v", err)
		}
		switch dir {
		case "L":
			counter += number / 100
			number = number % 100
			if number < pointed {
				pointed -= number
				continue
			}
			if pointed == 0 && number == 0 {
				continue
			}
			if number == pointed {
				pointed = 0
				counter++
				continue
			}
			if pointed != 0 {
				counter++
			}
			pointed = pointed - number + 100
		case "R":
			counter += number / 100
			number = number % 100
			if pointed+number < 100 {
				pointed += number
				continue
			}
			if pointed+number == 100 {
				pointed = 0
				counter++
				continue
			}
			counter++
			pointed = pointed + number - 100
		}
	}
	return counter
}

func readInput(filename string) ([]string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	var rotations []string

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		rotations = append(rotations, line)
	}
	return rotations, nil
}
