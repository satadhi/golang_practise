// Package weather provides functionality to describe the current weather
// conditions for a given location.
package weather

var (
	// CurrentCondition represents the current weather condition of a location.
	CurrentCondition string

	// CurrentLocation represents the name of the location for which the weather is being reported.
	CurrentLocation string
)

// Forecast returns a formatted string describing the current weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}