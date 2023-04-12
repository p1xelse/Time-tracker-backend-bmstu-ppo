package main

import (
	"src/models"
	"time"
)

func main() {
	t1 := time.Date(1984, time.November, 3, 13, 0, 0, 0, time.UTC)
	t2 := time.Date(1984, time.November, 3, 14, 0, 0, 0, time.UTC)
	g := models.Goal{TimeEnd: t2, TimeStart: t1}
	g.GetPrettyDuration()
}
