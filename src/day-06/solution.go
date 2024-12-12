package day06

import (
	"advent-of-code/src/utils"
)

type coords struct {
	x int
	y int
}

type direction int

const (
	North direction = iota // 0
	East                   // 1
	South                  // 2
	West                   // 3
	Err                    // -1
)

var puzzleInput = getPuzzleInput()

func GuardGallivantPartOne() int {
	var retVal int

	const obstacle string = "#"
	const visited string = "X"
	guardDirection := North
	guardCoords := coords{
		x: 89,
		y: 84,
	}
	nextCoords := getNextCoords(guardDirection, guardCoords)

	for {
		if (nextCoords.y < 0 || nextCoords.y > len(puzzleInput)-1) || (nextCoords.x < 0 || nextCoords.x > len(puzzleInput[0])-1) {
			puzzleInput[guardCoords.y][guardCoords.x] = visited
			break
		}

		if puzzleInput[nextCoords.y][nextCoords.x] == obstacle {
			guardDirection = turnGuard(guardDirection)
			nextCoords = getNextCoords(guardDirection, guardCoords)
		} else {
			previousCoords := coords{
				x: guardCoords.x,
				y: guardCoords.y,
			}
			guardCoords = getNextCoords(guardDirection, guardCoords)
			nextCoords = getNextCoords(guardDirection, guardCoords)
			puzzleInput[previousCoords.y][previousCoords.x] = visited
		}
	}

	for _, row := range puzzleInput {
		for _, cell := range row {
			if cell == visited {
				retVal++
			}
		}
	}

	return retVal
}

func turnGuard(currentDirection direction) direction {
	switch currentDirection {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	default:
		return Err
	}
}

func getNextCoords(d direction, currentCoords coords) coords {
	switch d {
	case North:
		return coords{
			x: currentCoords.x,
			y: currentCoords.y - 1,
		}
	case East:
		return coords{
			x: currentCoords.x + 1,
			y: currentCoords.y,
		}
	case South:
		return coords{
			x: currentCoords.x,
			y: currentCoords.y + 1,
		}
	case West:
		return coords{
			x: currentCoords.x - 1,
			y: currentCoords.y,
		}
	default:
		return coords{
			x: -1,
			y: -1,
		}
	}
}

func getPuzzleInput() [][]string {
	retVal := [][]string{}

	file := utils.GetPuzzleInput(6)
	for _, line := range file {
		r := []string{}
		for _, s := range line {
			r = append(r, string(s))
		}
		retVal = append(retVal, r)
	}

	return retVal
}
