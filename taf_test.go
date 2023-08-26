package taf

import (
	"strings"
	"testing"
	"time"

	"github.com/go-test/deep"
	"go.elara.ws/taf/airports"
	"go.elara.ws/taf/units"
)

func TestKLAX(t *testing.T) {
	const data = `KLAX 212011Z 2120/2224 26012KT P6SM FEW035 SCT050 SCT060
  FM212200 25010KT P6SM SCT040
  FM220300 VRB03KT P6SM BKN025
  FM221000 VRB03KT P6SM OVC025
  FM221700 26006KT P6SM BKN025
  FM222000 26012KT P6SM SCT030`

	expected := &Forecast{
		Identifier: "KLAX",
		Airport: airports.Airport{
			ICAO:      "KLAX",
			IATA:      "LAX",
			Name:      "Los Angeles International Airport",
			City:      "Los Angeles",
			State:     "California",
			Country:   "US",
			Elevation: 125,
			Latitude:  33.94250107,
			Longitude: -118.4079971,
			Timezone:  "America/Los_Angeles",
		},
		PublishTime: time.Date(2023, time.August, 21, 20, 11, 0, 0, time.UTC),
		Valid: ValidPair{
			From:     time.Date(2023, time.August, 21, 20, 0, 0, 0, time.UTC),
			To:       time.Date(2023, time.August, 23, 0, 0, 0, 0, time.UTC),
			Duration: time.Duration(100800000000000),
		},
		Visibility: Visibility{
			Plus:  true,
			Value: 6,
			Unit:  units.Miles,
		},
		Wind: Wind{
			Direction: Direction{
				Value: 260,
			},
			Speed: 12,
			Unit:  units.Knots,
		},
		SkyCondition: []SkyCondition{
			{
				Type:     Few,
				Altitude: 3500,
			},
			{
				Type:     Scattered,
				Altitude: 5000,
			},
			{
				Type:     Scattered,
				Altitude: 6000,
			},
		},
		Changes: []*Change{
			{
				Type: From,
				Valid: ValidPair{
					From: time.Date(2023, time.August, 21, 22, 0, 0, 0, time.UTC),
				},
				Visibility: Visibility{
					Plus:  true,
					Value: 6,
					Unit:  units.Miles,
				},
				Wind: Wind{
					Direction: Direction{
						Value: 250,
					},
					Speed: 10,
					Unit:  units.Knots,
				},
				SkyCondition: []SkyCondition{
					{
						Type:     Scattered,
						Altitude: 4000,
					},
				},
			},
			{
				Type: From,
				Valid: ValidPair{
					From: time.Date(2023, time.August, 22, 3, 0, 0, 0, time.UTC),
				},
				Visibility: Visibility{
					Plus:  true,
					Value: 6,
					Unit:  units.Miles,
				},
				Wind: Wind{
					Direction: Direction{
						Variable: true,
					},
					Speed: 3,
					Unit:  units.Knots,
				},
				SkyCondition: []SkyCondition{
					{
						Type:     Broken,
						Altitude: 2500,
					},
				},
			},
			{
				Type: From,
				Valid: ValidPair{
					From: time.Date(2023, time.August, 22, 10, 0, 0, 0, time.UTC),
				},
				Visibility: Visibility{
					Plus:  true,
					Value: 6,
					Unit:  units.Miles,
				},
				Wind: Wind{
					Direction: Direction{
						Variable: true,
					},
					Speed: 3,
					Unit:  units.Knots,
				},
				SkyCondition: []SkyCondition{
					{
						Type:     Overcast,
						Altitude: 2500,
					},
				},
			},
			{
				Type: From,
				Valid: ValidPair{
					From: time.Date(2023, time.August, 22, 17, 0, 0, 0, time.UTC),
				},
				Visibility: Visibility{
					Plus:  true,
					Value: 6,
					Unit:  units.Miles,
				},
				Wind: Wind{
					Direction: Direction{
						Value: 260,
					},
					Speed: 6,
					Unit:  units.Knots,
				},
				SkyCondition: []SkyCondition{
					{
						Type:     Broken,
						Altitude: 2500,
					},
				},
			},
			{
				Type: From,
				Valid: ValidPair{
					From: time.Date(2023, time.August, 22, 20, 0, 0, 0, time.UTC),
				},
				Visibility: Visibility{
					Plus:  true,
					Value: 6,
					Unit:  units.Miles,
				},
				Wind: Wind{
					Direction: Direction{
						Value: 260,
					},
					Speed: 12,
					Unit:  units.Knots,
				},
				SkyCondition: []SkyCondition{
					{
						Type:     Scattered,
						Altitude: 3000,
					},
				},
			},
		},
	}

	fc, err := DecodeWithOptions(strings.NewReader(data), Options{
		Month: time.August,
		Year:  2023,
	})
	if err != nil {
		t.Fatalf("Error during parsing: %s", err)
	}

	if diff := deep.Equal(fc, expected); diff != nil {
		t.Error(diff)
	}
}

func TestZGSZ(t *testing.T) {
	const data = `TAF AMD ZGSZ 211907Z 2118/2218 18004MPS 8000 SCT020 TX32/2206Z TN28/2122Z
  TEMPO 2120/2202 SHRA SCT020 FEW023CB
  TEMPO 2204/2208 TSRA SCT020 FEW023CB`

	expected := &Forecast{
		Identifier: "ZGSZ",
		Airport: airports.Airport{
			ICAO:      "ZGSZ",
			IATA:      "SZX",
			Name:      "Shenzhen Bao'an International Airport",
			City:      "Shenzhen",
			State:     "Guangdong",
			Country:   "CN",
			Elevation: 13,
			Latitude:  22.6392993927,
			Longitude: 113.8109970093,
			Timezone:  "Asia/Shanghai",
		},
		PublishTime: time.Date(2023, time.August, 21, 19, 7, 0, 0, time.UTC),
		Valid: ValidPair{
			From:     time.Date(2023, time.August, 21, 18, 0, 0, 0, time.UTC),
			To:       time.Date(2023, time.August, 22, 18, 0, 0, 0, time.UTC),
			Duration: time.Duration(86400000000000),
		},
		Visibility: Visibility{
			Value: 8000,
			Unit:  units.Meters,
		},
		Wind: Wind{
			Direction: Direction{
				Value: 180,
			},
			Speed: 4,
			Unit:  units.MetersPerSecond,
		},
		SkyCondition: []SkyCondition{
			{
				Type:     Scattered,
				Altitude: 2000,
			},
		},
		Temperature: []Temperature{
			{
				Type:  High,
				Value: 32,
				Time:  time.Date(2023, time.August, 22, 6, 0, 0, 0, time.UTC),
			},
			{
				Type:  Low,
				Value: 28,
				Time:  time.Date(2023, time.August, 21, 22, 0, 0, 0, time.UTC),
			},
		},
		Changes: []*Change{
			{
				Type: Temporary,
				Valid: ValidPair{
					From:     time.Date(2023, time.August, 21, 20, 0, 0, 0, time.UTC),
					To:       time.Date(2023, time.August, 22, 2, 0, 0, 0, time.UTC),
					Duration: time.Duration(21600000000000),
				},
				SkyCondition: []SkyCondition{
					{
						Type:     Scattered,
						Altitude: 2000,
					},
					{
						Type:      Few,
						Altitude:  2300,
						CloudType: CumuloNimbus,
					},
				},
				Weather: []Weather{
					{
						Descriptor:    Showers,
						Precipitation: Rain,
					},
				},
			},
			{
				Type: Temporary,
				Valid: ValidPair{
					From:     time.Date(2023, time.August, 22, 4, 0, 0, 0, time.UTC),
					To:       time.Date(2023, time.August, 22, 8, 0, 0, 0, time.UTC),
					Duration: time.Duration(14400000000000),
				},
				SkyCondition: []SkyCondition{
					{
						Type:     Scattered,
						Altitude: 2000,
					},
					{
						Type:      Few,
						Altitude:  2300,
						CloudType: CumuloNimbus,
					},
				},
				Weather: []Weather{
					{
						Descriptor:    Thunderstorm,
						Precipitation: Rain,
					},
				},
			},
		},
	}

	fc, err := DecodeWithOptions(strings.NewReader(data), Options{
		Month: time.August,
		Year:  2023,
	})
	if err != nil {
		t.Fatalf("Error during parsing: %s", err)
	}

	if diff := deep.Equal(fc, expected); diff != nil {
		t.Error(diff)
	}
}

func TestLFBD(t *testing.T) {
	const data = `TAF LFBD 211700Z 2118/2224 31010KT CAVOK TX37/2214Z TN22/2205Z
  BECMG 2118/2120 32004KT
  BECMG 2200/2202 26005KT
  BECMG 2213/2215 32010KT
  BECMG 2222/2224 24004KT`

	expected := &Forecast{
		Identifier: "LFBD",
		Airport: airports.Airport{
			ICAO:      "LFBD",
			IATA:      "BOD",
			Name:      "Bordeaux-Merignac (BA 106) Airport",
			City:      "Bordeaux/Merignac",
			State:     "Nouvelle-Aquitaine",
			Country:   "FR",
			Elevation: 162,
			Latitude:  44.8283004761,
			Longitude: -0.7155560255,
			Timezone:  "Europe/Paris",
		},
		PublishTime: time.Date(2023, time.August, 21, 17, 0, 0, 0, time.UTC),
		Valid: ValidPair{
			From:     time.Date(2023, time.August, 21, 18, 0, 0, 0, time.UTC),
			To:       time.Date(2023, time.August, 23, 0, 0, 0, 0, time.UTC),
			Duration: time.Duration(108000000000000),
		},
		Wind: Wind{
			Direction: Direction{
				Value: 310,
			},
			Speed: 10,
			Unit:  units.Knots,
		},
		Temperature: []Temperature{
			{
				Type:  High,
				Value: 37,
				Time:  time.Date(2023, time.August, 22, 14, 0, 0, 0, time.UTC),
			},
			{
				Type:  Low,
				Value: 22,
				Time:  time.Date(2023, time.August, 22, 5, 0, 0, 0, time.UTC),
			},
		},
		Changes: []*Change{
			{
				Type: Becoming,
				Valid: ValidPair{
					From:     time.Date(2023, time.August, 21, 18, 0, 0, 0, time.UTC),
					To:       time.Date(2023, time.August, 21, 20, 0, 0, 0, time.UTC),
					Duration: time.Duration(7200000000000),
				},
				Wind: Wind{
					Direction: Direction{
						Value: 320,
					},
					Speed: 4,
					Unit:  units.Knots,
				},
			},
			{
				Type: Becoming,
				Valid: ValidPair{
					From:     time.Date(2023, time.August, 22, 0, 0, 0, 0, time.UTC),
					To:       time.Date(2023, time.August, 22, 2, 0, 0, 0, time.UTC),
					Duration: time.Duration(7200000000000),
				},
				Wind: Wind{
					Direction: Direction{
						Value: 260,
					},
					Speed: 5,
					Unit:  units.Knots,
				},
			},
			{
				Type: Becoming,
				Valid: ValidPair{
					From:     time.Date(2023, time.August, 22, 13, 0, 0, 0, time.UTC),
					To:       time.Date(2023, time.August, 22, 15, 0, 0, 0, time.UTC),
					Duration: time.Duration(7200000000000),
				},
				Wind: Wind{
					Direction: Direction{
						Value: 320,
					},
					Speed: 10,
					Unit:  units.Knots,
				},
			},
			{
				Type: Becoming,
				Valid: ValidPair{
					From:     time.Date(2023, time.August, 22, 22, 0, 0, 0, time.UTC),
					To:       time.Date(2023, time.August, 23, 0, 0, 0, 0, time.UTC),
					Duration: time.Duration(7200000000000),
				},
				Wind: Wind{
					Direction: Direction{
						Value: 240,
					},
					Speed: 4,
					Unit:  units.Knots,
				},
			},
		},
		Flags: []Flag{
			CeilingAndVisibilityOK,
		},
	}

	fc, err := DecodeWithOptions(strings.NewReader(data), Options{
		Month: time.August,
		Year:  2023,
	})
	if err != nil {
		t.Fatalf("Error during parsing: %s", err)
	}

	if diff := deep.Equal(fc, expected); diff != nil {
		t.Error(diff)
	}
}

func TestUUEE(t *testing.T) {
	const data = `TAF UUEE 211958Z 2121/2221 VRB01MPS 9999 SCT030 TX20/2212Z TN12/2202Z
  TEMPO 2121/2204 BKN004
  PROB40
  TEMPO 2121/2204 0300 FG
  BECMG 2204/2206 24006MPS
  PROB40
  TEMPO 2209/2218 -TSRA BKN020CB`

	expected := &Forecast{
		Identifier: "UUEE",
		Airport: airports.Airport{
			ICAO:      "UUEE",
			IATA:      "SVO",
			Name:      "Sheremetyevo International Airport",
			City:      "Moscow",
			State:     "Moscow-Oblast",
			Country:   "RU",
			Elevation: 622,
			Latitude:  55.9725990295,
			Longitude: 37.4146003723,
			Timezone:  "Europe/Moscow",
		},
		PublishTime: time.Date(2023, time.August, 21, 19, 58, 0, 0, time.UTC),
		Valid: ValidPair{
			From:     time.Date(2023, time.August, 21, 21, 0, 0, 0, time.UTC),
			To:       time.Date(2023, time.August, 22, 21, 0, 0, 0, time.UTC),
			Duration: time.Duration(86400000000000),
		},
		Visibility: Visibility{
			Value: 9999,
			Unit:  units.Meters,
		},
		Wind: Wind{
			Direction: Direction{
				Variable: true,
			},
			Speed: 1,
			Unit:  units.MetersPerSecond,
		},
		SkyCondition: []SkyCondition{
			{
				Type:     Scattered,
				Altitude: 3000,
			},
		},
		Temperature: []Temperature{
			{
				Type:  High,
				Value: 20,
				Time:  time.Date(2023, time.August, 22, 12, 0, 0, 0, time.UTC),
			},
			{
				Type:  Low,
				Value: 12,
				Time:  time.Date(2023, time.August, 22, 2, 0, 0, 0, time.UTC),
			},
		},
		Changes: []*Change{
			{
				Type: Temporary,
				Valid: ValidPair{
					From:     time.Date(2023, time.August, 21, 21, 0, 0, 0, time.UTC),
					To:       time.Date(2023, time.August, 22, 4, 0, 0, 0, time.UTC),
					Duration: time.Duration(25200000000000),
				},
				SkyCondition: []SkyCondition{
					{
						Type:     Broken,
						Altitude: 400,
					},
				},
			},
			{
				Type: Temporary,
				Valid: ValidPair{
					From:     time.Date(2023, time.August, 21, 21, 0, 0, 0, time.UTC),
					To:       time.Date(2023, time.August, 22, 4, 0, 0, 0, time.UTC),
					Duration: time.Duration(25200000000000),
				},
				Visibility: Visibility{
					Value: 300,
					Unit:  units.Meters,
				},
				Weather: []Weather{
					{
						Obscuration: Fog,
					},
				},
				Probability: 40,
			},
			{
				Type: Becoming,
				Valid: ValidPair{
					From:     time.Date(2023, time.August, 22, 4, 0, 0, 0, time.UTC),
					To:       time.Date(2023, time.August, 22, 6, 0, 0, 0, time.UTC),
					Duration: time.Duration(7200000000000),
				},
				Wind: Wind{
					Direction: Direction{
						Value: 240,
					},
					Speed: 6,
					Unit:  units.MetersPerSecond,
				},
			},
			{
				Type: Temporary,
				Valid: ValidPair{
					From:     time.Date(2023, time.August, 22, 9, 0, 0, 0, time.UTC),
					To:       time.Date(2023, time.August, 22, 18, 0, 0, 0, time.UTC),
					Duration: time.Duration(32400000000000),
				},
				SkyCondition: []SkyCondition{
					{
						Type:      Broken,
						Altitude:  2000,
						CloudType: CumuloNimbus,
					},
				},
				Weather: []Weather{
					{
						Modifier:      Light,
						Descriptor:    Thunderstorm,
						Precipitation: Rain,
					},
				},
				Probability: 40,
			},
		},
	}

	fc, err := DecodeWithOptions(strings.NewReader(data), Options{
		Month: time.August,
		Year:  2023,
	})
	if err != nil {
		t.Fatalf("Error during parsing: %s", err)
	}

	if diff := deep.Equal(fc, expected); diff != nil {
		t.Error(diff)
	}
}

func TestEGLL(t *testing.T) {
	const data = `TAF EGLL 211658Z 2118/2224 22008KT 9999 FEW040
  BECMG 2201/2204 BKN007
  PROB30
  TEMPO 2202/2206 8000 BKN004
  BECMG 2207/2210 SCT025`

	expected := &Forecast{
		Identifier: "EGLL",
		Airport: airports.Airport{
			ICAO:      "EGLL",
			IATA:      "LHR",
			Name:      "London Heathrow Airport",
			City:      "London",
			State:     "England",
			Country:   "GB",
			Elevation: 83,
			Latitude:  51.4706001282,
			Longitude: -0.4619410038,
			Timezone:  "Europe/London",
		},
		PublishTime: time.Date(2023, time.August, 21, 16, 58, 0, 0, time.UTC),
		Valid: ValidPair{
			From:     time.Date(2023, time.August, 21, 18, 0, 0, 0, time.UTC),
			To:       time.Date(2023, time.August, 23, 0, 0, 0, 0, time.UTC),
			Duration: time.Duration(108000000000000),
		},
		Visibility: Visibility{
			Value: 9999,
			Unit:  units.Meters,
		},
		Wind: Wind{
			Direction: Direction{
				Value: 220,
			},
			Speed: 8,
			Unit:  units.Knots,
		},
		SkyCondition: []SkyCondition{
			{
				Type:     Few,
				Altitude: 4000,
			},
		},
		Changes: []*Change{
			{
				Type: Becoming,
				Valid: ValidPair{
					From:     time.Date(2023, time.August, 22, 1, 0, 0, 0, time.UTC),
					To:       time.Date(2023, time.August, 22, 4, 0, 0, 0, time.UTC),
					Duration: time.Duration(10800000000000),
				},
				SkyCondition: []SkyCondition{
					{
						Type:     Broken,
						Altitude: 700,
					},
				},
			},
			{
				Type: Temporary,
				Valid: ValidPair{
					From:     time.Date(2023, time.August, 22, 2, 0, 0, 0, time.UTC),
					To:       time.Date(2023, time.August, 22, 6, 0, 0, 0, time.UTC),
					Duration: time.Duration(14400000000000),
				},
				Visibility: Visibility{
					Value: 8000,
					Unit:  units.Meters,
				},
				SkyCondition: []SkyCondition{
					{
						Type:     Broken,
						Altitude: 400,
					},
				},
				Probability: 30,
			},
			{
				Type: Becoming,
				Valid: ValidPair{
					From:     time.Date(2023, time.August, 22, 7, 0, 0, 0, time.UTC),
					To:       time.Date(2023, time.August, 22, 10, 0, 0, 0, time.UTC),
					Duration: time.Duration(10800000000000),
				},
				SkyCondition: []SkyCondition{
					{
						Type:     Scattered,
						Altitude: 2500,
					},
				},
			},
		},
	}

	fc, err := DecodeWithOptions(strings.NewReader(data), Options{
		Month: time.August,
		Year:  2023,
	})
	if err != nil {
		t.Fatalf("Error during parsing: %s", err)
	}

	if diff := deep.Equal(fc, expected); diff != nil {
		t.Error(diff)
	}
}
