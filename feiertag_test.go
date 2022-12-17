package feiertag

import (
	"testing"
	"time"
)

func TestOstersonntagIn(t *testing.T) {
	if !OstersonntagIn(2022).
		Equal(time.Date(2022, 4, 17, 0, 0, 0, 0, time.UTC)) {
		t.Fail()
	}
}

func TestOstermontagIn(t *testing.T) {
	if !OstermontagIn(2022).
		Equal(time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC)) {
		t.Fail()
	}
}
func TestPfingstsonntagIn(t *testing.T) {
	if !PfingstsonntagIn(2022).
		Equal(time.Date(2022, 6, 5, 0, 0, 0, 0, time.UTC)) {
		t.Fail()
	}
}

func TestPfingstmontagIn(t *testing.T) {
	if !PfingstmontagIn(2022).
		Equal(time.Date(2022, 6, 6, 0, 0, 0, 0, time.UTC)) {
		t.Fail()
	}
}

func TestChristiHimmelfahrtIn(t *testing.T) {
	if !ChristiHimmelfahrtIn(2022).
		Equal(time.Date(2022, 5, 26, 0, 0, 0, 0, time.UTC)) {
		t.Fail()
	}
}

func TestFronleichnamIn(t *testing.T) {
	if !FronleichnamIn(2022).
		Equal(time.Date(2022, 6, 16, 0, 0, 0, 0, time.UTC)) {
		t.Fail()
	}
}

func TestAdvent_1In(t *testing.T) {
	if !Advent_1In(2022).
		Equal(time.Date(2022, 11, 27, 0, 0, 0, 0, time.UTC)) {
		t.Fail()
	}
}

func TestWeiberfastnachtIn(t *testing.T) {
	if !WeiberfastnachtIn(2022).
		Equal(time.Date(2022, 2, 24, 0, 0, 0, 0, time.UTC)) {
		t.Fail()
	}

}

func TestRosenmontagIn(t *testing.T) {
	if !RosenmontagIn(2022).
		Equal(time.Date(2022, 2, 28, 0, 0, 0, 0, time.UTC)) {
		t.Fail()
	}

}
