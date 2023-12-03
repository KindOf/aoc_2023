package day2

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

const (
	red, green, blue = 12, 13, 14
)

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err == nil {
		return n
	}
	return 0
}

func getGameId(line string) int {
	re := regexp.MustCompile(`^Game (\d*):`)
	s := re.FindStringSubmatch(line)
	if len(s) < 2 {
		return 0
	}
	return toInt(s[1])
}

func isPossibleRound(roll string, re *regexp.Regexp, max int) bool {
	match := re.FindAllStringSubmatch(roll, -1)
	for _, m := range match {
		if toInt(m[1]) > max {
			return false
		}
	}
	return true
}

func isPossible(line string) bool {
	rer := regexp.MustCompile(`(\d+) red`)
	reg := regexp.MustCompile(`(\d+) green`)
	reb := regexp.MustCompile(`(\d+) blue`)

	game := strings.Split(line, ": ")[1]
	rounds := strings.Split(game, "; ")

	for _, round := range rounds {
		if !isPossibleRound(round, rer, red) {
			return false
		}

		if !isPossibleRound(round, reg, green) {
			return false
		}

		if !isPossibleRound(round, reb, blue) {
			return false
		}
	}
	return true
}

func processGame(line string) (int, bool) {
	gameId := getGameId(line)
    possible := isPossible(line)

    return gameId, possible
}

func GetGameSum() {
    sum := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
        if n, ok := processGame(line); ok {
            sum += n
        }
	}

    fmt.Println(sum)
}
