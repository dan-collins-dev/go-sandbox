package conversion

func ToCelsius(tempInF float64) float64 {
	return (tempInF - 32) * (5.0 / 9.0)
}

func ToFahrenheit(tempInC float64) float64 {
	return tempInC*(9.0/5.0) + 32.0
}

func ToMiles(distInKM float64) (float64, error) {
	return 0.0, nil
}

func ToKilometers(distInMI float64) (float64, error) {
	return 0.0, nil
}
