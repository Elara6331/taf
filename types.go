package taf

import (
	"time"

	"go.elara.ws/taf/airports"
	"go.elara.ws/taf/units"
)

// Forecast represents a Terminal Aerodrome Forecast (TAF) weather report for a specific airport.
type Forecast struct {
	// Identifier holds the ICAO airport identifier for which this forecast was issued.
	Identifier string `json:"identifier,omitempty"`

	// Airport provides additional information about the airport for which this forecast was issued.
	Airport airports.Airport `json:"airport,omitempty"`

	// PublishTime indicates the time at which this forecast was issued.
	PublishTime time.Time `json:"publish_time,omitempty"`

	// Valid defines the period during which this forecast is applicable.
	Valid ValidPair `json:"valid,omitempty"`

	// Visibility describes the anticipated visibility conditions.
	Visibility Visibility `json:"visibility,omitempty"`

	// Wind describes the projected wind conditions.
	Wind Wind `json:"wind,omitempty"`

	// SkyCondition lists the expected sky conditions.
	SkyCondition []SkyCondition `json:"sky_condition,omitempty"`

	// Temperature lists the expected temperature values.
	Temperature []Temperature `json:"temperature,omitempty"`

	// Weather lists information about the expected weather conditions.
	Weather []Weather `json:"weather,omitempty"`

	// Probabilities contains the probabilities for potential conditions.
	Probabilities []*Probability `json:"probabilities,omitempty"`

	// Changes lists any expected changes in conditions.
	Changes []*Change `json:"changes,omitempty"`

	// Flags contains special flags associated with the forecast.
	Flags []Flag `json:"flags,omitempty"`
}

// Change represents a change in weather conditions within a forecast.
type Change struct {
	// Type specifies the nature of this weather change.
	Type ChangeType `json:"type,omitempty"`

	// Valid defines the period during which this change is applicable.
	Valid ValidPair `json:"valid,omitempty"`

	// Visibility describes the anticipated visibility conditions.
	Visibility Visibility `json:"visibility,omitempty"`

	// Wind describes the projected wind conditions.
	Wind Wind `json:"wind,omitempty"`

	// SkyCondition lists the expected sky conditions.
	SkyCondition []SkyCondition `json:"sky_condition,omitempty"`

	// Temperature lists the expected temperature values.
	Temperature []Temperature `json:"temperature,omitempty"`

	// Weather lists information about the expected weather conditions.
	Weather []Weather `json:"weather,omitempty"`

	// Flags contains special flags associated with the change.
	Flags []Flag `json:"flags,omitempty"`

	// Probability indicates the percent chance of this change occurring.
	Probability int `json:"probability,omitempty"`
}

// Probability represents the probability of potential conditions occurring within a forecast.
type Probability struct {
	// Valid defines the period during which these potential conditions are applicable.
	Valid ValidPair `json:"valid,omitempty"`

	// Value indicates the percent chance of these conditions occurring.
	Value int `json:"value,omitempty"`

	// Visibility describes the anticipated visibility conditions.
	Visibility Visibility `json:"visibility,omitempty"`

	// Wind describes the projected wind conditions.
	Wind Wind `json:"wind,omitempty"`

	// SkyCondition lists the expected sky conditions.
	SkyCondition []SkyCondition `json:"sky_condition,omitempty"`

	// Temperature lists the expected temperature values.
	Temperature []Temperature `json:"temperature,omitempty"`

	// Weather lists information about the expected weather conditions.
	Weather []Weather `json:"weather,omitempty"`

	// Flags contains special flags associated with the potential conditions.
	Flags []Flag `json:"flags,omitempty"`
}

// ValidPair represents a time interval for which weather data is valid.
type ValidPair struct {
	// From represents the time from which the data is valid.
	From time.Time `json:"from,omitempty"`

	// To indicates the time until which the data is valid.
	To time.Time `json:"to,omitempty"`

	// Duration contains the total duration for which the data remains valid
	Duration time.Duration `json:"duration,omitempty"`
}

// Visibility represents the visibility conditions in the forecast.
type Visibility struct {
	// Plus indicates whether visibility is expected to be greater than the specified value.
	Plus bool `json:"plus,omitempty"`

	// Value holds the visibility measurement. Its unit is determined by the Unit field.
	Value float64 `json:"value,omitempty"`

	// Unit specifies the unit of measurement for the visibility value.
	Unit units.Distance `json:"unit,omitempty"`
}

// SkyConditionType represents different types of sky conditions in the forecast.
type SkyConditionType string

// Sky Condition Types
const (
	Few                SkyConditionType = "Few"
	Scattered          SkyConditionType = "Scattered"
	Broken             SkyConditionType = "Broken"
	Overcast           SkyConditionType = "Overcast"
	VerticalVisibility SkyConditionType = "VerticalVisibility"
	SkyClear           SkyConditionType = "SkyClear"
)

// CloudType represents different types of cloud formations.
type CloudType string

// Cloud Types
const (
	CumuloNimbus    CloudType = "CumuloNumbus"
	ToweringCumulus CloudType = "ToweringCumulus"
)

// SkyCondition represents the condition of the sky, including cloud cover and altitude.
type SkyCondition struct {
	// Type specifies the nature of the expected sky condition.
	Type SkyConditionType `json:"type,omitempty"`

	// Altitude represents the altitude at which this sky condition is anticipated, in feet.
	Altitude int `json:"altitude,omitempty"`

	// CloudType defines the type of clouds expected in the sky.
	CloudType CloudType `json:"cloud_type,omitempty"`
}

// Wind represents wind-related information in a weather forecast.
type Wind struct {
	// Direction indicates the wind direction of the expected wind.
	Direction Direction `json:"direction,omitempty"`

	// WindShear specifies the altitude at which wind shear is expected.
	WindShear int `json:"wind_shear,omitempty"`

	// Speed represents the anticipated wind speed. The unit is determined by the Unit field.
	Speed int `json:"speed,omitempty"`

	// Gusts holds the projected gust speed. The unit is determined by the Unit field.
	Gusts int `json:"gusts,omitempty"`

	// Unit denotes the unit of measurement for wind and gust speeds.
	Unit units.Speed `json:"unit,omitempty"`
}

// Direction describes the wind direction, which can be variable.
type Direction struct {
	// Variable signifies if the wind direction is variable. When true, Value is set to zero.
	Variable bool `json:"variable,omitempty"`

	// Value specifies the wind direction in degrees.
	Value int `json:"value,omitempty"`
}

// Modifier represents modifiers for weather conditions, such as "Heavy" or "Light".
type Modifier string

// Weather Modifiers
const (
	// Heavy indicates that weather conditions are expected to be severe.
	Heavy Modifier = "Heavy"
	// Light indicates that weather conditions are expected to be mild.
	Light Modifier = "Light"
)

// Descriptor represents descriptors for weather conditions, such as "Shallow" or "Showers".
type Descriptor string

// Weather Descriptors
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

// Precipitation represents different types of precipitation.
type Precipitation string

// Precipitation Types
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

// Obscuration represents different types of atmospheric obscurations.
type Obscuration string

// Atmospheric Obscuration Types
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

// Phenomenon represents different atmospheric phenomena like whirls, squalls, etc.
type Phenomenon string

// Weather Phenomena
const (
	Whirls      Phenomenon = "Whirls"
	Squalls     Phenomenon = "Squalls"
	FunnelCloud Phenomenon = "FunnelCloud"
	Sandstorm   Phenomenon = "Sandstorm"
	Duststorm   Phenomenon = "Duststorm"
)

// Weather represents various weather-related conditions in the forecast.
type Weather struct {
	// Vicinity specifies if the described weather is occurring near the airport.
	Vicinity bool `json:"vicinity,omitempty"`

	// Modifier indicates the severity of the weather conditions.
	Modifier Modifier `json:"modifier,omitempty"`

	// Descriptor provides details about the specific type of expected weather.
	Descriptor Descriptor `json:"descriptor,omitempty"`

	// Precipitation indicates the anticipated type of precipitation.
	Precipitation Precipitation `json:"precipitation,omitempty"`

	// Obscuration describes any potential atmospheric obscurations expected.
	Obscuration Obscuration `json:"obscuration,omitempty"`

	// Phenomenon contains anticipated weather phenomena.
	Phenomenon Phenomenon `json:"phenomenon,omitempty"`
}

// TemperatureType represents different types of temperature data, like "High" or "Low".
type TemperatureType string

// Temperature Types
const (
	High TemperatureType = "High"
	Low  TemperatureType = "Low"
)

// Temperature represents temperature-related details in the forecast.
type Temperature struct {
	// Type specifies if this temperature is a high or low value.
	Type TemperatureType `json:"type,omitempty"`

	// Value holds the anticipated temperature in degrees Celsius.
	Value int `json:"value,omitempty"`

	// Time indicates the expected time for this temperature.
	Time time.Time `json:"time,omitempty"`
}

// ChangeType represents different types of changes in weather conditions.
type ChangeType string

// Change Types
const (
	// From indicates a rapid change.
	From ChangeType = "From"
	// Becoming indicates a slow or gradual change.
	Becoming ChangeType = "Becoming"
	// Temporary indicates that a change is expected to last generally less than an hour.
	Temporary ChangeType = "Temporary"
)

// Flag represents special flags for specific weather conditions.
type Flag string

// Flags
const (
	// CeilingAndVisibilityOK indicates that visibility is over 10km, that
	// there are no significant clouds, and no significant weather
	CeilingAndVisibilityOK Flag = "CeilingAndVisibilityOK"
)
