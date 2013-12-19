package math

import "testing"

var maxIntTests = []struct {
	v    []int
	want int
}{
	{[]int{}, 0},
	{[]int{0}, 0},
	{[]int{-1}, -1},
	{[]int{0, 0}, 0},
	{[]int{1, 1}, 1},
	{[]int{1, 2}, 2},
	{[]int{2, 1}, 2},
	{[]int{-1, 1}, 1},
	{[]int{1, -1}, 1},
	{[]int{-1, 0, 1}, 1},
	{[]int{100, 42, 17, 2, 3}, 100},
}

func TestMaxInt(t *testing.T) {
	for i, tt := range maxIntTests {
		got := MaxInt(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MaxInt(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var minIntTests = []struct {
	v    []int
	want int
}{
	{[]int{}, 0},
	{[]int{0}, 0},
	{[]int{-1}, -1},
	{[]int{0, 0}, 0},
	{[]int{1, 1}, 1},
	{[]int{1, 2}, 1},
	{[]int{2, 1}, 1},
	{[]int{-1, 1}, -1},
	{[]int{1, -1}, -1},
	{[]int{-1, 0, 1}, -1},
	{[]int{100, 42, 17, 2, 3}, 2},
}

func TestMinInt(t *testing.T) {
	for i, tt := range minIntTests {
		got := MinInt(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MinInt(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
