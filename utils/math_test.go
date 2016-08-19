package utils

import (
	"testing"

	"github.com/TerraFactory/tilegenerator/utils"
	"github.com/stretchr/testify/assert"
)

func TestDistanceBetweenPoints(t *testing.T) {
	var distTest = []struct {
		x1, y1, x2, y2 float64
		expected       float64
	}{
		{x1: 10, y1: 10, x2: 10, y2: 0, expected: 10},
		{x1: -20, y1: 0, x2: 0, y2: 0, expected: 20},
		{x1: -5, y1: 15, x2: -2, y2: 11, expected: 5},
	}

	for _, item := range distTest {
		distance := utils.DistanceBetweenPoints(item.x1, item.y1, item.x2, item.y2)
		assert.Equal(t, item.expected, distance)
	}
}

func TestLengthPolyline(t *testing.T) {
	var lenghtTest = []struct {
		xs, ys   []float64
		expected float64
	}{
		{xs: []float64{0, 0}, ys: []float64{0, 0}, expected: 0.0},
		{xs: []float64{0}, ys: []float64{0}, expected: 0.0},
		{xs: []float64{-10, -10, 0, 0, 15, 15}, ys: []float64{0, -5, -5, 0, 0, -10}, expected: 45},
		{xs: []float64{0, 10, 20}, ys: []float64{0, 5, 5}, expected: 21.18033988749895},
	}

	for _, item := range lenghtTest {
		length := utils.LengthPolyline(item.xs, item.ys)
		assert.Equal(t, item.expected, length)
	}
}
