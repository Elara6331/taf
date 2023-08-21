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

func parseTime(s string) (time.Time, error) {
	t, err := time.Parse(TimeFormat, s)
	if err != nil {
		return time.Time{}, err
	}
	now := time.Now().UTC()
	return t.AddDate(now.Year(), int(now.Month()), 0), nil
}

func parseValid(v *parser.ValidPair) (ValidPair, error) {
	start, err := parseValidTime(v.Start)
	if err != nil {
		return ValidPair{}, err
	}

	end, err := parseValidTime(v.End)
	if err != nil {
		return ValidPair{}, err
	}

	return ValidPair{
		From:     start,
		To:       end,
		Duration: end.Sub(start),
	}, nil
}

func parseValidTime(s string) (time.Time, error) {
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

	now := time.Now().UTC()
	return t.AddDate(now.Year(), int(now.Month()), addDays), nil
}
