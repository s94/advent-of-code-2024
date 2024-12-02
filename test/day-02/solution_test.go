package day02_test

import (
	day02 "advent-of-code/src/day-02"
	"testing"
)

func TestRedNosedReportsPartOne(t *testing.T) {
	getAnswer := day02.RedNosedReportsPartOne()

	if getAnswer != 591 {
		t.Fatal("Incorrect answer")
	}
}

func TestRedNosedReportsPartTwo(t *testing.T) {
	getAnswer := day02.RedNosedReportsPartTwo()

	if getAnswer != 621 {
		t.Fatal("Incorrect answer")
	}
}
