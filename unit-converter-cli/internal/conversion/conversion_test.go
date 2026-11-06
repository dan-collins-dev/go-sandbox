package conversion

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestToCelsius(t *testing.T) {
	conversionTests := []struct {
		name     string
		recieved float64
		expected float64
	}{
		{name: "32 degrees fahrenheit to 0 degrees celsius", recieved: 32, expected: 0},
		{name: "94.4 degrees fahrenheit to 34.6 degrees celsius", recieved: 94.4, expected: 34.6},
		{name: "44 degrees fahrenheit to 6.6 degrees celsius", recieved: 44, expected: 6.66},
	}

	for _, tt := range conversionTests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToCelsius(tt.recieved)

			toleranceOpt := cmpopts.EquateApprox(1e-2, 1e-9)

			if diff := cmp.Diff(tt.expected, got, toleranceOpt); diff != "" {
				t.Errorf("expected '%f' but got '%f'", tt.expected, got)
			}
		})
	}
}

func TestToFahrenheit(t *testing.T) {
	conversionTests := []struct {
		name     string
		recieved float64
		expected float64
	}{
		{name: "0 degrees celsius to 32 degrees fahrenheight", recieved: 0, expected: 32},
		{name: "34.6 degrees celsius to 94.28 degrees fahrenheit", recieved: 34.6, expected: 94.28},
		{name: "6.6 degrees celsius to 43.8 degrees fahrenheit", recieved: 6.6, expected: 43.8},
	}

	for _, tt := range conversionTests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToFahrenheit(tt.recieved)

			toleranceOpt := cmpopts.EquateApprox(1e-2, 1e-9)

			if diff := cmp.Diff(tt.expected, got, toleranceOpt); diff != "" {
				t.Errorf("expected '%f' but got '%f'", tt.expected, got)
			}
		})
	}
}
