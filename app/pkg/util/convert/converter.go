package convert

import (
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05.000000"
)

func TimeToStr(t time.Time) string {
	str := t.Format(timeFormat)
	return str
}