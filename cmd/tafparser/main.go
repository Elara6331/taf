package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/alecthomas/repr"
	"github.com/spf13/pflag"
	"go.elara.ws/logger"
	"go.elara.ws/logger/log"
	"go.elara.ws/taf"
	"go.elara.ws/taf/units"
)

func init() {
	log.Logger = logger.NewPretty(os.Stderr)
}

func main() {
	pretty := pflag.BoolP("pretty", "p", true, "Pretty-print the JSON output")
	printGo := pflag.BoolP("print-go", "G", false, "Print Go code instead of JSON")
	convertDist := pflag.StringP("convert-distance", "d", "", "Convert all the distances to the given unit. (valid units: mi, m, km)")
	convertSpd := pflag.StringP("convert-speed", "s", "", "Convert all the speeds to the given unit. (valid units: m/s, kph, kts, mph)")
	identifier := pflag.StringP("identifier", "i", "", "Automatically fetch the TAF report for the specified ICAO identifier")
	pflag.Parse()

	var opts taf.Options

	if *convertDist != "" {
		d, ok := units.ParseDistance(*convertDist)
		if !ok {
			log.Fatal("Invalid distance unit").Send()
		}
		opts.DistanceUnit = d
	}

	if *convertSpd != "" {
		s, ok := units.ParseSpeed(*convertSpd)
		if !ok {
			log.Fatal("Invalid speed unit").Send()
		}
		opts.SpeedUnit = s
	}

	var r io.Reader
	var err error
	if pflag.NArg() > 0 {
		fl, err := os.Open(pflag.Arg(0))
		if err != nil {
			log.Fatal("Error opening file").Err(err).Send()
		}
		defer fl.Close()
		r = fl
	} else if *identifier != "" {
		// Identifiers must be uppercase
		*identifier = strings.ToUpper(*identifier)
		// Get the TAF report from aviationweather.gov's beta endpoint
		res, err := http.Get("https://aviationweather.gov/cgi-bin/data/taf.php?ids=" + *identifier)
		if err != nil {
			log.Fatal("Error getting TAF report").Err(err).Send()
		}
		// The backend doesn't return an error for non-existent reports, so check the content length instead
		if res.ContentLength == 0 {
			log.Fatal("Couldn't find a TAF report for the specified airport").Str("id", *identifier).Send()
		}
		defer res.Body.Close()
		r = res.Body
	} else {
		r = os.Stdin
	}

	fc, err := taf.DecodeWithOptions(r, opts)
	if err != nil {
		log.Fatal("Error parsing TAF data").Err(err).Send()
	}

	if *printGo {
		repr.New(os.Stdout, repr.ScalarLiterals()).Println(fc)
	} else {
		enc := json.NewEncoder(os.Stdout)
		if *pretty {
			enc.SetIndent("", "  ")
		}
		err = enc.Encode(fc)
		if err != nil {
			log.Fatal("Error encoding forecast").Err(err).Send()
		}
	}
}
