// Package weather provides tools to forecast weather in city.
package weather

// CurrentCondition is a variable for the current condition.
var CurrentCondition string
//CurrentLocation is a variable for the current location.
var CurrentLocation string

// Forecast takes in a city and conditions string values and returns a string value explaining the current weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
