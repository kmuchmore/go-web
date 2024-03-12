package domain

import "time"

type Filters struct {
	DateRange DateRange
	Glob      string
}

type DateRange struct {
	Start time.Time
	End   time.Time
}

func (d DateRange) Valid() bool {
	return d.Start.Before(d.End)
}
