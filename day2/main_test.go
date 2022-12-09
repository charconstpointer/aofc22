package main

import "testing"

func TestRound(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      Res
		wantScore int
	}{
		{
			name:      "A Y",
			input:     "A Y",
			want:      Win,
			wantScore: 8,
		},
		{
			name:      "B X",
			input:     "B X",
			want:      Loss,
			wantScore: 1,
		},
		{
			name:      "C Z",
			input:     "C Z",
			want:      Draw,
			wantScore: 6,
		},
		{
			name:      "A X",
			input:     "A X",
			want:      Draw,
			wantScore: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			round := parse(tt.input)
			got := round.Result()
			if got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
			score := calcScore(round)
			if score != tt.wantScore {
				t.Errorf("Round.score = %v, want %v", score, tt.wantScore)
			}
		})
	}
}

func TestCounter2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "A Y",
			input: "A Y",
			want:  4,
		},
		{
			name:  "B X",
			input: "B X",
			want:  1,
		},
		{
			name:  "C Z",
			input: "C Z",
			want:  7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			round := parse(tt.input)
			if got := round.Counter2(); got != tt.want {
				t.Errorf("calcScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
