package day3

import (
	"strings"
	"testing"
)

func TestToInt(t *testing.T) {
	s := "."
	got, ok := toInt(s)

	if ok {
		t.Fatalf("Got = %v", got)
	}
}

func TestGetSum(t *testing.T) {
	schematic :=
		`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	got := getSum(schematic)

	if got != 4361 {
		t.Fatalf("Got = %d, want 4361\n", got)

	}

}

func TestDetectAdjacement(t *testing.T) {
	schematic := strings.Split(strings.TrimSpace(
		`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592....*
.......755
...$.*....
.664.598..`), "\n")

	tests := []struct {
		i, j    int
		want    string
		wantHas bool
	}{
		{0, 1, "", false},
		{0, 2, "1:3", true},
		{2, 2, "1:3", true},
		{2, 3, "1:3", true},
		{2, 3, "1:3", true},
		{0, 5, "", false},
        {7, 9, "6:9", true},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t1 *testing.T) {
			got, has := detectAdjacement(schematic, schematic[tt.i], tt.i, tt.j)
			if got != tt.want && has != tt.wantHas {
				t1.Fatalf(
					"Got = %s; want = %s; Has = %v; wantHas = %v",
					got, tt.want, has, tt.wantHas,
				)
			}
		})
	}
}

func TestPrintSum(t *testing.T) {
	PrintSum()

	t.Fatal("Not Implemented")
}

func TestPrintSumV2(t *testing.T) {
	PrintSumV2()

	t.Fatal("Not Implemented")
}
