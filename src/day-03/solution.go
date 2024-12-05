package day03

import (
	"advent-of-code/src/utils"
	"regexp"
	"strconv"
	"strings"
)

var puzzleInput []string = utils.GetPuzzleInput(3)

func MullItOverPartOne() int {
	var retVal int
	regex := regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)`)

	for _, line := range puzzleInput {
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			retVal += processMatch(match)
		}
	}

	return retVal
}

func MullItOverPartTwo() int {
	var retVal int
	regex := regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)`)
	joinedStrings := strings.Join(puzzleInput, "")
	doArray := strings.Split(joinedStrings, "do()")

	for _, d := range doArray {
		doString := strings.Split(d, "don't()")[0]

		matches := regex.FindAllStringSubmatch(doString, -1)
		for _, match := range matches {
			retVal += processMatch([]string{match[1]})
		}
	}

	return retVal
}

func processMatch(match []string) int {
	var retVal int
	for _, m := range match {
		s := strings.Split(m, ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])
		retVal += (x * y)
	}

	return retVal
}
