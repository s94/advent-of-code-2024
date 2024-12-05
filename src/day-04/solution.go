package day04

import (
	"advent-of-code/src/utils"
)

var puzzleInput [][]string = getPuzzleInput()
var directions [][]int = [][]int{
	{-1, -1}, // up left
	{-1, 0},  // up
	{-1, 1},  //up right
	{0, -1},  // left
	{0, 1},   // right
	{1, -1},  // down left
	{1, 0},   // down
	{1, 1},   // down right
}

func CeresSearchPartOne() int {
	var retVal int
	for i, line := range puzzleInput {
		for j, char := range line {
			if char == "X" {
				for _, d := range directions {
					retVal += checkDirection(i, j, d[0], d[1])
				}
			}
		}
	}

	return retVal
}

func CeresSearchPartTwo() int {
	var retVal int

	for i, line := range puzzleInput {
		for j, char := range line {
			if char == "A" {
				retVal += checkForX(i, j)
			}
		}
	}

	return retVal
}

func checkForX(x int, y int) int {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	uL := puzzleInput[x+directions[0][0]][y+directions[0][1]]
	uR := puzzleInput[x+directions[2][0]][y+directions[2][1]]
	dL := puzzleInput[x+directions[5][0]][y+directions[5][1]]
	dR := puzzleInput[x+directions[7][0]][y+directions[7][1]]

	if (uL == "M" && dR == "S") || (uL == "S" && dR == "M") {
		if (uR == "M" && dL == "S") || (uR == "S" && dL == "M") {
			return 1
		}
	}

	return 0
}

func checkDirection(x int, y int, xDirection int, yDirection int) int {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	found := func(letter string) bool {
		x += xDirection
		y += yDirection
		return puzzleInput[x][y] == letter
	}

	if found("M") {
		if found("A") {
			if found("S") {
				return 1
			}
		}
	}

	return 0
}

func getPuzzleInput() [][]string {
	var retVal [][]string
	file := utils.GetPuzzleInput(4)

	for _, line := range file {
		var row []string
		lineLength := len(line)

		for i := 0; i < lineLength; i++ {
			row = append(row, string(line[i]))
		}
		retVal = append(retVal, row)
	}

	return retVal
}
