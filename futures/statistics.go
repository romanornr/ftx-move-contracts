package futures

import (
	"time"
)

type Statistic struct {
	Type string
	Day time.Weekday
	Time string
	ExpirationPrice float64
	AverageExpirationPrice float64
}

type Statistics struct {
	Static []Statistic
	Day time.Weekday
	AverageExpirationPrice float64
}

func (statistics Statistics) DailyAverage(day string) Statistics{
	var daysCount = 0
	statistics.AverageExpirationPrice = 0 // reset
	for i := len(statistics.Static)-1; i >= 0; i-- {
		if statistics.Static[i].Day.String() != day {
			///fmt.Println(statistics.Static[i].Day)
			continue
		}
		daysCount += 1
		statistics.AverageExpirationPrice += statistics.Static[i].ExpirationPrice
		statistics.Day = statistics.Static[i].Day
	}
	statistics.AverageExpirationPrice = statistics.AverageExpirationPrice / float64(daysCount)
	return statistics
}