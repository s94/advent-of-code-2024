package day05

import (
	"advent-of-code/src/utils"
	"strconv"
	"strings"
)

type puzzleInput struct {
	pageOrderingRules []rules
	updates           [][]int
	validUpdates      [][]int
}

type rules struct {
	a int
	b int
}

var getPuzzleInput = convertPuzzleInput()

func PrintQueuePartOne() int {
	var retVal int

	processPuzzleInput()
	for _, update := range getPuzzleInput.validUpdates {
		retVal += update[int(len(update)/2)]
	}

	return retVal
}

func processPuzzleInput() {
	for _, update := range getPuzzleInput.updates {
		updateValid := true
		aFound := false
		bFound := false
		for i, u := range update {
			if i == 0 {
				for _, r := range getPuzzleInput.pageOrderingRules {
					if r.b == u {
						for _, p := range update {
							if r.a == p {
								updateValid = false
							}
						}
					}
				}
			} else {
				for _, r := range getPuzzleInput.pageOrderingRules {
					if r.b == u {
						for _, a := range update {
							if a == r.a {
								aFound = true
							}
						}
					} else if r.a == u {
						for j, b := range update {
							if j < i {
								if b == r.b {
									updateValid = false
								}
							}
							if b == r.b {
								bFound = true
							}
						}
					}
				}
			}
		}
		if updateValid && aFound && bFound {
			getPuzzleInput.validUpdates = append(getPuzzleInput.validUpdates, update)
		}
	}
}

func convertPuzzleInput() puzzleInput {
	file := utils.GetPuzzleInput(5)
	var rulesy []rules
	var updates [][]int
	withinUpdateSection := false

	getInt := func(s string) int {
		x, _ := strconv.Atoi(s)
		return x
	}

	for _, line := range file {
		if line == "" {
			withinUpdateSection = true
			continue
		}
		if withinUpdateSection {
			// process updates
			s := strings.Split(line, ",")
			var update []int
			for _, u := range s {
				update = append(update, getInt(u))
			}
			updates = append(updates, update)
		} else {
			// process rules
			s := strings.Split(line, "|")
			rulesy = append(rulesy, rules{
				a: getInt(s[0]),
				b: getInt(s[1]),
			})
		}
	}

	return puzzleInput{
		pageOrderingRules: rulesy,
		updates:           updates,
	}
}
