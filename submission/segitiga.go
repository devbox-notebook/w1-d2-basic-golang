package submission

func LuasSegitiga(alas, tinggi float64) float64 {
	// not negative value
	if alas <= 0 || tinggi <= 0 {
		return 0
	}
	return (alas * tinggi) / 2
}
