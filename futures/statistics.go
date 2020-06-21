package futures

import "time"

type Statistic struct {
	Type string
	Day time.Weekday
	Time string
	ExpirationPrice float64
}

type Statistics struct {
	Static []Statistic
}