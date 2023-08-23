package taf

func convertSkyConditionType(s string) SkyConditionType {
	switch s {
	case "FEW":
		return Few
	case "SCT":
		return Scattered
	case "BKN":
		return Broken
	case "OVC":
		return Overcast
	case "VV":
		return VerticalVisibility
	case "SKC":
		return SkyClear
	default:
		return ""
	}
}

func convertCloudType(s string) CloudType {
	switch s {
	case "CB":
		return CumuloNimbus
	case "TCU":
		return ToweringCumulus
	default:
		return ""
	}
}

func convertModifier(s string) Modifier {
	switch s {
	case "+":
		return Heavy
	case "-":
		return Light
	default:
		return ""
	}
}

func convertDescriptor(s string) Descriptor {
	switch s {
	case "MI":
		return Shallow
	case "BC":
		return Patches
	case "DC":
		return LowDrifting
	case "BL":
		return Blowing
	case "SH":
		return Showers
	case "TS":
		return Thunderstorm
	case "FZ":
		return Freezing
	case "PR":
		return Partial
	default:
		return ""
	}
}

func convertPrecipitation(s string) Precipitation {
	switch s {
	case "DZ":
		return Drizzle
	case "RA":
		return Rain
	case "SN":
		return Snow
	case "SG":
		return SnowGrains
	case "IC":
		return IceCrystals
	case "PL":
		return IcePellets
	case "GR":
		return Hail
	case "GS":
		return SmallHail
	case "UP":
		return Unknown
	default:
		return ""
	}
}

func convertObscuration(s string) Obscuration {
	switch s {
	case "BR":
		return Mist
	case "FG":
		return Fog
	case "FU":
		return Smoke
	case "DU":
		return Dust
	case "SA":
		return Sand
	case "HZ":
		return Haze
	case "PY":
		return Spray
	case "VA":
		return VolcanicAsh
	default:
		return ""
	}
}

func convertPhenomenon(s string) Phenomenon {
	switch s {
	case "PO":
		return Whirls
	case "SQ":
		return Squalls
	case "FC":
		return FunnelCloud
	case "SS":
		return Sandstorm
	case "DS":
		return Duststorm
	default:
		return ""
	}
}

func convertTemperatureType(s string) TemperatureType {
	switch s {
	case "TX":
		return High
	case "TN":
		return Low
	default:
		return ""
	}
}

func convertChangeType(s string) ChangeType {
	switch s {
	case "FM":
		return From
	case "BECMG":
		return Becoming
	case "TEMPO":
		return Temporary
	default:
		return ""
	}
}
