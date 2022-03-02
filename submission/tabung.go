package submission

import "math"

func LuasPermukaanTabung(r, t float64) float64 {
	return 2 * math.Pi * r * (r + t)
}
