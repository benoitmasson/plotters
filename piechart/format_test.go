package piechart

import "testing"

func TestFormatFloat(t *testing.T) {
	type test struct {
		name     string
		value    float64
		expected string
	}
	testData := []*test{
		{"zero", 0, "0"},
		{"1-digit integer", 1, "1"},
		{"multiple of 10", 10, "10"},
		{"large integer", 1_000_000, "1000000"},
		{"1-digit decimal", 0.1, "0.1"},
		{"2-digits decimal", 0.12, "0.12"},
		{"multiple-digits decimal", 0.12345, "0.12"},
		{"multiple-digits rounded decimal", 0.12789, "0.13"},
		{"generic decimal", 12.3456789, "12.35"},
		{"large decimal", 1_234.56789, "1234.57"},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			actual := formatFloat(data.value)
			if actual != data.expected {
				t.Errorf("formatFloat(%f) returned %s, expected %s", data.value, actual, data.expected)
			}
		})
	}
}
