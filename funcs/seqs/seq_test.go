package seqs

import (
	"reflect"
	"strconv"
	"testing"
)

func TestApply(t *testing.T) {
	type args struct {
		ts []int
		f  func(int) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "apply function to each element of the slice",
			args: args{
				ts: []int{1, 2, 3},
				f:  func(x int) int { return x * 2 },
			},
			want: []int{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Apply(tt.args.ts, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMap(t *testing.T) {
	type args struct {
		ts []int
		f  func(int) (string, int)
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "convert slice to map",
			args: args{
				ts: []int{1, 2, 3},
				f:  func(x int) (string, int) { return strconv.Itoa(x), x },
			},
			want: map[string]int{"1": 1, "2": 2, "3": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMap(tt.args.ts, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
