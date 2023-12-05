package utils

import (
	"reflect"
	"testing"
)

func TestIntersectionInts(t *testing.T) {
	tests := []struct {
		name  string
		input [2][]int
		want  []int
	}{
		{"empty inputs", [2][]int{}, []int{}},
		{"one match", [2][]int{{1}, {1}}, []int{1}},
		{"one match 2", [2][]int{{1, 3}, {1, 2}}, []int{1}},
		{"one match 3", [2][]int{{5, 3}, {5, 2, 3}}, []int{5, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			got := IntersectionInts(tt.input[0], tt.input[1])
			if !reflect.DeepEqual(tt.want, got) {
				t1.Fatalf("Got = %v; want = %v", got, tt.want)
			}
		})
	}

}
