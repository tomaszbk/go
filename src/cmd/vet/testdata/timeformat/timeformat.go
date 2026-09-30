package timeformat

import "time"

func _(t time.Time) {
	_ = t.Format("2006-02-01") // ERROR "2006-02-01 should be 2006-01-02"
}
