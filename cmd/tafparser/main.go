package main

import (
	"encoding/json"
	"os"

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

	var fl *os.File
	var err error
	if pflag.NArg() > 0 {
		fl, err = os.Open(pflag.Arg(0))
		if err != nil {
			log.Fatal("Error opening file").Err(err).Send()
		}
	} else {
		fl = os.Stdin
	}

	fc, err := taf.DecodeWithOptions(fl, opts)
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
