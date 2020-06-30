package main

import (
	"fmt"
	"github.com/romanornr/ftx-move-contracts/futures"
	"time"
	"bufio"
	"os"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}


func main() {
	expired := futures.GetExpiredFutures()
	MOVEContractsData := expired.GetDailyMOVEContracts()
	contracts := MOVEContractsData.AverageDailyMOVEContractsThisYear()
	averageYearlyExpirationPrice := MOVEContractsData.AverageDailyMOVEContractsThisYear().AverageExpirationPrice

	days := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}
	months := []time.Month{time.January, time.February, time.March, time.April, time.May, time.June, time.July, time.August, time.September, time.October, time.November, time.December}

	currentTime := time.Now()
	// open output file
    f, err := os.Create(currentTime.Format("2006-01-02")+".txt")
    check(err)
    defer f.Close()

	w := bufio.NewWriter(f)

	for i := len(contracts.Expired) - 1; i >= 0; i-- {
		if contracts.Expired[i].Mark > averageYearlyExpirationPrice {
			_, err = fmt.Fprintf(w, "%s\tGroup: %s\tDay:%s\tExpiration price: $%.2f\n", contracts.Expired[i].Description, contracts.Expired[i].Group, contracts.Expired[i].Expiry.Weekday().String(), contracts.Expired[i].Mark)
		} else {
			_, err = fmt.Fprintf(w, "%s\tGroup: %s\tDay:%s\tExpiration price: $%.2f\n", contracts.Expired[i].Description, contracts.Expired[i].Group, contracts.Expired[i].Expiry.Weekday().String(), contracts.Expired[i].Mark)
		}
	}
	_, err = fmt.Fprintf(w, "\n")


	for i, month := range months {
		move := contracts.AverageMonth(month)
		if move.Expired[i].Expiry.Month() == time.Now().Month() {
			break // break out of the loop
		}
		if move.AverageExpirationPrice > averageYearlyExpirationPrice {
			_, err = fmt.Fprintf(w, "Average %s %s MOVE Contract expiration price on %s\t$%2.f\n", move.Expired[i].Group, move.Expired[i].UnderlyingDescription, month, move.AverageExpirationPrice)
		} else {
			_, err = fmt.Fprintf(w, "Average %s %s MOVE Contract expiration price on %s\t$%2.f\n", move.Expired[i].Group, move.Expired[i].UnderlyingDescription, month, move.AverageExpirationPrice)
		}
	}

	// current month, using this method because index out of range [0] with length 0 by using months
	currentMonth := contracts.CurrentAverageMonth()
	if currentMonth.AverageExpirationPrice > averageYearlyExpirationPrice {
		_, err = fmt.Fprintf(w, "Average %s %s MOVE Contract expiration price on %s\t$%2.f\n",  currentMonth.Expired[0].Group, currentMonth.Expired[0].UnderlyingDescription, time.Now().Month(), currentMonth.AverageExpirationPrice)
	} else {
		_, err = fmt.Fprintf(w, "Average %s %s MOVE Contract expiration price on %s\t$%2.f\n", currentMonth.Expired[0].Group, currentMonth.Expired[0].UnderlyingDescription, time.Now().Month(), currentMonth.AverageExpirationPrice)

	}

	_, err = fmt.Fprintf(w, "\n")

	for _, day := range days {
		move := contracts.AverageDayWeek(day)
		if move.AverageExpirationPrice > averageYearlyExpirationPrice {
			_, err = fmt.Fprintf(w, "Average %s MOVE Contract expiration price on %s\t$%2.f\n", move.Expired[0].UnderlyingDescription, day, move.AverageExpirationPrice)
		} else {
			_, err = fmt.Fprintf(w, "Average %s MOVE Contract expiration price on %s\t$%2.f\n", move.Expired[0].UnderlyingDescription, day, move.AverageExpirationPrice)
		}
	}

	_, err = fmt.Fprintf(w, "\nCurrent average daily FTX MOVE contracts expiration price for %d: $%2.f\n\n", time.Now().Year(), contracts.AverageExpirationPrice)
	_, err = fmt.Fprintf(w, "Signup on FTX with this referral link receive a 10%% discount instead of the regular 5%% discount: %s\n", "https://ftx.com/#a=10percentDiscountOnFees")
	w.Flush()
}
