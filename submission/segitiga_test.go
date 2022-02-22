package submission

import "testing"

func TestLuasSegitiga(t *testing.T) {
	type test struct {
		alas, tinggi float64
		hasil        float64
	}

	tests := []test{
		{5, 2, 5},
		{5, 3, 7.5},
		{20, 25, 250},
	}

	for _, tt := range tests {
		actual := LuasSegitiga(tt.alas, tt.tinggi)
		if actual != tt.hasil {
			t.Errorf("LuasSegitiga(%f, %f) = %f, want %f", tt.alas, tt.tinggi, actual, tt.hasil)
		}
	}
}
