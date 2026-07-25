/*
	Package weather allows a user to produce a pretty
 	formatted weather condition string that includes the
 	provided weather condition and provided city.
*/
package weather

// CurrentCondition package level variable to store the current condition.
var CurrentCondition string
// CurrentLocation package level variable to store the current location.
var CurrentLocation string

// Forecast generate a pretty printed string for the given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
