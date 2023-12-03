package day2

import (
	"testing"
)

func TestIsPossible(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
		{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
		{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false},
		{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false},
		{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t1 *testing.T) {
			got := isPossible(tt.line)
			if got != tt.want {
				t1.Fatalf("Got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestGetGameId(t *testing.T) {
	tests := []struct {
		line string
		want int
	}{
		{"Game 123:some text", 123},
		{"Game 124", 0},
		{"Game", 0},
		{"Not a Game 123", 0},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t1 *testing.T) {
			got := getGameId(tt.line)
			if got != tt.want {
				t.Fatalf("got = %v, want = %d", got, tt.want)
			}
		})
	}
}

func TestGetGameSum(t *testing.T) {
	GetGameSum()
	t.Fatal("Not imlemented\n")
}
