package taf

import (
	"strings"
	"time"

	"go.elara.ws/taf/internal/parser"
)

const (
	TimeFormat  = "021504"
	ValidFormat = "0215"
)

func parseTime(s string, m time.Month, year int) (time.Time, error) {
	t, err := time.Parse(TimeFormat, s)
	if err != nil {
		return time.Time{}, err
	}
	return t.AddDate(year, int(m)-1, 0), nil
}

func parseValid(v *parser.ValidPair, m time.Month, year int) (ValidPair, error) {
	start, err := parseValidTime(v.Start, m, year)
	if err != nil {
		return ValidPair{}, err
	}

	end, err := parseValidTime(v.End, m, year)
	if err != nil {
		return ValidPair{}, err
	}

	return ValidPair{
		From:     start,
		To:       end,
		Duration: end.Sub(start),
	}, nil
}

func parseValidTime(s string, m time.Month, year int) (time.Time, error) {
	addDays := 0
	// Go doesn't know what to do with hour 24,
	// so we set it to 00 the next day
	if strings.HasSuffix(s, "24") {
		s = s[:2] + "00"
		addDays = 1
	}

	t, err := time.Parse(ValidFormat, s)
	if err != nil {
		return time.Time{}, err
	}

	return t.AddDate(year, int(m)-1, addDays), nil
}
