package taf

import (
	"io"
	"io/fs"
	"math/big"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/participle/v2"
	"go.elara.ws/taf/airports"
	"go.elara.ws/taf/internal/parser"
	"go.elara.ws/taf/units"
)

// DecodeString decodes a TAF string and returns a Forecast.
// This is equivalent to Decode(strings.NewReader(s)).
func DecodeString(s string) (*Forecast, error) {
	return Decode(strings.NewReader(s))
}

// DecodeFile decodes a TAF string and returns a Forecast.
// This is equivalent to opening a file and passing it
// to Decode().
func DecodeFile(path string) (*Forecast, error) {
	fl, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fl.Close()

	return Decode(fl)
}

// Decode decodes the data in a reader using default options and
// returns a Forecast
func Decode(r io.Reader) (*Forecast, error) {
	return DecodeWithOptions(r, Options{})
}

// Options contains options for the decoder
type Options struct {
	// If this is set, all distance units in the forecast
	// will be converted to the given unit
	DistanceUnit units.Distance

	// If this is set, all speed units in the forecast will
	// be converted to the given unit
	SpeedUnit units.Speed

	// The Year field is used to calculate the full date that this
	// report was published. If it's unset, the current year will be used.
	Year int

	// The Month field is used to calculate the full date that this
	// report was published. If it's unset, the current month will be used.
	Month time.Month
}

// DecodeWithOptions decodes the data in a reader and returns a Forecast
func DecodeWithOptions(r io.Reader, opts Options) (*Forecast, error) {
	filename := "unknown"
	switch r := r.(type) {
	case *os.File:
		filename = r.Name()
	case fs.File:
		fi, err := r.Stat()
		if err == nil {
			filename = fi.Name()
		}
	case *strings.Reader:
		filename = "string"
	}

	if opts.Year == 0 {
		opts.Year = time.Now().Year()
	}

	if opts.Month == 0 {
		opts.Month = time.Now().Month()
	}

	ast, err := parser.Parser.Parse(filename, r)
	if err != nil {
		return nil, err
	}

	fc := &Forecast{}
	out := reflect.ValueOf(fc).Elem()

	for _, item := range ast.Items {
		switch {
		case item.ID != nil:
			fc.Identifier = *item.ID
			if a, ok := airports.Airports[fc.Identifier]; ok {
				fc.Airport = a
			}
		case item.Time != nil:
			t, err := parseTime(*item.Time, opts.Month, opts.Year)
			if err != nil {
				return nil, participle.Errorf(item.Pos, "time: %s", err)
			}
			setField(out, "PublishTime", t)

			// The Time item always comes with a Valid as well because
			// of the way it's parsed into the AST
			vp, err := parseValid(item.Valid, opts.Month, opts.Year)
			if err != nil {
				return nil, participle.Errorf(item.Pos, "time: %s", err)
			}
			setField(out, "Valid", vp)
		case item.Weather != nil:
			appendField(out, "Weather", Weather{
				Modifier:      convertModifier(item.Weather.Modifier),
				Descriptor:    convertDescriptor(item.Weather.Descriptor),
				Precipitation: convertPrecipitation(item.Weather.Precipitation),
				Obscuration:   convertObscuration(item.Weather.Obscuration),
				Phenomenon:    convertPhenomenon(item.Weather.Other),
			})
		case item.Vicinity != nil:
			appendField(out, "Weather", Weather{
				Vicinity:      true,
				Descriptor:    convertDescriptor(item.Vicinity.Descriptor),
				Precipitation: convertPrecipitation(item.Vicinity.Precipitation),
			})
		case item.SkyCondition != nil:
			var altitude int
			if item.SkyCondition.Altitude != "" {
				altitude, err = strconv.Atoi(item.SkyCondition.Altitude)
				if err != nil {
					return nil, participle.Errorf(item.SkyCondition.Pos, "sky: %s", err)
				}
			}

			appendField(out, "SkyCondition", SkyCondition{
				Altitude:  altitude * 100, // Scale factor for altitude is 100
				Type:      convertSkyConditionType(item.SkyCondition.Type),
				CloudType: convertCloudType(item.SkyCondition.CloudType),
			})
		case item.Temperature != nil:
			vt, err := parseValidTime(item.Temperature.Time, opts.Month, opts.Year)
			if err != nil {
				return nil, participle.Errorf(item.Temperature.Pos, "temp: %s", err)
			}

			val, err := strconv.Atoi(item.Temperature.Value)
			if err != nil {
				return nil, participle.Errorf(item.Temperature.Pos, "temp: %s", err)
			}

			appendField(out, "Temperature", Temperature{
				Type:  convertTemperatureType(item.Temperature.Type),
				Time:  vt,
				Value: val,
			})
		case item.Visibility != nil:
			// This value may have a space at the end if there's no unit
			item.Visibility.Value = strings.TrimSpace(item.Visibility.Value)

			// Create a new rational number
			ratNum := new(big.Rat)
			// If there's a space, this is a mixed number, split it at the space
			if before, after, ok := strings.Cut(item.Visibility.Value, " "); ok {
				// Set the rational number to the fraction of the mixed number
				ratNum, ok = ratNum.SetString(after)
				if !ok {
					return nil, participle.Errorf(item.Visibility.Pos, "visibility: invalid fraction %q", after)
				}

				// Create a new rational number and set it to the whole part of
				// the mixed number
				add, ok := new(big.Rat).SetString(before)
				if !ok {
					return nil, participle.Errorf(item.Visibility.Pos, "visibility: invalid whole number %q", before)
				}

				// Add the whole part to the fractional part
				ratNum = ratNum.Add(ratNum, add)
			} else {
				// There's no space, so this is just a fraction or a whole number.
				// Just set the rational number to the whole string.
				ratNum, ok = ratNum.SetString(before)
				if !ok {
					return nil, participle.Errorf(item.Visibility.Pos, "visibility: invalid fraction %q", after)
				}
			}

			// If there's no unit, set the unit to meters
			if item.Visibility.Unit == "" {
				item.Visibility.Unit = "M"
			}

			unit, ok := units.ParseDistance(item.Visibility.Unit)
			if !ok {
				return nil, participle.Errorf(item.Visibility.Pos, "visibility: invalid unit %q", item.Visibility.Unit)
			}

			val, _ := ratNum.Float64()

			if opts.DistanceUnit != "" {
				val = unit.Convert(opts.DistanceUnit, val)
				unit = opts.DistanceUnit
			}

			setField(out, "Visibility", Visibility{
				Plus:  item.Visibility.Plus,
				Value: val,
				Unit:  unit,
			})
		case item.WindSpeed != nil:
			var direction int
			// If the wind speed is variable, there's no direction to worry about
			if !item.WindSpeed.Variable {
				// The length of the value must be at least 5 (3 characters for direction and 2 for speed)
				if len(item.WindSpeed.Value) < 5 {
					return nil, participle.Errorf(item.WindSpeed.Pos, "wind: invalid length (%d)", len(item.WindSpeed.Value))
				}

				// First three characters are the direction
				direction, err = strconv.Atoi(item.WindSpeed.Value[:3])
				if err != nil {
					return nil, participle.Errorf(item.WindSpeed.Pos, "wind: %s", err)
				}

				// Set the value to the last two characters so it can be processed
				// as just a speed.
				item.WindSpeed.Value = item.WindSpeed.Value[3:]

				// The direction is in degrees so it may not go above 360 or below 0
				if direction > 360 || direction < 0 {
					return nil, participle.Errorf(item.WindSpeed.Pos, "wind: invalid direction (%d)", direction)
				}
			}

			// If there was a direction, it was removed above, so now we can just
			// get the speed by parsing the string
			speed, err := strconv.Atoi(item.WindSpeed.Value)
			if err != nil {
				return nil, participle.Errorf(item.WindSpeed.Pos, "wind: %s", err)
			}

			var gusts int
			if item.WindSpeed.Gusts != "" {
				gusts, err = strconv.Atoi(item.WindSpeed.Gusts)
				if err != nil {
					return nil, participle.Errorf(item.WindSpeed.Pos, "wind: %s", err)
				}
			}

			var windshear int
			if item.WindSpeed.WindShear != "" {
				windshear, err = strconv.Atoi(item.WindSpeed.WindShear)
				if err != nil {
					return nil, participle.Errorf(item.WindSpeed.Pos, "wind: %s", err)
				}
			}

			unit, ok := units.ParseSpeed(item.WindSpeed.Unit)
			if !ok {
				return nil, participle.Errorf(item.WindSpeed.Pos, "wind: invalid unit %q", item.Visibility.Unit)
			}

			if opts.SpeedUnit != "" {
				speed = unit.Convert(opts.SpeedUnit, speed)
				if gusts != 0 {
					gusts = unit.Convert(opts.SpeedUnit, gusts)
				}
				unit = opts.SpeedUnit
			}

			setField(out, "Wind", Wind{
				Gusts:     gusts,
				Speed:     speed,
				WindShear: windshear * 100, // Scale factor for altitude is 100
				Direction: Direction{
					Variable: item.WindSpeed.Variable,
					Value:    direction,
				},
				Unit: unit,
			})
		case item.Flag != nil:
			switch {
			case item.Flag.CAVOK:
				appendField(out, "Flags", CeilingAndVisibilityOK)
			}
		case item.Change != nil:
			ch := &Change{
				Type: convertChangeType(item.Change.Type),
			}

			// FM changes don't have a valid pair, they only come with a single time string
			if ch.Type == From {
				t, err := parseTime(item.Change.Time, opts.Month, opts.Year)
				if err != nil {
					return nil, participle.Errorf(item.Change.Pos, "changes: %s", err)
				}
				ch.Valid = ValidPair{From: t}
			} else {
				vp, err := parseValid(item.Change.Valid, opts.Month, opts.Year)
				if err != nil {
					return nil, participle.Errorf(item.Change.Pos, "changes: %s", err)
				}
				ch.Valid = vp
			}

			fc.Changes = append(fc.Changes, ch)

			// Set out to the change value so that future mutations
			// happen to the change rather than the root forecast.
			out = reflect.ValueOf(ch).Elem()
		case item.Probability != nil:
			// If the time is empty, this probability belongs to the
			// previous change.
			if item.Probability.Time == "" {
				prob, err := strconv.Atoi(item.Probability.Value)
				if err != nil {
					return nil, participle.Errorf(item.Probability.Pos, "prob: %s", err)
				}
				setField(out, "Probability", prob)
			} else {
				pr := &Probability{}

				// The probability time string must have 4 characters,
				// 2 for starting time and 2 for ending time
				if len(item.Probability.Time) < 4 {
					return nil, participle.Errorf(item.Probability.Pos, "prob: invalid time %q", item.Probability.Time)
				}

				startStr := item.Probability.Time[:2]
				endStr := item.Probability.Time[2:]

				start, err := strconv.Atoi(startStr)
				if err != nil {
					return nil, participle.Errorf(item.Probability.Pos, "prob: %s", err)
				}

				end, err := strconv.Atoi(endStr)
				if err != nil {
					return nil, participle.Errorf(item.Probability.Pos, "prob: %s", err)
				}

				t := fc.PublishTime
				pr.Valid = ValidPair{
					From: time.Date(t.Year(), t.Month(), t.Day(), start, 0, 0, 0, time.UTC),
					To:   time.Date(t.Year(), t.Month(), t.Day(), end, 0, 0, 0, time.UTC),
				}

				// Get the duration by subtracting the from time from the to time
				pr.Valid.Duration = pr.Valid.To.Sub(pr.Valid.From)

				fc.Probabilities = append(fc.Probabilities, pr)

				// Set out to the probability value so that future mutations
				// happen to the probability rather than the root forecast.
				out = reflect.ValueOf(pr).Elem()
			}

		}
	}

	return fc, nil
}

// setField sets a field of a struct to a value.
//
// This is used to allow mutations to happen on either
// the root forecast or a change or probability. It makes it
// easier to handle the different types.
func setField(rv reflect.Value, name string, to any) {
	rv.FieldByName(name).Set(reflect.ValueOf(to))
}

// appendField appends a value to a slice in the field of a struct.
//
// This is used to allow mutations to happen on either
// the root forecast or a change or probability. It makes it
// easier to handle the different types.
func appendField(rv reflect.Value, name string, items ...any) {
	f := rv.FieldByName(name)
	f.Set(reflect.Append(f, anyToValues(items)...))
}

// anyToValues converts a slice of any type to a slice
// of reflect values.
func anyToValues(items []any) []reflect.Value {
	out := make([]reflect.Value, len(items))
	for i, item := range items {
		out[i] = reflect.ValueOf(item)
	}
	return out
}
