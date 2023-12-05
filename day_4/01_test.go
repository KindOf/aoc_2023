package day4

import (
	"fmt"
	"reflect"
	"testing"
)

const inputCardsString = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func TestCalculatePoints(t *testing.T) {
	got := CalculatePoints(inputCardsString)
	want := 13

	if got != want {
		t.Fatalf("Got = %d; want = %d", got, want)
	}
}

func TestEvaluateCard(t *testing.T) {
	tests := []struct {
		card string
		want int
	}{
		{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", 8},
		{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", 2},
		{"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", 2},
		{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", 1},
		{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", 0},
		{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", 0},
	}

	for _, tt := range tests {
		t.Run(tt.card, func(t1 *testing.T) {
			got := EvaluateCard(tt.card)

			if got != tt.want {
				t1.Fatalf("Got = %d; want = %d", got, tt.want)
			}
		})
	}
}

func TestGetWinningNumbers(t *testing.T) {
	tests := []struct {
		card string
		want []int
	}{
		{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			[]int{41, 48, 83, 86, 17},
		},
		{
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			[]int{13, 32, 20, 16, 61},
		},
		{
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			[]int{1, 21, 53, 59, 44},
		},
		{
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			[]int{41, 92, 73, 84, 69},
		},
		{
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			[]int{87, 83, 26, 28, 32},
		},
		{
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			[]int{31, 18, 13, 56, 72},
		},
	}

	for _, tt := range tests {
		t.Run(tt.card, func(t1 *testing.T) {
			got := getWinningNumbers(tt.card)

			if !reflect.DeepEqual(got, tt.want) {
				t1.Fatalf("Got = %d; want = %d", got, tt.want)
			}
		})
	}
}

func TestGetPick(t *testing.T) {
	tests := []struct {
		card string
		want []int
	}{
		{
			"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			[]int{83, 86, 6, 31, 17, 9, 48, 53},
		},
		{
			"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			[]int{61, 30, 68, 82, 17, 32, 24, 19},
		},
		{
			"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			[]int{69, 82, 63, 72, 16, 21, 14, 1},
		},
		{
			"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			[]int{59, 84, 76, 51, 58, 5, 54, 83},
		},
		{
			"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			[]int{88, 30, 70, 12, 93, 22, 82, 36},
		},
		{
			"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			[]int{74, 77, 10, 23, 35, 67, 36, 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.card, func(t1 *testing.T) {
			got := getPick(tt.card)

			if !reflect.DeepEqual(got, tt.want) {
				t1.Fatalf("Got = %d; want = %d", got, tt.want)
			}
		})
	}
}

func TestEvaluateResults(t *testing.T) {
    tests := []struct{
        len int
        want int
    }{
        {0, 0}, {1, 1}, {2, 2}, {3, 4}, {4, 8}, {5, 16},
    }

    for _, tt := range tests {
        t.Run(fmt.Sprintf("Match length: %d", tt.len), func(t1 *testing.T) {
            got := evaluateResults(tt.len)

            if got != tt.want {
                t1.Fatalf("Got = %d; want = %d", got, tt.want)
            }
        })
    }
}

func TestCalculateCards(t *testing.T) {
    cardInput := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11` 

    got := calculateCards(cardInput)
    want := 30

    if got != want {
        t.Fatalf("Got = %d; want = %d", got, want)
    }
}

func TestPrintPoints(t *testing.T) {
    PrintPoints()
    t.Fatal("Not implemented\n")
}

func TestPrintCardNum(t *testing.T) {
    PrintCardNum()
    t.Fatal("Not implemented\n")
}
