package utilities

import "testing"

func TestValid(t *testing.T) {
	type args struct {
		tetro [][]string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "valid", args : args{tetro: [][]string{{"...#","...#","...#","...#"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Valid(tt.args.tetro)
		})
	}
}
