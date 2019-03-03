package main

import (
	"regexp"
	"strconv"
	"time"
)

var apiDateTimeFormat = "2006-01-02T15:04:05"

func MustParseRFC3339(sTime string) (rt time.Time) {
	rt, _ = time.Parse(apiDateTimeFormat, sTime)
	return rt
}

// rfc3339 datetime format
func formatTime(t time.Time) string {
	return t.UTC().Format(apiDateTimeFormat)
}

var reg = regexp.MustCompile("/Date\\((.*?)\\)/")

// ParseODataTimeStamp from /Date(d)/
func ParseODataTimeStamp(in string) time.Time {
	tss := reg.FindAllStringSubmatch(in, -1)
	if len(tss) == 1 && len(tss[0]) == 2 {
		oDataTimeStamp, _ := strconv.ParseInt(tss[0][1], 10, 64)
		return time.Unix(0, oDataTimeStamp*1000000)
	} else {
		return time.Unix(0, 0)
	}

}
