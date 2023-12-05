package utils

import "strconv"

func ToInt(a string) (int, bool) {
    n, err := strconv.Atoi(a)
    return n, err != nil
}
