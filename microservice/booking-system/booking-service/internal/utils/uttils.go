package utils

import "time"

func FormatDate(times time.Time) (datetime string) {
	datetime = times.Format("2006-01-02T15:04:06")
	return datetime
}
