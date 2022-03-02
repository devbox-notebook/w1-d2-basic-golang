package main

import (
	"fmt"

	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/w1-d2-basic-golang/submission"
)

func main() {

	// soal 1
	nama := "Kobar"

	message, _ := submission.Hello(nama)

	fmt.Println(message)

	// soal 2
	var alas, tinggi float64 = 20, 25

	var hasil = submission.LuasSegitiga(alas, tinggi)

	fmt.Printf("Luas Segitiga dengan alas %.2f cm2 dan tinggi %.2f cm adalah %.2f cm2\n", alas, tinggi, hasil)

	// soal 3
	var r, T float64 = 4, 20

	luas := submission.LuasPermukaanTabung(r, T)

	fmt.Printf("Luas permukaan Tabung dengan tinggi %.2f cm dan radius %.2f cm adalah %.2f cm2\n", r, T, luas)

}
