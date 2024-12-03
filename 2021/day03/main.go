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
	power := powerConsumption(lines)
	fmt.Println(power)

	val := powerConsumption2(lines)
	fmt.Println(val)

}

func powerConsumption(lines []string) int64 {
	var (
		gamma   []byte
		epsilon []byte
	)
	var bits = lines[0]

	for i := 0; i < len(bits); i++ {
		var (
			count0Bits int
			count1Bits int
		)
		for _, line := range lines {

			if line[i] == '0' {
				count0Bits++
			} else {
				count1Bits++
			}

		}
		if count0Bits > count1Bits {
			gamma = append(gamma, '0')
			epsilon = append(epsilon, '1')
		} else {
			gamma = append(gamma, '1')
			epsilon = append(epsilon, '0')
		}
	}
	gammaDec, _ := strconv.ParseInt(string(gamma), 2, 64)
	epsilonDec, _ := strconv.ParseInt(string(epsilon), 2, 64)

	return gammaDec * epsilonDec
}

func powerConsumption2(lines []string) int64 {
	var (
		bits         = lines[0]
		oxygenSource = lines
		co2Source    = lines
	)
	for i := 0; i < len(bits); i++ {
		zeroBitCount, oneBitCount := mostCommonValue(oxygenSource, i)
		if zeroBitCount > oneBitCount {
			oxygenSource = subList(oxygenSource, i, '0')
		} else {
			oxygenSource = subList(oxygenSource, i, '1')
		}
		if len(oxygenSource) == 1 {
			break
		}
	}

	for i := 0; i < len(bits); i++ {
		zeroBitCount, oneBitCount := mostCommonValue(co2Source, i)
		if zeroBitCount < oneBitCount || zeroBitCount == oneBitCount {
			co2Source = subList(co2Source, i, '0')
		} else {
			co2Source = subList(co2Source, i, '1')
		}
		if len(co2Source) == 1 {
			break
		}
	}

	oxygen := oxygenSource[0]
	co2 := co2Source[0]

	oxygenVal, _ := strconv.ParseInt(oxygen, 2, 64)
	co2Val, _ := strconv.ParseInt(co2, 2, 64)

	return oxygenVal * co2Val
}

func mostCommonValue(lines []string, pos int) (zeroBitCount int, oneBitCount int) {
	for _, line := range lines {
		if line[pos] == '0' {
			zeroBitCount++
		} else {
			oneBitCount++
		}
	}
	return
}

func subList(lines []string, pos int, targetVal byte) []string {
	var list []string
	for _, line := range lines {
		if line[pos] == targetVal {
			list = append(list, line)
		}
	}
	return list
}
