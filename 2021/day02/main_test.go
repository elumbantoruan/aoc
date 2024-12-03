package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productHorizontalAndDepth(t *testing.T) {
	lines := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	product := productHorizontalAndDepth(lines)
	assert.Equal(t, 150, product)
}

func Test_productHorizontalAndDepth2(t *testing.T) {
	lines := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	product := productHorizontalAndDepth2(lines)
	assert.Equal(t, 900, product)
}
