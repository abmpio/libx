package math

import "testing"

func Test_RoundX(t *testing.T) {
	v := 3.14159265
	result := Round(v, 3)
	if result != 3.142 {
		t.Fatalf("Expect: 3.14159,but actual: %f", result)
	}
}

func Test_TruncateFloat64(t *testing.T) {
	v := 3.14159265
	result := TruncateFloat64(v, 2)
	if result != 3.14 {
		t.Fatalf("Expect: 3.14,but actual: %f", result)
	}
}
