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
	Few                SkyConditionType = "Few"
	Scattered          SkyConditionType = "Scattered"
	Broken             SkyConditionType = "Broken"
	Overcast           SkyConditionType = "Overcast"
	VerticalVisibility SkyConditionType = "VerticalVisibility"
	SkyClear           SkyConditionType = "SkyClear"
)

type CloudType string

const (
	CumuloNimbus    CloudType = "CumuloNumbus"
	ToweringCumulus CloudType = "ToweringCumulus"
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
	Heavy Modifier = "Heavy"
	Light Modifier = "Light"
)

type Descriptor string

const (
	Shallow      Descriptor = "Shallow"
	Patches      Descriptor = "Patches"
	LowDrifting  Descriptor = "LowDrifting"
	Blowing      Descriptor = "Blowing"
	Showers      Descriptor = "Showers"
	Thunderstorm Descriptor = "Thunderstorm"
	Freezing     Descriptor = "Freezing"
	Partial      Descriptor = "Partial"
)

type Precipitation string

const (
	Drizzle     Precipitation = "Drizzle"
	Rain        Precipitation = "Rain"
	Snow        Precipitation = "Snow"
	SnowGrains  Precipitation = "SnowGrains"
	IceCrystals Precipitation = "IceCrystals"
	IcePellets  Precipitation = "IcePellets"
	Hail        Precipitation = "Hail"
	SmallHail   Precipitation = "SmallHail"
	Unknown     Precipitation = "Unknown"
)

type Obscuration string

const (
	Mist        Obscuration = "Mist"
	Fog         Obscuration = "Fog"
	Smoke       Obscuration = "Smoke"
	Dust        Obscuration = "Dust"
	Sand        Obscuration = "Sand"
	Haze        Obscuration = "Haze"
	Spray       Obscuration = "Spray"
	VolcanicAsh Obscuration = "VolcanicAsh"
)

type Phenomenon string

const (
	Whirls      Phenomenon = "Whirls"
	Squalls     Phenomenon = "Squalls"
	FunnelCloud Phenomenon = "FunnelCloud"
	Sandstorm   Phenomenon = "Sandstorm"
	Duststorm   Phenomenon = "Duststorm"
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
	High TemperatureType = "High"
	Low  TemperatureType = "Low"
)

type Temperature struct {
	Type  TemperatureType `json:"type,omitempty"`
	Value int             `json:"value,omitempty"`
	Time  time.Time       `json:"time,omitempty"`
}

type ChangeType string

const (
	From      ChangeType = "From"
	Becoming  ChangeType = "Becoming"
	Temporary ChangeType = "Temporary"
)

type Flag string

const (
	CeilingAndVisibilityOK Flag = "CeilingAndVisibilityOK"
)
