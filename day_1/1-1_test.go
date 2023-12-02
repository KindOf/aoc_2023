package day1

import "testing"

func TestCalculateForLine(t *testing.T) {
	tests := []struct {
		line string
		want int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
		{"fivepqxlpninevh2xxsnsgg63pbvdnqptmg", 53},
		{"649twomktwonebx", 61},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t1 *testing.T) {
			got := calculateForLine(tt.line)

			if got != tt.want {
				t1.Fatalf("Got = %d, want = %d", got, tt.want)
			}
		})

	}
}

func TestReadFile(t *testing.T) {
	readFile()
	t.Fatal("Not implemented")
}
