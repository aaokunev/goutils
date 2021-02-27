package timeformat

import (
	"fmt"
	"time"
)

func DurationToHMS(d time.Duration) string {
	d = d.Round(time.Second)

	sign := ""

	if d < 0 {
		sign = "-"
		d = d * -1
	}

	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second

	return fmt.Sprintf("%s%02d:%02d:%02d", sign, h, m, s)
}
