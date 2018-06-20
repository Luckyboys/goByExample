package lib

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestParseFromTime(t *testing.T) {

	testTime := time.Date(2018, 12, 8, 10, 11, 40, 4000, time.UTC)
	dt := ParseFromTime(&testTime)
	assert.Equal(t, int16(2018), dt.Year)
	assert.Equal(t, int8(12), dt.Month)
	assert.Equal(t, int8(8), dt.Day)
	assert.Equal(t, int8(10), dt.Hour)
	assert.Equal(t, int8(11), dt.Minute)
	assert.Equal(t, int8(40), dt.Second)
	assert.Equal(t, 4000, dt.Nanosecond)
}

func TestDiffTime_Sub(t *testing.T) {

	time1 := time.Date(2018, 10, 10, 0, 0, 0, 0, time.UTC)
	time2 := time.Date(2018, 10, 8, 0, 0, 0, 0, time.UTC)

	dt := ParseFromTime(&time1)
	dt.Sub(&time2)

	assert.Equal(t, int16(0), dt.Year)
	assert.Equal(t, int8(0), dt.Month)
	assert.Equal(t, int8(2), dt.Day)
	assert.Equal(t, int8(0), dt.Hour)
	assert.Equal(t, int8(0), dt.Minute)
	assert.Equal(t, int8(0), dt.Second)
	assert.Equal(t, 0, dt.Nanosecond)
}

func TestDiffTime_Sub2(t *testing.T) {

	time1 := time.Date(2018, 10, 10, 0, 0, 0, 0, time.UTC)
	time2 := time.Date(2018, 9, 8, 0, 0, 0, 0, time.UTC)

	dt := ParseFromTime(&time1)
	dt.Sub(&time2)

	assert.Equal(t, int16(0), dt.Year)
	assert.Equal(t, int8(1), dt.Month)
	assert.Equal(t, int8(2), dt.Day)
	assert.Equal(t, int8(0), dt.Hour)
	assert.Equal(t, int8(0), dt.Minute)
	assert.Equal(t, int8(0), dt.Second)
	assert.Equal(t, 0, dt.Nanosecond)
}

func TestDiffTime_Sub3(t *testing.T) {

	time1 := time.Date(2018, 10, 10, 0, 0, 0, 0, time.UTC)
	time2 := time.Date(2018, 9, 10, 0, 0, 0, 0, time.UTC)

	dt := ParseFromTime(&time1)
	dt.Sub(&time2)

	assert.Equal(t, int16(0), dt.Year)
	assert.Equal(t, int8(1), dt.Month)
	assert.Equal(t, int8(0), dt.Day)
	assert.Equal(t, int8(0), dt.Hour)
	assert.Equal(t, int8(0), dt.Minute)
	assert.Equal(t, int8(0), dt.Second)
	assert.Equal(t, 0, dt.Nanosecond)
}

func TestDiffTime_Sub4(t *testing.T) {

	time1 := time.Date(2018, 10, 10, 0, 0, 0, 0, time.UTC)
	time2 := time.Date(2018, 9, 11, 0, 0, 0, 0, time.UTC)

	dt := ParseFromTime(&time1)
	dt.Sub(&time2)

	assert.Equal(t, int16(0), dt.Year)
	assert.Equal(t, int8(0), dt.Month)
	assert.Equal(t, int8(29), dt.Day)
	assert.Equal(t, int8(0), dt.Hour)
	assert.Equal(t, int8(0), dt.Minute)
	assert.Equal(t, int8(0), dt.Second)
	assert.Equal(t, 0, dt.Nanosecond)
}

func TestDiffTime_Sub5(t *testing.T) {

	time1 := time.Date(2018, 10, 10, 0, 0, 0, 0, time.UTC)
	time2 := time.Date(2017, 11, 10, 0, 0, 0, 0, time.UTC)

	dt := ParseFromTime(&time1)
	dt.Sub(&time2)

	assert.Equal(t, int16(0), dt.Year)
	assert.Equal(t, int8(11), dt.Month)
	assert.Equal(t, int8(0), dt.Day)
	assert.Equal(t, int8(0), dt.Hour)
	assert.Equal(t, int8(0), dt.Minute)
	assert.Equal(t, int8(0), dt.Second)
	assert.Equal(t, 0, dt.Nanosecond)
}

func TestDiffTime_Sub6(t *testing.T) {

	time1 := time.Date(2018, 10, 10, 0, 0, 0, 0, time.UTC)
	time2 := time.Date(2017, 10, 10, 0, 0, 0, 0, time.UTC)

	dt := ParseFromTime(&time1)
	dt.Sub(&time2)

	assert.Equal(t, int16(1), dt.Year)
	assert.Equal(t, int8(0), dt.Month)
	assert.Equal(t, int8(0), dt.Day)
	assert.Equal(t, int8(0), dt.Hour)
	assert.Equal(t, int8(0), dt.Minute)
	assert.Equal(t, int8(0), dt.Second)
	assert.Equal(t, 0, dt.Nanosecond)
}

func TestDiffTime_Sub7(t *testing.T) {

	time1 := time.Date(2018, 10, 10, 0, 0, 0, 0, time.UTC)
	time2 := time.Date(2017, 10, 10, 0, 0, 1, 0, time.UTC)

	dt := ParseFromTime(&time1)
	dt.Sub(&time2)

	assert.Equal(t, int16(0), dt.Year)
	assert.Equal(t, int8(11), dt.Month)
	assert.Equal(t, int8(29), dt.Day)
	assert.Equal(t, int8(23), dt.Hour)
	assert.Equal(t, int8(59), dt.Minute)
	assert.Equal(t, int8(59), dt.Second)
	assert.Equal(t, 0, dt.Nanosecond)
}

func TestDateDiff(t *testing.T) {

	time1 := time.Date(2018, 10, 10, 0, 0, 0, 0, time.UTC)
	time2 := time.Date(2017, 10, 11, 0, 0, 0, 0, time.UTC)

	dt := DateDiff(&time1,&time2)

	assert.Equal(t, int16(0), dt.Year)
	assert.Equal(t, int8(11), dt.Month)
	assert.Equal(t, int8(29), dt.Day)
	assert.Equal(t, int8(0), dt.Hour)
	assert.Equal(t, int8(0), dt.Minute)
	assert.Equal(t, int8(0), dt.Second)
	assert.Equal(t, 0, dt.Nanosecond)
}