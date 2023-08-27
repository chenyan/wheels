// BEGIN: 8d7f5a3b7c5a
package seq

import (
	"reflect"
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

// END: 8d7f5a3b7c5a
