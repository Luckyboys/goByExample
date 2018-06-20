package lib

import (
	"time"
)

type DiffTime struct {
	Year int16
	Month int8
	Day int8
	Hour int8
	Minute int8
	Second int8
	Nanosecond int
}

func ParseFromTime(t *time.Time) (diff *DiffTime) {

	diff = &DiffTime{
		Year:       int16(t.Year()),
		Month:      int8(t.Month()),
		Day:        int8(t.Day()),
		Hour:       int8(t.Hour()),
		Minute:     int8(t.Minute()),
		Second:     int8(t.Second()),
		Nanosecond: int(t.Nanosecond()),
	}

	return
}

func (diff *DiffTime) Sub(time2 *time.Time) {

	diff.Nanosecond -= time2.Nanosecond()
	if diff.Nanosecond < 0 {
		diff.Second -=1
		diff.Nanosecond += int(time.Second)
	}

	diff.Second -= int8(time2.Second())
	if diff.Second < 0 {
		diff.Minute -= 1
		diff.Second += 60
	}

	diff.Minute -= int8(time2.Minute())
	if diff.Minute < 0 {
		diff.Hour -= 1
		diff.Minute += 60
	}

	diff.Hour -= int8(time2.Hour())

	theTime := time.Date(int(diff.Year), time.Month(diff.Month), int(diff.Day), 0, 0, 0, 0, time.UTC)

	if diff.Hour < 0 {
		theTime = theTime.AddDate(0, 0, -1)
		diff.Hour += 24
	}

	if time2.Day() == theTime.Day() {

		diff.Day = 0
	} else {
		theTime = theTime.AddDate(0, 0, -time2.Day())
		diff.Day = int8(theTime.Day())
	}

	diff.Month = int8(theTime.Month() - time2.Month())
	if diff.Month < 0 {
		diff.Year -= 1
		diff.Month += 12
	}

	diff.Year = diff.Year - int16(time2.Year())

	return
}

func DateDiff(time1, time2 *time.Time) (returnValue *DiffTime) {

	returnValue = ParseFromTime(time1)
	returnValue.Sub(time2)

	return
}