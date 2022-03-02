package submission

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLuasSegitiga(t *testing.T) {
	type test struct {
		alas, tinggi float64
		hasil        float64
	}

	tests := []test{
		{5, 2, 5},
		{5, 3, 7.5},
		{20, 25, 250},
		{0, 0, 0},
		{-5, -2, 0},
	}

	for _, tt := range tests {
		actual := LuasSegitiga(tt.alas, tt.tinggi)
		assert.Equal(t, tt.hasil, actual, "LuasSegitiga(%f, %f) = %f, want %f", tt.alas, tt.tinggi, actual, tt.hasil)
	}
}
