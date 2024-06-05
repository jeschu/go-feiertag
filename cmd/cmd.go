package main

import (
	"fmt"
	"github.com/jeschu/go-feiertag"
	"os"
	"strconv"
	"time"
)

func main() {
	argCount := len(os.Args)
	if argCount < 2 {
		usage()
	}
	bundesland := feiertag.ParseBundesland(os.Args[1])
	year := time.Now().Year()
	if argCount == 3 {
		if y, err := strconv.Atoi(os.Args[2]); err != nil {
			panic(err)
		} else {
			year = y
		}
	}
	fmt.Printf("Feiertage %s in %d:", bundesland.String(), year)
	feiertageOnWorkdays := 0
	for _, ftag := range bundesland.Feiertage() {
		inYear := ftag.InYear(year)
		weekday := inYear.Weekday()
		if weekday != time.Sunday && weekday != time.Saturday {
			feiertageOnWorkdays++
		}
		fmt.Printf("\n  %s %s", inYear.Format("Mon 02.01.2006"), ftag.String())
	}
	saturdays := 0
	sundays := 0
	workdays := 0
	holidayOnWeekend := 0
	holidayOnWorkday := 0
	for date := time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local); date.Year() == year; date = date.AddDate(0, 0, 1) {
		weekday := date.Weekday()
		ftag := bundesland.IsFeiertag(date)
		if weekday == time.Sunday {
			sundays++
			if ftag {
				holidayOnWeekend++
			}
		} else if weekday == time.Saturday {
			saturdays++
			if ftag {
				holidayOnWeekend++
			}
		} else if ftag {
			holidayOnWorkday++
		} else {
			workdays++
		}
	}
	fmt.Printf("\n              Arbeitstage: %d\n", workdays)
	fmt.Printf("                 Samstage: %d\n", saturdays)
	fmt.Printf("                 Sonntage: %d\n", sundays)
	fmt.Printf("Feiertage an Arbeitstagen: %d\n", holidayOnWorkday)
	fmt.Printf("   Feiertage an Wochenden: %d\n", holidayOnWeekend)
	fmt.Printf("             Tage im Jahr: %d\n", holidayOnWorkday+workdays+saturdays+sundays)
}

func usage() {
	fmt.Printf("Usage: %s <command> <state> [<year>]", os.Args[0])
	fmt.Printf("\n       <state> one of:")
	for _, bundesland := range feiertag.BundeslandAll {
		fmt.Printf("\n         %s (%s)", bundesland.Code(), bundesland.String())
	}
	fmt.Println()
}
