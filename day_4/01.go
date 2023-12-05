package day4

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/KindOf/aoc_2023/utils"
)

//go:embed input
var input string

func splitCards(inputCards string) []string {
	return strings.Split(strings.TrimSpace(inputCards), "\n")
}

func getWinningNumbers(card string) (nums []int) {
	curr := strings.Builder{}
	state := "start"

	for i, char := range card {
		if char == '|' {
			break
		}
		if char == ' ' {
			continue
		}
		if state == "start" && char != ':' {
			continue
		}
		if state == "start" && char == ':' {
			state = "nums"
			continue
		}

		if _, ok := utils.ToInt(string(char)); ok && state == "nums" {
			curr.WriteRune(char)
		}

		if !utils.IsNextInt(card, i) {
			n, ok := utils.ToInt(curr.String())
			curr.Reset()
			if ok {
				nums = append(nums, n)
			}
		}
	}

	return
}

func getPick(card string) (nums []int) {
	curr := strings.Builder{}
	state := "start"

	for i, char := range card {
		if state == "start" && char != '|' {
			continue
		}

		if state == "start" && char == '|' {
			state = "nums"

			continue
		}

		if char == ' ' {
			continue
		}

		if _, ok := utils.ToInt(string(char)); ok && state == "nums" {
			curr.WriteRune(char)
		}

		if !utils.IsNextInt(card, i) {
			n, ok := utils.ToInt(curr.String())

			curr.Reset()

			if ok {
				nums = append(nums, n)
			}
		}

	}

	return
}

func evaluateResults(matchLength int) int  {
    if matchLength == 0 {
        return 0
    }
    res := 1

    for i := 0; i < matchLength -1; i++ {
        res *= 2
    }

    return res
}

func EvaluateCard(card string) int {
	winning := getWinningNumbers(card)
	pick := getPick(card)
    match := utils.IntersectionInts(winning, pick)

	return evaluateResults(len(match))
}

func CalculatePoints(inputCards string) (sum int) {
	for _, card := range splitCards(inputCards) {
		sum += EvaluateCard(card)
	}
	return
}

func PrintPoints() {
    fmt.Println("Result:", CalculatePoints(input))
}
