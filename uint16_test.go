package math

import "testing"

var maxUint16Tests = []struct {
	v    []uint16
	want uint16
}{
	{[]uint16{}, 0},
	{[]uint16{0}, 0},
	{[]uint16{0, 0}, 0},
	{[]uint16{1, 1}, 1},
	{[]uint16{1, 2}, 2},
	{[]uint16{2, 1}, 2},
	{[]uint16{1, 0, 2}, 2},
	{[]uint16{100, 200, 42, 17, 2, 3}, 200},
}

func TestMaxUint16(t *testing.T) {
	for i, tt := range maxUint16Tests {
		got := MaxUint16(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MaxUint16(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}

var minUint16Tests = []struct {
	v    []uint16
	want uint16
}{
	{[]uint16{}, 0},
	{[]uint16{0}, 0},
	{[]uint16{0, 0}, 0},
	{[]uint16{1, 1}, 1},
	{[]uint16{1, 2}, 1},
	{[]uint16{2, 1}, 1},
	{[]uint16{1, 0, 2}, 0},
	{[]uint16{100, 200, 42, 17, 2, 3}, 2},
}

func TestMinUint16(t *testing.T) {
	for i, tt := range minUint16Tests {
		got := MinUint16(tt.v...)
		if tt.want != got {
			t.Errorf("%d: MinUint16(%v) = %v, want %v", i+1, tt.v, got, tt.want)
		}
	}
}
