package day02

import (
	"advent-of-code/src/utils"
	"strconv"
	"strings"
)

var puzzleInput [][]int = convertPuzzleInput()

func RedNosedReportsPartOne() int {
	var retVal int

	for _, line := range puzzleInput {
		isIncreasing := line[0] < line[len(line)-1]
		if isSafe(line, isIncreasing, false) {
			retVal++
		}
	}

	return retVal
}

func RedNosedReportsPartTwo() int {
	var retVal int

	for _, line := range puzzleInput {
		isIncreasing := line[0] < line[len(line)-1]
		if isSafe(line, isIncreasing, true) {
			retVal++
		}
	}

	return retVal
}

func isSafe(levels []int, isIncreasing bool, tolerateSingleBadLevel bool) bool {
	for i, l := range levels {
		if i == 0 {
			continue
		}

		var x int
		if isIncreasing {
			x = l - levels[i-1]
		} else {
			x = levels[i-1] - l
		}

		if x < 1 || x > 3 {
			if !tolerateSingleBadLevel {
				return false
			}

			newLevelsVariantA := append([]int{}, levels[:i-1]...)
			newLevelsVariantA = append(newLevelsVariantA, levels[i:]...)

			newLevelsVariantB := append([]int{}, levels[:i]...)
			newLevelsVariantB = append(newLevelsVariantB, levels[i+1:]...)

			return isSafe(newLevelsVariantA, isIncreasing, false) || isSafe(newLevelsVariantB, isIncreasing, false)
		}
	}

	return true
}

func convertPuzzleInput() [][]int {
	var puzzleInput [][]int

	for _, line := range utils.GetPuzzleInput(2) {
		var splitLine []string = strings.Split(line, " ")
		var levels []int

		for _, l := range splitLine {
			x, _ := strconv.Atoi(l)
			levels = append(levels, x)
		}

		puzzleInput = append(puzzleInput, levels)
	}

	return puzzleInput
}
