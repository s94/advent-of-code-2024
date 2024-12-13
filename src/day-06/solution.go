package day06

import (
	"advent-of-code/src/utils"
	"sync"
)

type coords struct {
	x int
	y int
}

type direction int

const (
	north    direction = iota // 0
	east                      // 1
	south                     // 2
	west                      // 3
	err                       // 4
	obstacle string    = "#"
	visited  string    = "X"
)

func GuardGallivantPartOne() int {
	var retVal int

	puzzleInput := simulateScenario(coords{}, make(chan int, 1), nil)
	for _, row := range puzzleInput {
		for _, cell := range row {
			if cell == visited {
				retVal++
			}
		}
	}

	return retVal
}

func GuardGallivantPartTwo() int {
	var retVal int
	ch := make(chan int)
	var wg sync.WaitGroup

	for i, row := range getPuzzleInput() {
		for j, cell := range row {
			if cell == obstacle || cell == "^" {
				continue
			} else {
				wg.Add(1)
				go simulateScenario(coords{
					x: j,
					y: i,
				}, ch, &wg)
			}
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		if result == 6000 {
			retVal++
		}
	}

	return retVal
}

func simulateScenario(obstacleCoords coords, ch chan<- int, wg *sync.WaitGroup) [][]string {
	puzzleInput := getPuzzleInput()
	puzzleInput[obstacleCoords.y][obstacleCoords.x] = obstacle

	guardDirection := north
	guardCoords := coords{
		x: 89,
		y: 84,
	}
	nextCoords := getNextCoords(guardDirection, guardCoords)

	counter := 0
	for {
		if counter >= 6000 {
			break
		}
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
		counter++
	}
	ch <- counter
	if wg != nil {
		wg.Done()
	}

	return puzzleInput
}

func turnGuard(currentDirection direction) direction {
	switch currentDirection {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	case west:
		return north
	default:
		return err
	}
}

func getNextCoords(d direction, currentCoords coords) coords {
	switch d {
	case north:
		return coords{
			x: currentCoords.x,
			y: currentCoords.y - 1,
		}
	case east:
		return coords{
			x: currentCoords.x + 1,
			y: currentCoords.y,
		}
	case south:
		return coords{
			x: currentCoords.x,
			y: currentCoords.y + 1,
		}
	case west:
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
