package simplemath

import "testing"

func TestC3sum(t *testing.T) {
	C3sum()
}

func TestSqrt(t *testing.T) {
	const in, out = 100, 10
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, expect %v", in, x, out)
	}
}
