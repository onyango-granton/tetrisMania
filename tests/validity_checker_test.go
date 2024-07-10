package tests

import (
	"testing"
	"tetris/utilities"
)

func TestValid(t *testing.T) {
	type args struct {
		tetro [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "valid", args: args{tetro: [][]string{{"...#", "...#", "...#", "...#"}}}, want: "ok"},
		{name: "Invalid", args: args{tetro: [][]string{{"....", "....", "....", "...."}}}, want: "Invalid File"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utilities.Valid(tt.args.tetro); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
