package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func GetPuzzleInput(day int) []string {
	fileSystem := os.DirFS("../../src/day-" + fmt.Sprintf("%02d", day))
	file, err := fileSystem.Open("puzzle-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}
