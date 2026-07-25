package space

type Planet string

var planetYearRatios = map[Planet]float64{
	"Mercury": 0.2408467 * 31557600,
	"Venus":   0.61519726 * 31557600,
	"Earth":   1.0 * 31557600,
	"Mars":    1.8808158 * 31557600,
	"Jupiter": 11.862615 * 31557600,
	"Saturn":  29.447498 * 31557600,
	"Uranus":  84.016846 * 31557600,
	"Neptune": 164.79132 * 31557600,
}

// Age returns the number of years old someone would be on the given planet,
// based on the number of seconds they've been alive.
func Age(seconds float64, planet Planet) float64 {
	planetRatio, ok := planetYearRatios[planet]
	if !ok {
		return -1 // error: invalid planet
	}
	return seconds / planetRatio
}
