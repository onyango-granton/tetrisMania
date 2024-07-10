package utilities

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		board       [][]string
		tetrominoes [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "Box", args: args{board: [][]string{{"..", ".."}}, tetrominoes: [][]string{{"AA", "AA"}}}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.board, tt.args.tetrominoes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
