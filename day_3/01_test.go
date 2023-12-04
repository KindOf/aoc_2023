package day3

import (
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

func TestPrintSum(t *testing.T) {
   PrintSum()

   t.Fatal("Not Implemented")
}
