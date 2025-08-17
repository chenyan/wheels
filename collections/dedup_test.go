package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDedup(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 2, 1}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			got := Dedup(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dedup() = %v, want %v", got, tt.want)
			}
		})
	}
}
