package math

import "testing"

func TestTruncateFloat64(t *testing.T) {
	v := 3.14159265
	result := TruncateFloat64(v, 2)
	if result != 3.14 {
		t.Fatalf("Expect: 3.14,but actual: %f", result)
	}
}
