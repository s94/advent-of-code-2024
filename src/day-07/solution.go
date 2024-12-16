package day07

import (
	"advent-of-code/src/utils"
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

type equation struct {
	answer  int
	numbers []int
}

func BridgeRepairPartOne() int {
	var retVal int
	ch := make(chan int)
	var wg sync.WaitGroup

	puzzleInput := convertPuzzleInput()

	for _, equation := range puzzleInput {
		wg.Add(1)
		go evaluateEquation(equation, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		retVal += result
	}

	return retVal
}

func evaluateEquation(e equation, ch chan<- int, wg *sync.WaitGroup) {
	var operatorsArray []string
	n := len(e.numbers) - 1
	combinations := 1 << n
	for {
		if len(operatorsArray) == combinations {
			break
		}
		s := generateRandomString(n)
		found := false
		for _, operators := range operatorsArray {
			if operators == s {
				found = true
			}
		}
		if !found {
			operatorsArray = append(operatorsArray, s)
		}
	}
	for _, operators := range operatorsArray {
		sum := 0
		for j, x := range e.numbers {
			if j == 0 {
				continue
			}
			if sum == 0 {
				sum = e.numbers[j-1]
			}
			if operators[j-1] == '+' {
				sum = sum + x
			} else {
				sum = sum * x
			}
		}
		if sum == e.answer {
			ch <- sum
			wg.Done()
			return
		}
	}
	wg.Done()
}

func generateRandomString(n int) string {
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		if rand.Intn(2) == 0 {
			result[i] = '+'
		} else {
			result[i] = '*'
		}
	}

	return string(result)
}

func convertPuzzleInput() []equation {
	var retVal []equation

	for i, line := range utils.GetPuzzleInput(7) {
		if i != 0 {

		}
		a, _ := strconv.Atoi(strings.Split(line, ":")[0])
		n := strings.Split(strings.Split(line, ":")[1], " ")
		numbers := []int{}
		for j, x := range n {
			if j == 0 {
				continue
			}
			y, _ := strconv.Atoi(x)
			numbers = append(numbers, y)
		}
		retVal = append(retVal, equation{a, numbers})
	}

	return retVal
}
