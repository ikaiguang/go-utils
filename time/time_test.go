package gotime

import (
	"testing"
)

func TestFormat(t *testing.T) {
	tNow := now()

	t.Log(FormatSecond(tNow))
	t.Log(FormatMinute(tNow))
	t.Log(FormatHour(tNow))
	t.Log(FormatDay(tNow))
	t.Log(FormatMonth(tNow))
	t.Log(tNow.Year())
	t.Log(FormatRFC3339(tNow))
}
