package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func toInt(s string) (int, bool) {
	n, err := strconv.Atoi(s)
	return n, err == nil
}

func findIntInString(s string) int {
	wordMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine)`)
	match := re.FindString(s)

	if match != "" {
		return wordMap[match]
	}

	return 0
}

func calculateForLine(line string) int {
	lw, rw := "", ""
    l, r := 0, 0

	for i := 0; i < len(line); i++ {
        w := findIntInString(lw)
        if w != 0 {
            l = w
            break
        }
		if n, ok := toInt(string(line[i])); ok {
            l = n
			break
		}
        lw += string(line[i])
	}

	for j := len(line)-1; j >= 0; j-- {
		w := findIntInString(rw)
        if w != 0 {
            r = w
            break
        }
		if n, ok := toInt(string(line[j])); ok {
            r = n
			break
		}
        rw = string(line[j]) + rw
	}

    if res, ok := toInt(fmt.Sprintf("%d%d", l, r)); ok {
        return res
    }

	return 0
}

func readFile() {
	f, err := os.Open("1-1-input")
	if err != nil {
		panic(err)
	}

	info, err := f.Stat()
	if err != nil {
		panic(err)
	}
	size := info.Size()
	offset := 0

	var buffer int
	r := bufio.NewReader(f)

	for offset < int(size) {
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		offset += len(line)
		buffer += calculateForLine(line)
	}

	fmt.Println(buffer)
}
