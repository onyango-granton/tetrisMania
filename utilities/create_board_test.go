package utilities

import (
	"reflect"
	"testing"
)

func TestCreateInitialBoard(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "Three", args: args{size: 3}, want: [][]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}}},
		{name: "Two", args: args{size: 2}, want: [][]string{{".", "."}, {".", "."}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateInitialBoard(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInitialBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
