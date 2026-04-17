// Package weather provides tools to check the
// temperature of the weather.
package weather

var (
// CurrentCondition represents the city.
	CurrentCondition string
    
// CurrentLocation ... represents the condition.
	CurrentLocation  string
)

// Forecast returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
