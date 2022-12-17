// All returned time.Time values have zero time (hour, minute, second, nanosecond) and locale time.UTC.
package feiertag

import (
	"fmt"
	"time"
)

func NeujahrIn(year int) time.Time                  { return Neujahr.InYear(year) }
func HeiligeDreiKoenigeIn(year int) time.Time       { return HeiligeDreiKoenige.InYear(year) }
func InternationalerFrauenTagIn(year int) time.Time { return InternationalerFrauenTag.InYear(year) }
func WeiberfastnachtIn(year int) time.Time          { return Weiberfastnacht.InYear(year) }
func RosenmontagIn(year int) time.Time              { return Rosenmontag.InYear(year) }
func GruendonnerstagIn(year int) time.Time          { return Gruendonnerstag.InYear(year) }
func KarfreitagIn(year int) time.Time               { return Karfreitag.InYear(year) }
func OstersonntagIn(year int) time.Time             { return Ostersonntag.InYear(year) }
func OstermontagIn(year int) time.Time              { return Ostermontag.InYear(year) }
func TagDerArbeitIn(year int) time.Time             { return TagDerArbeit.InYear(year) }
func ChristiHimmelfahrtIn(year int) time.Time       { return ChristiHimmelfahrt.InYear(year) }
func PfingstsonntagIn(year int) time.Time           { return Pfingstsonntag.InYear(year) }
func PfingstmontagIn(year int) time.Time            { return Pfingsmontag.InYear(year) }
func FronleichnamIn(year int) time.Time             { return Fronleichnam.InYear(year) }
func AugsburgerFriedensfestIn(year int) time.Time   { return AugsburgerFriedensfest.InYear(year) }
func MariaeHimmelfahrtIn(year int) time.Time        { return MariaeHimmelfahrt.InYear(year) }
func WeltkindertagIn(year int) time.Time            { return Weltkindertag.InYear(year) }
func TagDerDeutschenEinheitIn(year int) time.Time   { return TagDerDeutschenEinheit.InYear(year) }
func ReformationstagIn(year int) time.Time          { return Reformationstag.InYear(year) }
func AllerheiligenIn(year int) time.Time            { return Allerheiligen.InYear(year) }
func BussUndBettagIn(year int) time.Time            { return BussUndBettag.InYear(year) }
func Advent_1In(year int) time.Time                 { return Advent_1.InYear(year) }
func Advent_2In(year int) time.Time                 { return Advent_2.InYear(year) }
func Advent_3In(year int) time.Time                 { return Advent_3.InYear(year) }
func Advent_4In(year int) time.Time                 { return Advent_4.InYear(year) }
func HeiligAbendIn(year int) time.Time              { return HeiligAbend.InYear(year) }
func Weihnachtsfeiertag_1In(year int) time.Time     { return Weihnachtsfeiertag_1.InYear(year) }
func Weihnachtsfeiertag_2In(year int) time.Time     { return Weihnachtsfeiertag_2.InYear(year) }

var easterCache = make(map[int]time.Time)

func easterByYear(year int) time.Time {
	if easter, ok := easterCache[year]; ok {
		return easter
	} else {
		var offsetDaysFromMarchFirst int
		a := year % 19
		if year >= 1583 {
			b := year / 100
			c := year % 100
			d := b / 4
			e := b % 4
			f := (b + 8) / 25
			g := (b - f + 1) / 3
			h := (19*a + b - d - g + 15) % 30
			i := c / 4
			k := c % 4
			l := (32 + 2*e + 2*i - h - k) % 7
			m := (a + 11*h + 22*l) / 451
			offsetDaysFromMarchFirst = 22 + h + l - 7*m
		} else {
			b := year % 7
			c := year % 4
			d := (19*a + 15) % 30
			e := (2*c + 4*b - d + 34) % 7
			offsetDaysFromMarchFirst = 22 + d + e
		}
		easter = time.Date(year, time.March, offsetDaysFromMarchFirst, 0, 0, 0, 0, time.UTC)
		easterCache[year] = easter
		return easter
	}
}

func easterByYearOffset(year, offset int) time.Time { return easterByYear(year).AddDate(0, 0, offset) }

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

type Feiertag uint8

const (
	Neujahr Feiertag = iota
	HeiligeDreiKoenige
	InternationalerFrauenTag
	Weiberfastnacht
	Rosenmontag
	Gruendonnerstag
	Karfreitag
	Ostersonntag
	Ostermontag
	TagDerArbeit
	ChristiHimmelfahrt
	Pfingstsonntag
	Pfingsmontag
	Fronleichnam
	AugsburgerFriedensfest
	MariaeHimmelfahrt
	Weltkindertag
	TagDerDeutschenEinheit
	Reformationstag
	Allerheiligen
	BussUndBettag
	Advent_1
	Advent_2
	Advent_3
	Advent_4
	HeiligAbend
	Weihnachtsfeiertag_1
	Weihnachtsfeiertag_2
	Sylvester
)

var fbw = map[Bundesland][]Feiertag{
	BB: []Feiertag{Neujahr, Karfreitag, Ostersonntag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingstsonntag, Pfingsmontag, TagDerDeutschenEinheit, Reformationstag, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	BE: []Feiertag{Neujahr, InternationalerFrauenTag, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, TagDerDeutschenEinheit, Reformationstag, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	BW: []Feiertag{Neujahr, HeiligeDreiKoenige, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, Fronleichnam, TagDerDeutschenEinheit, Allerheiligen, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	BY: []Feiertag{Neujahr, HeiligeDreiKoenige, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, Fronleichnam, TagDerDeutschenEinheit, Allerheiligen, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	HB: []Feiertag{Neujahr, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, TagDerDeutschenEinheit, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	HE: []Feiertag{Neujahr, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, Fronleichnam, TagDerDeutschenEinheit, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	HH: []Feiertag{Neujahr, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, TagDerDeutschenEinheit, Reformationstag, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	MV: []Feiertag{Neujahr, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, TagDerDeutschenEinheit, Reformationstag, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	NI: []Feiertag{Neujahr, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, TagDerDeutschenEinheit, Reformationstag, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	NW: []Feiertag{Neujahr, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, Fronleichnam, TagDerDeutschenEinheit, Allerheiligen, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	RP: []Feiertag{Neujahr, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, Fronleichnam, TagDerDeutschenEinheit, Allerheiligen, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	SH: []Feiertag{Neujahr, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, TagDerDeutschenEinheit, Reformationstag, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	SL: []Feiertag{Neujahr, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, Fronleichnam, MariaeHimmelfahrt, TagDerDeutschenEinheit, Allerheiligen, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	SN: []Feiertag{Neujahr, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, TagDerDeutschenEinheit, Reformationstag, BussUndBettag, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	ST: []Feiertag{Neujahr, HeiligeDreiKoenige, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, TagDerDeutschenEinheit, Reformationstag, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
	TH: []Feiertag{Neujahr, Karfreitag, Ostermontag, TagDerArbeit, ChristiHimmelfahrt, Pfingsmontag, Weltkindertag, TagDerDeutschenEinheit, Reformationstag, Weihnachtsfeiertag_1, Weihnachtsfeiertag_2},
}

func (feiertag Feiertag) feiertagIn(bundesland Bundesland) bool {
	if feiertage, ok := fbw[bundesland]; !ok {
		return false
	} else {
		for _, f := range feiertage {
			if f == feiertag {
				return true
			}
		}
		return false
	}
}

func (feiertag Feiertag) InYear(year int) time.Time {
	switch feiertag {
	case Neujahr:
		return date(year, 1, 1)
	case HeiligeDreiKoenige:
		return date(year, 1, 6)
	case InternationalerFrauenTag:
		return date(year, 3, 8)
	case Weiberfastnacht:
		return easterByYearOffset(2022, -52)
	case Rosenmontag:
		return easterByYearOffset(2022, -48)
	case Gruendonnerstag:
		return easterByYearOffset(year, -3)
	case Karfreitag:
		return easterByYearOffset(year, -2)
	case Ostersonntag:
		return easterByYear(year)
	case Ostermontag:
		return easterByYearOffset(year, 1)
	case TagDerArbeit:
		return date(year, 5, 1)
	case ChristiHimmelfahrt:
		return easterByYearOffset(year, 39)
	case Pfingstsonntag:
		return easterByYearOffset(year, 49)
	case Pfingsmontag:
		return easterByYearOffset(year, 50)
	case Fronleichnam:
		return easterByYearOffset(year, 60)
	case AugsburgerFriedensfest:
		return date(year, 8, 8)
	case MariaeHimmelfahrt:
		return date(year, 8, 15)
	case Weltkindertag:
		return date(year, 9, 20)
	case TagDerDeutschenEinheit:
		return date(year, 10, 3)
	case Reformationstag:
		return date(year, 10, 31)
	case Allerheiligen:
		return date(year, 11, 1)
	case BussUndBettag:
		return Advent_1.InYear(year).AddDate(0, 0, -11)
	case Advent_1:
		d := time.Date(year, 11, 27, 0, 0, 0, 0, time.UTC)
		for d.Weekday() != time.Sunday {
			d.AddDate(0, 0, 1)
		}
		return d
	case Advent_2:
		return Advent_1.InYear(year).AddDate(0, 0, 7)
	case Advent_3:
		return Advent_1.InYear(year).AddDate(0, 0, 14)
	case Advent_4:
		return Advent_1.InYear(year).AddDate(0, 0, 21)
	case HeiligAbend:
		return date(year, 12, 24)
	case Weihnachtsfeiertag_1:
		return date(year, 12, 25)
	case Weihnachtsfeiertag_2:
		return date(year, 12, 26)
	case Sylvester:
		return date(year, 12, 31)
	default:
		panic(fmt.Sprintf("unknown holiday (%v)", feiertag))
	}
}

func (bundesland Bundesland) Feiertage() []Feiertag { return fbw[bundesland] }
