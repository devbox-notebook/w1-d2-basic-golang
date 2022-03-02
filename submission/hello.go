package submission

import (
	"fmt"
	"regexp"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("nama tidak boleh kosong")
	}

	regex := regexp.MustCompile(`^[a-zA-Z]+$`)
	if !regex.MatchString(name) {
		return "", fmt.Errorf("nama hanya boleh berisi huruf")
	}

	greeting := fmt.Sprintf("Hello %s, Saya Golang Bahasa yang Menyenangkan!", name)

	return greeting, nil
}
