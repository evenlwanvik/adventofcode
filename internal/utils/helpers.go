package utils

import (
	"os"
)

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func PrintLines(lines []string) {
	for _, line := range lines {
		println(line)
	}
}
