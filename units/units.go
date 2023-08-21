package units

import (
	"strings"
)

// Speed represents a unit of speed
type Speed string

// Speed units
const (
	MetersPerSecond   Speed = "MPS"
	KilometersPerHour Speed = "KMH"
	Knots             Speed = "KT"
)

var speedNames = map[Speed]string{
	MetersPerSecond:   "m/s",
	KilometersPerHour: "kph",
	Knots:             "kts",
}

func (su Speed) String() string {
	name, ok := speedNames[su]
	if !ok {
		return "<unknown>"
	}
	return name
}

// Convert converts a value from one unit to another
func (sf Speed) Convert(st Speed, val int) int {
	switch {
	case sf == MetersPerSecond && st == KilometersPerHour:
		return int(float64(val) * 3.6)
	case sf == KilometersPerHour && st == MetersPerSecond:
		return int(float64(val) / 3.6)
	case sf == Knots && st == KilometersPerHour:
		return int(float64(val) * 1.852)
	case sf == KilometersPerHour && st == Knots:
		return int(float64(val) / 1.852)
	case sf == MetersPerSecond && st == Knots:
		return int(float64(val) * 1.94384)
	case sf == Knots && st == MetersPerSecond:
		return int(float64(val) / 1.94384)
	default:
		return val
	}
}

// ParseSpeed parses a speed value. Valid inputs include:
// mps, m/s, kmh, kph, kt, and kts.
// This function is case-insensitive.
func ParseSpeed(s string) (Speed, bool) {
	if _, ok := speedNames[Speed(s)]; ok {
		return Speed(s), true
	}

	for su, name := range speedNames {
		if strings.EqualFold(s, name) ||
			strings.EqualFold(s, string(su)) {
			return su, true
		}
	}

	return "", false
}

// Distance represents a unit of distance
type Distance string

// Distance units
const (
	StatuteMiles Distance = "SM"
	Meters       Distance = "M"
)

var distanceNames = map[Distance]string{
	StatuteMiles: "mi",
	Meters:       "m",
}

func (du Distance) String() string {
	name, ok := distanceNames[du]
	if !ok {
		return "<unknown>"
	}
	return name
}

// Convert converts a value from one unit to another
func (df Distance) Convert(dt Distance, val float64) float64 {
	switch {
	case df == StatuteMiles && dt == Meters:
		return val * 1609
	case df == Meters && dt == StatuteMiles:
		return val / 1609
	default:
		return val
	}
}

// ParseDistance parses a speed value. Valid inputs include:
// sm, mi, and m. This function is case-insensitive.
func ParseDistance(s string) (Distance, bool) {
	if _, ok := distanceNames[Distance(s)]; ok {
		return Distance(s), true
	}

	for d, name := range distanceNames {
		if strings.EqualFold(s, name) ||
			strings.EqualFold(s, string(d)) {
			return d, true
		}
	}

	return "", false
}
