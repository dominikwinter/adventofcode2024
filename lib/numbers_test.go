package lib

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		{input: int(-5), expected: int(5)},
		{input: int32(-5), expected: int32(5)},
		{input: int64(-5), expected: int64(5)},
		{input: float32(-5.5), expected: float32(5.5)},
		{input: float64(-5.5), expected: float64(5.5)},
		{input: int(5), expected: int(5)},
		{input: int32(5), expected: int32(5)},
		{input: int64(5), expected: int64(5)},
		{input: float32(5.5), expected: float32(5.5)},
		{input: float64(5.5), expected: float64(5.5)},
	}

	for _, test := range tests {
		switch v := test.input.(type) {
		case int:
			if result := Abs(v); result != test.expected {
				t.Errorf("Abs(%d) = %d; want %d", v, result, test.expected)
			}
		case int32:
			if result := Abs(v); result != test.expected {
				t.Errorf("Abs(%d) = %d; want %d", v, result, test.expected)
			}
		case int64:
			if result := Abs(v); result != test.expected {
				t.Errorf("Abs(%d) = %d; want %d", v, result, test.expected)
			}
		case float32:
			if result := Abs(v); result != test.expected {
				t.Errorf("Abs(%f) = %f; want %f", v, result, test.expected)
			}
		case float64:
			if result := Abs(v); result != test.expected {
				t.Errorf("Abs(%f) = %f; want %f", v, result, test.expected)
			}
		}
	}
}
