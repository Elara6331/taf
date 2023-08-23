package units

import (
	"strings"
)

// Speed represents a unit of speed
type Speed string

// Speed units
const (
	MetersPerSecond   Speed = "MetersPerSecond"
	KilometersPerHour Speed = "KilometersPerHour"
	Knots             Speed = "Knots"
)

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
	switch strings.ToLower(s) {
	case "m/s", "mps", "meterspersecond", "meters per second":
		return MetersPerSecond, true
	case "kmh", "kph", "km/h", "kilometersperhour", "kilometers per hour":
		return KilometersPerHour, true
	case "kt", "kts", "knot", "knots":
		return Knots, true
	default:
		return "", false
	}
}

// Distance represents a unit of distance
type Distance string

// Distance units
const (
	Miles  Distance = "Miles"
	Meters Distance = "Meters"
)

// Convert converts a value from one unit to another
func (df Distance) Convert(dt Distance, val float64) float64 {
	switch {
	case df == Miles && dt == Meters:
		return val * 1609
	case df == Meters && dt == Miles:
		return val / 1609
	default:
		return val
	}
}

// ParseDistance parses a speed value. Valid inputs include:
// sm, mi, and m. This function is case-insensitive.
func ParseDistance(s string) (Distance, bool) {
	switch strings.ToLower(s) {
	case "sm", "mi", "mile", "miles":
		return Miles, true
	case "m", "meter", "meters":
		return Meters, true
	default:
		return "", false
	}
}
