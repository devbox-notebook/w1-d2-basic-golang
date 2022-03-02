package submission

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	got, _ := Hello("Kobar")
	expect := "Hello Kobar, Saya Golang Bahasa yang Menyenangkan!"

	assert.Equal(t, expect, got)

	got, _ = Hello("Adi")
	expect = "Hello Adi, Saya Golang Bahasa yang Menyenangkan!"

	assert.Equal(t, expect, got)

	_, err := Hello("-adw")
	assert.EqualError(t, err, "nama hanya boleh berisi huruf")

	_, err = Hello("")
	assert.EqualError(t, err, "nama tidak boleh kosong")

	got, _ = Hello("Adi Gunawan")
	expect = "Adi Gunawan"

	assert.NotEqual(t, expect, got)

}
