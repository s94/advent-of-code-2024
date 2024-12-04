package day01

import (
	"advent-of-code/src/utils"
	"sort"
	"strconv"
	"strings"
)

type input struct {
	left  []int
	right []int
}

var puzzleInput = orderInput()

func HistorianHysteriaPartOne() int {
	retVal := 0

	for i := 0; i < len(puzzleInput.left); i++ {
		if puzzleInput.left[i] > puzzleInput.right[i] {
			retVal += puzzleInput.left[i] - puzzleInput.right[i]
			continue
		}

		retVal += puzzleInput.right[i] - puzzleInput.left[i]
	}

	return retVal
}

func HistorianHysteriaPartTwo() int {
	retVal := 0

	for i := 0; i < len(puzzleInput.left); i++ {
		if !rightHasValue(puzzleInput.left[i]) {
			continue
		}

		count := 0
		for _, r := range puzzleInput.right {
			if r == puzzleInput.left[i] {
				count++
			}

			if r > puzzleInput.left[i] {
				break
			}
		}

		retVal += (puzzleInput.left[i] * count)
	}

	return retVal
}

func rightHasValue(value int) bool {
	for _, r := range puzzleInput.right {
		if r == value {
			return true
		}
	}

	return false
}

func orderInput() input {
	left := []int{}
	right := []int{}

	for _, line := range utils.GetPuzzleInput(1) {
		x := strings.Split(line, "   ")
		l, _ := strconv.Atoi(x[0])
		r, _ := strconv.Atoi(x[1])
		left = append(left, l)
		right = append(right, r)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	return input{
		left:  left,
		right: right,
	}
}
