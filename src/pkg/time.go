package pkg

import (
	"time"
)

func GetPrettyDuration(tStart time.Time, tEnd time.Time) string {
	dur := tEnd.Sub(tStart)
	return dur.String()
}

func GetToday() (time.Time, time.Time) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return start, end
}
