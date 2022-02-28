package simplemath

import "testing"

func TestSqrt(t *testing.T) {
	v := Sqrt(9)
	if v != 81 {
		t.Errorf("Sqrt(9) failed. Got %v, expect 3.", v)
	}
}
