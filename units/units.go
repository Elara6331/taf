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
	MilesPerHour      Speed = "MilesPerHour"
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
	case sf == MilesPerHour && st == KilometersPerHour:
		return int(float64(val) * 1.60934)
	case sf == KilometersPerHour && st == MilesPerHour:
		return int(float64(val) / 1.60934)
	case sf == MilesPerHour && st == MetersPerSecond:
		return int(float64(val) * 0.44704)
	case sf == MetersPerSecond && st == MilesPerHour:
		return int(float64(val) / 0.44704)
	case sf == MilesPerHour && st == Knots:
		return int(float64(val) * 0.868976)
	case sf == Knots && st == MilesPerHour:
		return int(float64(val) / 0.868976)
	default:
		return val
	}
}

// ParseSpeed parses a speed value. Valid inputs include:
// mps, m/s, kmh, kph, kt, kts, mph, and milesperhour.
// This function is case-insensitive.
func ParseSpeed(s string) (Speed, bool) {
	switch strings.ToLower(s) {
	case "m/s", "mps", "meterspersecond", "meters per second", "metrespersecond", "metres per second":
		return MetersPerSecond, true
	case "kmh", "kph", "km/h", "kilometersperhour", "kilometers per hour", "kilometresperhour", "kilometres per hour":
		return KilometersPerHour, true
	case "kt", "kts", "knot", "knots":
		return Knots, true
	case "mph", "milesperhour", "miles per hour":
		return MilesPerHour, true
	default:
		return "", false
	}
}

// Distance represents a unit of distance
type Distance string

// Distance units
const (
	Miles      Distance = "Miles"
	Meters     Distance = "Meters"
	Kilometers Distance = "Kilometers"
)

// Convert converts a value from one unit to another
func (df Distance) Convert(dt Distance, val float64) float64 {
	switch {
	case df == Miles && dt == Meters:
		return val * 1609
	case df == Meters && dt == Miles:
		return val / 1609
	case df == Kilometers && dt == Meters:
		return val * 1000
	case df == Meters && dt == Kilometers:
		return val / 1000
	case df == Miles && dt == Kilometers:
		return val * 1.60934
	case df == Kilometers && dt == Miles:
		return val / 1.60934
	default:
		return val
	}
}

// ParseDistance parses a distance value. Valid inputs include:
// sm, mi, m, km, and kilometers.
// This function is case-insensitive.
func ParseDistance(s string) (Distance, bool) {
	switch strings.ToLower(s) {
	case "sm", "mi", "mile", "miles":
		return Miles, true
	case "m", "meter", "meters", "metre", "metres":
		return Meters, true
	case "km", "kilometer", "kilometers", "kilometre", "kilometres":
		return Kilometers, true
	default:
		return "", false
	}
}
