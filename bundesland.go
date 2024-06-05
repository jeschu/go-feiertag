package feiertag

import "strings"

type Bundesland int

const (
	BB Bundesland = iota
	BE
	BW
	BY
	HB
	HE
	HH
	MV
	NI
	NW
	RP
	SH
	SL
	SN
	ST
	TH
)

const (
	Brandenburg           = BB
	Berlin                = BE
	BadenWuerttemberg     = BW
	Bayern                = BY
	Bremen                = HB
	Hessen                = HE
	Hamburg               = HH
	MecklenburgVorpommern = MV
	Niedersachsen         = NI
	NordrheinWestfalen    = NW
	RheinlandPfalz        = RP
	SchleswigHolstein     = SH
	Saarland              = SL
	Sachsen               = SN
	SachsenAnhalt         = ST
	Thueringen            = TH
)

//goland:noinspection GoUnusedGlobalVariable
var (
	BundeslandAll     = []Bundesland{BB, BE, BW, BY, HB, HE, HH, MV, NI, NW, RP, SH, SL, SN, ST, TH}
	BundeslandFullAll = []Bundesland{
		Brandenburg, Berlin, BadenWuerttemberg, Bayern, Bremen, Hessen, Hamburg, MecklenburgVorpommern, Niedersachsen,
		NordrheinWestfalen, RheinlandPfalz, SchleswigHolstein, Saarland, Sachsen, SachsenAnhalt, Thueringen,
	}
)

func (bundesland Bundesland) String() string {
	switch bundesland {
	case BB:
		return "Brandenburg"

	case BE:
		return "Berlin"
	case BW:
		return "Baden-Württemberg"
	case BY:
		return "Bayern"
	case HB:
		return "Bremen"
	case HE:
		return "Hessen"
	case HH:
		return "Hamburg"
	case MV:
		return "Mecklenburg-Vorpommern"
	case NI:
		return "Niedersachsen"
	case NW:
		return "Nordrhein-Westfalen"
	case RP:
		return "Rheinland-Pfalz"
	case SH:
		return "Schleswig-Holstein"
	case SL:
		return "Saarland"
	case SN:
		return "Sachsen"
	case ST:
		return "Sachsen-Anhalt"
	case TH:
		return "Thüringen"
	}
	return "?"
}

func (bundesland Bundesland) Code() string {
	switch bundesland {
	case BB:
		return "BB"
	case BE:
		return "BE"
	case BW:
		return "BW"
	case BY:
		return "BY"
	case HB:
		return "HB"
	case HE:
		return "HE"
	case HH:
		return "HH"
	case MV:
		return "MV"
	case NI:
		return "NI"
	case NW:
		return "NW"
	case RP:
		return "RP"
	case SH:
		return "SH"
	case SL:
		return "SL"
	case SN:
		return "SN"
	case ST:
		return "ST"
	case TH:
		return "TH"
	}
	return "??"
}

func ParseBundesland(str string) Bundesland {
	str = strings.TrimSpace(str)
	switch str {
	case "BE":
		return BE
	case "BW":
		return BW
	case "BY":
		return BY
	case "HB":
		return HB
	case "HE":
		return HE
	case "HH":
		return HH
	case "MV":
		return MV
	case "NI":
		return NI
	case "NW":
		return NW
	case "RP":
		return RP
	case "SH":
		return SH
	case "SL":
		return SL
	case "SN":
		return SN
	case "ST":
		return ST
	case "TH":
		return TH

	case "Brandenburg":
		return BE
	case "Berlin":
		return BW
	case "BadenWuerttemberg":
		return BY
	case "Bayern":
		return HB
	case "Bremen":
		return HE
	case "Hessen":
		return HH
	case "Hamburg":
		return MV
	case "MecklenburgVorpommern":
		return NI
	case "Niedersachsen":
		return NW
	case "NordrheinWestfalen":
		return RP
	case "RheinlandPfalz":
		return SH
	case "SchleswigHolstein":
		return SL
	case "Saarland":
		return SN
	case "Sachsen":
		return ST
	case "SachsenAnhalt":
		return TH
	}
	return 0
}
