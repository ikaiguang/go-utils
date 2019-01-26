package gotime

import (
	"time"
)

// Date Format
const (
	YmdHms  = "2006-01-02 15:04:05"
	YmdHm   = "2006-01-02 15:04"
	YmdH    = "2006-01-02 15"
	Ymd     = "2006-01-02"
	Ym      = "2006-01"
	RFC3339 = time.RFC3339
)

// now time.Now()
func now() time.Time {
	return time.Now()
}

// FormatSecond to YmdHms
func FormatSecond(t time.Time) string {
	return t.Format(YmdHms)
}

// FormatMinute to YmdHm
func FormatMinute(t time.Time) string {
	return t.Format(YmdHm)
}

// FormatHour to YmdH
func FormatHour(t time.Time) string {
	return t.Format(YmdH)
}

// FormatDay to Ymd
func FormatDay(t time.Time) string {
	return t.Format(Ymd)
}

// FormatMonth to Ym
func FormatMonth(t time.Time) string {
	return t.Format(Ym)
}

// FormatRFC3339 to RFC3339
func FormatRFC3339(t time.Time) string {
	return t.Format(RFC3339)
}
