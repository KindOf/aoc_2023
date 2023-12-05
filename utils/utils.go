package utils

import (
	"sort"
	"strconv"
)

func ToInt(a string) (int, bool) {
	n, err := strconv.Atoi(a)
	return n, err == nil
}

func IsNextInt(s string, index int) bool {
	if index == len(s)-1 {
		return false
	}
	_, isInt := ToInt(string(s[index+1]))
	return isInt
}

func IntersectionInts(a []int, b []int) []int {
	set := make([]int, 0)
	hash := make(map[int]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}

func Intersection[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		idx := sort.Search(len(b), func(i int) bool {
			return b[i] == v
		})
		if idx < len(b) && b[idx] == v {
			set = append(set, v)
		}
	}

	return set
}
