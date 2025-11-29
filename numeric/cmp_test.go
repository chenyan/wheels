package numeric

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		xs   []int
		want int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{3, 2, 1}, 3},
		{[]int{1, 3, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.xs), func(t *testing.T) {
			got := Max(tt.xs...)
			if got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		xs   []int
		want int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{3, 2, 1}, 1},
		{[]int{1, 3, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.xs), func(t *testing.T) {
			got := Min(tt.xs...)
			if got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}
