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
	increased := sonarSweep(lines)
	fmt.Println(increased)

	increased = sonarSweep2(lines)
	fmt.Println(increased)
}

func sonarSweep(vals []string) int {
	var (
		increased int
		nums      = convertInts(vals)
	)

	for i := 1; i < len(nums); i++ {
		current := nums[i]
		prev := nums[i-1]
		if current > prev {
			increased++
		}
	}
	return increased
}

func sonarSweep2(vals []string) int {
	var (
		increased       int
		nums            = convertInts(vals)
		sumThreePrev    = nums[0] + nums[1] + nums[2]
		sumThreeCurrent int
		left            = 1
	)
	for right := 3; right < len(nums); right++ {
		sumThreeCurrent += nums[left] + nums[right-1] + nums[right]
		if sumThreeCurrent > sumThreePrev {
			increased++
		}
		left++
		sumThreePrev = sumThreeCurrent
		sumThreeCurrent = 0
	}
	return increased
}

func convertInts(vals []string) []int {
	var nums = make([]int, len(vals))
	for i, val := range vals {
		nums[i], _ = strconv.Atoi(val)
	}
	return nums
}
