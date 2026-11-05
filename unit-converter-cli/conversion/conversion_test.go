package conversion

import "testing"

func TestToCelsius(t *testing.T) {
	conversionTests := []struct {
		name     string
		recieved float64
		expected float64
	}{
		{name: "32 degrees fahrenheit to 0 degrees celsius", recieved: 32, expected: 0},
		{name: "98.6 degrees fahrenheit to 37 degrees celsius", recieved: 98.6, expected: 37},
	}

	for _, tt := range conversionTests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToCelsius(tt.recieved)
			if got != tt.expected {
				t.Errorf("expected '%.2f' but got '%.2f'", tt.expected, got)
			}
		})
	}
}
