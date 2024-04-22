package lang

import "testing"

func Test_ValueToPtr(t *testing.T) {
	v := 3.14159265
	result := ValueToPtr(v)
	if *result != 3.14159265 {
		t.Fatalf("Expect: 3.14159,but actual: %f", *result)
	}
}
