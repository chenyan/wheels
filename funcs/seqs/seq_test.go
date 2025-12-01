package seqs

import (
	"iter"
	"reflect"
	"strconv"
	"testing"

	"github.com/chenyan/wheels/types"
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

func TestMap(t *testing.T) {
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
			name: "map function to each element of the slice",
			args: args{
				ts: []int{1, 2, 3},
				f:  func(x int) int { return x * 2 },
			},
			want: []int{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.ts, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSeq(t *testing.T) {
	type args struct {
		ts []int
	}
	tests := []struct {
		name string
		args args
		want iter.Seq[int]
	}{
		{
			name: "convert slice to sequence",
			args: args{
				ts: []int{1, 2, 3},
			},
			want: ToSeq([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSeq(tt.args.ts)
			items := make(map[int]struct{})
			got(func(t int) bool {
				items[t] = struct{}{}
				return true
			})
			if len(items) != 3 {
				t.Errorf("ToSeq() yielded %d items, want 3", len(items))
			}
		})
	}
}

func TestZip(t *testing.T) {
	type args struct {
		as iter.Seq[int]
		bs iter.Seq[string]
	}
	tests := []struct {
		name string
		args args
		want iter.Seq2[int, string]
	}{
		{
			name: "zip two sequences",
			args: args{
				as: ToSeq([]int{1, 2, 3, 4}),
				bs: ToSeq([]string{"a", "b", "c", "d"}),
			},
			want: func(yield func(int, string) bool) {
				yield(1, "a")
				yield(2, "b")
				yield(3, "c")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Zip(tt.args.as, tt.args.bs)
			items := make(map[types.Pair[int, string]]struct{})
			got(func(a int, b string) bool {
				items[types.NewPair(a, b)] = struct{}{}
				return true
			})
			if len(items) != 4 {
				t.Errorf("Zip() = %v, want %v", len(items), 3)
			}
		})
	}
}
