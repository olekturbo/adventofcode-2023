package utils

import (
	"os"
	"strings"
)

func ReadInput(path string) []string {
	f, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(f), "\n")
}
