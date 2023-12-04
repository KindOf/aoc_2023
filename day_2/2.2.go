package day2

import (
	"fmt"
	"regexp"
	"strings"
)

func getMinForRounds(rounds string, re *regexp.Regexp) int {
    var min int = 0
	match := re.FindAllStringSubmatch(rounds, -1)

	for _, m := range match {
		n := toInt(m[1])

        if min == 0 {
            min = n
        }

		if n > min {
            min = n
		}
	}

	return min
}

func getGameProduct(game string) int {
	rer := regexp.MustCompile(`(\d+) red`)
	reg := regexp.MustCompile(`(\d+) green`)
	reb := regexp.MustCompile(`(\d+) blue`)

	rounds := strings.Split(game, ": ")[1]
    minr := getMinForRounds(rounds, rer)
    ming := getMinForRounds(rounds, reg)
    minb := getMinForRounds(rounds, reb)

	return minr * ming * minb
}

func GetMinSet() {
	sum := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		n := getGameProduct(line)
		sum += n
	}
    fmt.Println(sum)
}

