package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)
	lines := strings.Split(text, "\n")
	product := productHorizontalAndDepth(lines)
	fmt.Println(product)

	product = productHorizontalAndDepth2(lines)
	fmt.Println(product)

}

func productHorizontalAndDepth(lines []string) int {
	var (
		horizontal int
		depth      int
	)

	for _, line := range lines {
		if strings.HasPrefix(line, "forward") {
			parts := strings.Split(line, " ")
			num, _ := strconv.Atoi(parts[1])
			horizontal += num
		} else if strings.HasPrefix(line, "down") {
			parts := strings.Split(line, " ")
			num, _ := strconv.Atoi(parts[1])
			depth += num
		} else if strings.HasPrefix(line, "up") {
			parts := strings.Split(line, " ")
			num, _ := strconv.Atoi(parts[1])
			depth -= num
		}
	}
	return horizontal * depth
}

func productHorizontalAndDepth2(lines []string) int {
	var (
		horizontal int
		depth      int
		aim        int
	)

	for _, line := range lines {
		if strings.HasPrefix(line, "forward") {
			parts := strings.Split(line, " ")
			num, _ := strconv.Atoi(parts[1])
			horizontal += num
			depth += num * aim
		} else if strings.HasPrefix(line, "down") {
			parts := strings.Split(line, " ")
			num, _ := strconv.Atoi(parts[1])
			aim += num
		} else if strings.HasPrefix(line, "up") {
			parts := strings.Split(line, " ")
			num, _ := strconv.Atoi(parts[1])
			aim -= num
		}
	}
	return horizontal * depth
}
