package utils

import (
	"fmt"
	"time"
)

type Time time.Time

// 2006-01-02 15:04:05
const (
	timeFormat = "2006年01月02日 15点04分"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format(timeFormat))
	return []byte(stamp), nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t Time) StringFormat(formatStr ...string) string {
	if len(formatStr) > 0 {
		return time.Time(t).Format(formatStr[0])
	}
	return time.Time(t).Format("01月02日 15:04")
}
