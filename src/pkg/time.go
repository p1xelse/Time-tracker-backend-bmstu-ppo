package pkg

import (
	"time"
)

func GetPrettyDuration(tStart time.Time, tEnd time.Time) string {
	dur := tEnd.Sub(tStart)
	return dur.String()
}
