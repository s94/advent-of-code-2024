package day01_test

import (
	day01 "advent-of-code/src/day-01"
	"testing"
)

func TestHistorianHysteriaPartOne(t *testing.T) {
	getAnswer := day01.HistorianHysteriaPartOne()

	if getAnswer != 1879048 {
		t.Fatal("Incorrect")
	}
}

func TestHistorianHysteriaPartTwo(t *testing.T) {
	getAnswer := day01.HistorianHysteriaPartTwo()

	if getAnswer != 21024792 {
		t.Fatal("Incorrect")
	}
}
