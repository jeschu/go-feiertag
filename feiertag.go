package feiertag

import "time"

func EasterSunday(year int) time.Time {
	return easterByYear(year)
}

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
