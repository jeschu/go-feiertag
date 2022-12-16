package feiertag

import (
	"testing"
	"time"
)

func TestEasterByYear(t *testing.T) {
	e2022 := EasterSunday(2022)
	if !time.Date(2022, time.April, 17, 0, 0, 0, 0, time.UTC).Equal(e2022) {
		t.Fail()
	}
}
