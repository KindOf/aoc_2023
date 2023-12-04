package day3

import (
	"fmt"
	"strings"
)

func detectAdjacement(schematic []string, row string, i, j int) (string, bool) {
	adjacents := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, a := range adjacents {
		ai := i + a[0]
		aj := j + a[1]

		if ai >= 0 && aj >= 0 && ai <= len(schematic)-1 && aj <= len(row)-1 {
			isStar := schematic[ai][aj] == '*'

			if isStar {
				return fmt.Sprintf("%d:%d", ai, aj), true
			}
		}
	}
	return "", false
}

func CollectParts() map[string][]int {
	dir := make(map[string][]int)
	schematic := strings.Split(strings.TrimSpace(input), "\n")

	for i, row := range schematic {
		curr := strings.Builder{}
		adj := ""

		for j, char := range row {
			_, isInt := toInt(string(char))

			if isInt {
				curr.WriteRune(char)

				if key, has := detectAdjacement(schematic, row, i, j); has && adj == "" {
					adj = key
                    if len(dir[adj]) == 0 {
					    dir[adj] = make([]int, 0)
                    }
				}
			}

			if isNextInt(row, j) {
				continue
			}

			if curr.Len() > 0 {
				n, ok := toInt(curr.String())

				if ok && adj != "" {
					dir[adj] = append(dir[adj], n)
				}

				curr.Reset()
			}

			adj = ""
		}
	}

	return dir
}

func PrintSumV2() {
	partMap := CollectParts()
	sum := 0

	for _, v := range partMap {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}

	fmt.Println(sum)
}
