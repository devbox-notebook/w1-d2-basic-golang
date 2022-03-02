package submission

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLuasPermukaanTabung(t *testing.T) {
	type test struct {
		t float64
		r float64
		l float64
	}
	var result1 float64 = 192
	var result2 float64 = 522
	var result3 float64 = 32
	var result4 float64 = 24
	var result5 float64 = 40
	tests := []test{
		{20, 4, result1 * math.Pi},
		{20, 9, result2 * math.Pi},
		{0, 4, result3 * math.Pi},
		{-1, 4, result4 * math.Pi},
		{-1, -4, result5 * math.Pi},
	}

	for _, tt := range tests {
		actual := LuasPermukaanTabung(tt.r, tt.t)
		assert.Equal(t, tt.l, actual, "LuasPermukaanTabung(%f, %f) = %f, want %f", tt.r, tt.t, actual, tt.l)
	}
}
