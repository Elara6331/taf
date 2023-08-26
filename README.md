# taf

[![Go Report Card](https://goreportcard.com/badge/go.elara.ws/taf)](https://goreportcard.com/report/go.elara.ws/taf)
[![Go Reference](https://pkg.go.dev/badge/go.elara.ws/taf.svg)](https://pkg.go.dev/go.elara.ws/taf)

This is a library and command-line tool that parses and decodes [TAF forecasts](https://en.wikipedia.org/wiki/Terminal_aerodrome_forecast).

TAF stands for Terminal Aerodrome Forecast. It's the weather forecast format used in aviation. TAF reports are useful as a free source of accurate weather.

Here's an example of a TAF report from JFK airport:

```
KJFK 212335Z 2200/2306 33012G18KT P6SM FEW060 BKN250
  FM220300 36014KT P6SM FEW060 SCT150
  FM221400 01015G21KT P6SM SCT060
  FM221900 04011KT P6SM SCT060
  FM230000 03007KT P6SM FEW060
  FM230300 35006KT P6SM FEW060
```

Try parsing it by installing the tafparser tool using

```bash
go install go.elara.ws/taf/cmd/tafparser@latest
```

and then running

```bash
tafparser <<EOF
KJFK 212335Z 2200/2306 33012G18KT P6SM FEW060 BKN250
  FM220300 36014KT P6SM FEW060 SCT150
  FM221400 01015G21KT P6SM SCT060
  FM221900 04011KT P6SM SCT060
  FM230000 03007KT P6SM FEW060
  FM230300 35006KT P6SM FEW060
EOF
```

That should return a JSON object containing all the decoded data from the TAF report.

You can also give the `tafparser` tool a file to read from using `tafparser file.txt`.

Units in TAF reports are inconsistent between different countries. `tafparser` can convert the units for you! Just pass it the units you want to use for speed and/or distance like so:

```bash
tafparser -s m/s -d m
```

This tells `tafparser` to convert all speed units to meters per second and distance units to meters.

`tafparser` can also fetch TAF reports for you using the [aviationweather.gov](https://aviationweather.gov) site. Use the `-i <identifier>` flag to tell it to do that, like so:

```bash
tafparser -i EGLL
```

That should automatically fetch the report for London Heathrow and parse it.