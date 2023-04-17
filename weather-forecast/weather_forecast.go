// Package weather is made to forecast the
// current weather condition of various cities
// in Goblinocus.
package weather

// CurrentCondition is a string describing the current
// weather condition. It can be used in the Forecast function.
var CurrentCondition string

// CurrentLocation is a string with the city's name
// that can be used in the Forecast function.
var CurrentLocation string

// Forecast function returns the current weather
// condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
