package day3

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func toInt(a string) (int, bool) {
	n, err := strconv.Atoi(a)
	if err != nil {
		return 0, false
	}
	return n, true
}

func hasAdjacement(schematic []string, row string, i, j int) bool {
	adjacents := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, a := range adjacents {
		ai := i + a[0]
		aj := j + a[1]

		if ai >= 0 && aj >= 0 && ai <= len(schematic)-1 && aj <= len(row)-1 {
			_, isInt := toInt(string(schematic[ai][aj]))
			isDot := schematic[ai][aj] == '.'

			if !isInt && !isDot {
				return true
			}
		}
	}
	return false
}

func isNexInt(s string, index int) bool {
	if index == len(s)-1 {
		return false
	}
	_, isInt := toInt(string(s[index+1]))
	return isInt
}

func getSum(input string) (numSum int) {
	schematic := strings.Split(strings.TrimSpace(input), "\n")

	for i, row := range schematic {
		curr := strings.Builder{}
		hasAdj := false

		for j, c := range row {
			_, isInt := toInt(string(c))

			if isInt {
				curr.WriteRune(c)

				if !hasAdj {
					hasAdj = hasAdjacement(schematic, row, i, j)
				}

			}

            if isNexInt(row, j) {
                continue
            }

			if curr.Len() > 0 {
				n, ok := toInt(curr.String())
				fmt.Println(curr.String(), hasAdj)

				if ok && hasAdj {
					numSum += n
				}

			    curr.Reset()
			}

			hasAdj = false
		}
	}

	return
}

func PrintSum() {
	fmt.Println(getSum(input))
}
