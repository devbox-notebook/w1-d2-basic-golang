package submission

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Kobar")
	expect := "Hello Kobar, Saya Golang Bahasa yang Menyenangkan!"

	if got != expect {
		t.Errorf("Hello(\"Kobar\") = %q, want %q", got, expect)
	}
}
