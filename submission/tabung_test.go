package submission

import "testing"

func TestLuasPermukaanTabung(t *testing.T) {
	type test struct {
		t float64
		r float64
		l float64
	}
	tests := []test{
		{20, 4, 602.88},
	}

	for _, tt := range tests {
		actual := LuasPermukaanTabung(tt.r, tt.t)
		if actual != tt.l {
			t.Errorf("LuasPermukaanTabung(%f, %f) = %f, want %f", tt.r, tt.t, actual, tt.l)
		}
	}
}
