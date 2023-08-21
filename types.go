package taf

import (
	"time"

	"go.elara.ws/taf/units"
)

const None = ""

type Forecast struct {
	Identifier    string         `json:"identifier,omitempty"`
	PublishTime   time.Time      `json:"publish_time,omitempty"`
	Valid         ValidPair      `json:"valid,omitempty"`
	Visibility    Visibility     `json:"visibility,omitempty"`
	Wind          Wind           `json:"wind,omitempty"`
	SkyCondition  []SkyCondition `json:"sky_condition,omitempty"`
	Temperature   []Temperature  `json:"temperature,omitempty"`
	Weather       []Weather      `json:"weather,omitempty"`
	Probabilities []*Probability `json:"probabilities,omitempty"`
	Changes       []*Change      `json:"changes,omitempty"`
	Flags         []Flag         `json:"flags,omitempty"`
}

type Change struct {
	Type         ChangeType     `json:"type,omitempty"`
	Valid        ValidPair      `json:"valid,omitempty"`
	Visibility   Visibility     `json:"visibility,omitempty"`
	Wind         Wind           `json:"wind,omitempty"`
	SkyCondition []SkyCondition `json:"sky_condition,omitempty"`
	Temperature  []Temperature  `json:"temperature,omitempty"`
	Weather      []Weather      `json:"weather,omitempty"`
	Flags        []Flag         `json:"flags,omitempty"`
	Probability  int            `json:"probability,omitempty"`
}

type Probability struct {
	Valid        ValidPair      `json:"valid,omitempty"`
	Value        int            `json:"value,omitempty"`
	Visibility   Visibility     `json:"visibility,omitempty"`
	Wind         Wind           `json:"wind,omitempty"`
	SkyCondition []SkyCondition `json:"sky_condition,omitempty"`
	Temperature  []Temperature  `json:"temperature,omitempty"`
	Weather      []Weather      `json:"weather,omitempty"`
	Flags        []Flag         `json:"flags,omitempty"`
}

type ValidPair struct {
	From     time.Time     `json:"from,omitempty"`
	To       time.Time     `json:"to,omitempty"`
	Duration time.Duration `json:"duration,omitempty"`
}

type Visibility struct {
	Plus  bool           `json:"plus,omitempty"`
	Value float64        `json:"value,omitempty"`
	Unit  units.Distance `json:"unit,omitempty"`
}

type SkyConditionType string

const (
	Few                SkyConditionType = "FEW"
	Scattered          SkyConditionType = "SCT"
	Broken             SkyConditionType = "BKN"
	Overcast           SkyConditionType = "OVC"
	VerticalVisibility SkyConditionType = "VV"
	SkyClear           SkyConditionType = "SKC"
)

type CloudType string

const (
	CumuloNimbus    CloudType = "CB"
	ToweringCumulus CloudType = "TCU"
)

type SkyCondition struct {
	Type      SkyConditionType `json:"type,omitempty"`
	Altitude  int              `json:"altitude,omitempty"`
	CloudType CloudType        `json:"cloud_type,omitempty"`
}

type Wind struct {
	Direction Direction   `json:"direction,omitempty"`
	WindShear int         `json:"wind_shear,omitempty"`
	Speed     int         `json:"speed,omitempty"`
	Gusts     int         `json:"gusts,omitempty"`
	Unit      units.Speed `json:"unit,omitempty"`
}

type Direction struct {
	Variable bool `json:"variable,omitempty"`
	Value    int  `json:"value,omitempty"`
}

type Modifier string

const (
	Heavy Modifier = "+"
	Light Modifier = "-"
)

type Descriptor string

const (
	Shallow      Descriptor = "MI"
	Patches      Descriptor = "BC"
	LowDrifting  Descriptor = "DC"
	Blowing      Descriptor = "BL"
	Showers      Descriptor = "SH"
	Thunderstorm Descriptor = "TS"
	Freezing     Descriptor = "FZ"
	Partial      Descriptor = "PR"
)

type Precipitation string

const (
	Drizzle     Precipitation = "DZ"
	Rain        Precipitation = "RA"
	Snow        Precipitation = "SN"
	SnowGrains  Precipitation = "SG"
	IceCrystals Precipitation = "IC"
	IcePellets  Precipitation = "PL"
	Hail        Precipitation = "GR"
	SmallHail   Precipitation = "GS"
	Unknown     Precipitation = "UP"
)

type Obscuration string

const (
	Mist        Obscuration = "BR"
	Fog         Obscuration = "FG"
	Smoke       Obscuration = "FU"
	Dust        Obscuration = "DU"
	Sand        Obscuration = "SA"
	Haze        Obscuration = "HZ"
	Spray       Obscuration = "PY"
	VolcanicAsh Obscuration = "VA"
)

type Phenomenon string

const (
	Whirls      Phenomenon = "PO"
	Squalls     Phenomenon = "SQ"
	FunnelCloud Phenomenon = "FC"
	SandStorm   Phenomenon = "SS"
	DustStorm   Phenomenon = "DS"
)

type Weather struct {
	Vicinity      bool          `json:"vicinity,omitempty"`
	Modifier      Modifier      `json:"modifier,omitempty"`
	Descriptor    Descriptor    `json:"descriptor,omitempty"`
	Precipitation Precipitation `json:"precipitation,omitempty"`
	Obscuration   Obscuration   `json:"obscuration,omitempty"`
	Phenomenon    Phenomenon    `json:"phenomenon,omitempty"`
}

type TemperatureType string

const (
	High TemperatureType = "TX"
	Low  TemperatureType = "TN"
)

type Temperature struct {
	Type  TemperatureType `json:"type,omitempty"`
	Value int             `json:"value,omitempty"`
	Time  time.Time       `json:"time,omitempty"`
}

type ChangeType string

const (
	From      ChangeType = "FM"
	Becoming  ChangeType = "BECMG"
	Temporary ChangeType = "TEMPO"
)

type Flag string

const (
	CeilingAndVisibilityOK Flag = "CAVOK"
)
