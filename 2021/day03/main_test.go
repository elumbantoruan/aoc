package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_powerConsumption(t *testing.T) {
	lines := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	power := powerConsumption(lines)
	assert.Equal(t, int64(198), power)
}

func Test_powerConsumption2(t *testing.T) {
	lines := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	power := powerConsumption2(lines)
	assert.Equal(t, int64(230), power)
}
