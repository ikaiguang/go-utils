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

// Today today
func Today() time.Time {
	t := time.Now()

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// ToDay 2019-08-21 22:07:07 -> 2019-08-21 00:00:00
func ToDay(t time.Time) time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// ToHour 2019-08-21 22:07:07 -> 2019-08-21 22:00:00
func ToHour(t time.Time) time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// ToMinute 2019-08-21 22:07:07 -> 2019-08-21 22:07:00
func ToMinute(t time.Time) time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// ThisMonth 2019-08-21 22:07:07 -> 2019-08-01 00:00:00
func ThisMonth(t time.Time) time.Time {
	y, m, _ := t.Date()

	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// ThisYear 2019-08-21 22:07:07 -> 2019-01-01 00:00:00
func ThisYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// TimestampToTime timestamp to Time
func TimestampToTime(u int64) time.Time {
	return time.Unix(u, 0)
}

// TimestampToDate timestamp to date
func TimestampToDate(u int64, format string) string {
	return time.Unix(u, 0).Format(format)
}

// DateToTime date to time
func DateToTime(format, date string) (time.Time, error) {
	return time.ParseInLocation(format, date, time.Local)
}

// FormatRFC3339 to RFC3339
func FormatRFC3339(t time.Time) string {
	return t.Format(RFC3339)
}
