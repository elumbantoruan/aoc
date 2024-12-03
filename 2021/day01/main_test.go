package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sonarSweep(t *testing.T) {
	vals := []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}
	increased := sonarSweep(vals)
	assert.Equal(t, 7, increased)

}

func Test_sonarSweep2(t *testing.T) {
	vals := []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}
	increased := sonarSweep2(vals)
	assert.Equal(t, 5, increased)
}
