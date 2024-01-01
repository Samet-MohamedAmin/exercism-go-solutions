package space

const secondsInYear = 31557600

var planetEarthYear = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Planet is planet name
type Planet string

// Age returns age in earth years for a given planet
func Age(seconds float64, p Planet) float64 {

	return seconds / (secondsInYear * planetEarthYear[p])
}
