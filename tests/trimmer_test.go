package utilities

import (
	"reflect"
	"testing"
)

func TestTrimmer(t *testing.T) {
	type args struct {
		tetro [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "stick", args: args{tetro: [][]string{{"...A", "...A", "...A", "...A"}, {"....", "BB..", "BB..", "...."}}}, want: [][]string{{"A", "A", "A", "A"}, {"BB", "BB"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trimmer(tt.args.tetro); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Trimmer() = %v, want %v", got, tt.want)
			}
		})
	}
}
