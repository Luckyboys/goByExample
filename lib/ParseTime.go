package lib

import (
	"log"
	"time"
)

const timeFormat1 = "2006-01-02T15:04:05.0000000Z"
const timeFormat2 = "2006-01-02T15:04:05Z"
const timeFormat3 = "2006-01-02 15:04:05 +00:00 UTC"
const timeFormat4 = "2006-01-02 15:04:05 +00:00"
const timeFormat5 = "2006-01-02T15:04:05.0000000Z"
const timeFormatNonTimeZone1 = "2006-01-02"
const timeFormatNonTimeZone2 = "2006-01-02 15:04:05"
const timeFormatNonTimeZone3 = "20060102150405"

var (
	timeFormats = []string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		timeFormat1,
		timeFormat2,
		timeFormat3,
		timeFormat4,
		timeFormat5,
	}

	timeFormatsNonTimeZone = []string{
		timeFormatNonTimeZone1,
		timeFormatNonTimeZone2,
		timeFormatNonTimeZone3,
	}
)

func ParseTime(timeData interface{}) (int64, bool) {
	var timeStamp int64
	switch timeData.(type) {
	case time.Time:
		updateTime, ok := timeData.(time.Time)
		if !ok {
			log.Println("Convert UpdateTime(time.Time) Failure")
			return 0, false
		}
		timeStamp = updateTime.Unix()

	case int:

		updateTime, ok := timeData.(int)
		if !ok {
			log.Println("Convert UpdateTime(int) Failure")
			return 0, false
		}

		timeStamp = int64(updateTime)

	case int64:

		updateTime, ok := timeData.(int64)
		if !ok {
			log.Println("Convert UpdateTime(int64) Failure")
			return 0, false
		}

		timeStamp = updateTime

	case float64:

		updateTime, ok := timeData.(float64)
		if !ok {
			log.Println("Convert UpdateTime(float64) Failure")
			return 0, false
		}

		timeStamp = int64(updateTime)

	case string:

		updateTime, ok := timeData.(string)
		if !ok {
			log.Println("Convert UpdateTime(string) Failure")
			return 0, false
		}

		timeStamp, ok = tryToParseTime(updateTime)
		if !ok {
			log.Printf("Convert update time field format error, time: %s\n", updateTime)
			return 0, false
		}
	}

	return timeStamp, true
}

func tryToParseTime(str string) (int64, bool) {

	var err error
	var theTime time.Time
	const CST = 28800
	var CSTZone = time.FixedZone("CST", CST)

	for _, timeFormat := range timeFormatsNonTimeZone {

		theTime, err = time.ParseInLocation(timeFormat, str, CSTZone)
		if err == nil {

			return theTime.Unix(), true
		}
	}

	for _, timeFormat := range timeFormats {

		theTime, err = time.ParseInLocation(timeFormat, str, CSTZone)
		if err == nil {

			return theTime.Unix(), true
		}
	}

	return 0, false
}
