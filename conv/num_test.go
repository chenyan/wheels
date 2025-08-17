package conv

import (
	"reflect"
	"testing"
)

func TestParseI64s(t *testing.T) {
	tests := []struct {
		s    string
		sep  byte
		want []int64
	}{
		{
			s:    "1,2,3",
			sep:  ',',
			want: []int64{1, 2, 3},
		},
		{
			s:    ",",
			sep:  ',',
			want: []int64{},
		},
		{
			s:    "",
			sep:  ',',
			want: []int64{},
		},
		{
			s:    ",,",
			sep:  ',',
			want: []int64{},
		},
		{
			s:    "1,2,3,",
			sep:  ',',
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got, err := ParseI64s(tt.s, tt.sep)
			if err != nil {
				t.Errorf("ParseI64s() = %v, want nil", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseI64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkParseI64s(b *testing.B) {
	for b.Loop() {
		ParseI64s("1,2,3", ',')
	}
}

func TestJoinI64s(t *testing.T) {
	tests := []struct {
		i64s []int64
		sep  byte
		want string
	}{
		{
			i64s: []int64{1, 2, 3},
			sep:  ',',
			want: "1,2,3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := JoinI64s(tt.i64s, tt.sep)
			if got != tt.want {
				t.Errorf("JoinI64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkJoinI64s(b *testing.B) {
	for b.Loop() {
		JoinI64s([]int64{1, 2, 3}, ',')
	}
}
