package main

import (
	"fmt"
	"github.com/bclicn/color"
	"github.com/romanornr/ftx-move-contracts/futures"
	"time"
)

func main() {
	expired := futures.GetExpiredFutures()
	MOVEContractsData := expired.GetDailyMOVEContracts()
	contracts := MOVEContractsData.AverageDailyMOVEContractsThisYear()
	averageYearlyExpirationPrice := MOVEContractsData.AverageDailyMOVEContractsThisYear().AverageExpirationPrice

	days := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}

	for i := len(contracts.Expired) - 1; i >= 0; i-- {
		if contracts.Expired[i].Mark > averageYearlyExpirationPrice {
			fmt.Printf("%s\tGroup: %s\tDay:%s\tExpiration price: "+color.Green("$%.2f\n"), contracts.Expired[i].Description, contracts.Expired[i].Group, contracts.Expired[i].Expiry.Weekday().String(), contracts.Expired[i].Mark)
		} else {
			fmt.Printf("%s\tGroup: %s\tDay:%s\tExpiration price: "+color.Red("$%.2f\n"), contracts.Expired[i].Description, contracts.Expired[i].Group, contracts.Expired[i].Expiry.Weekday().String(), contracts.Expired[i].Mark)
		}
	}
	fmt.Println()

	for _, day := range days {
		move := contracts.AverageDayWeek(day)
		if move.AverageExpirationPrice > averageYearlyExpirationPrice {
			fmt.Printf("Average %s MOVE Contract expiration price on %s\t"+color.Green("$%2.f\n"), move.Expired[0].UnderlyingDescription, day, move.AverageExpirationPrice)
		} else {
			fmt.Printf("Average %s MOVE Contract expiration price on %s\t"+color.Red("$%2.f\n"), move.Expired[0].UnderlyingDescription, day, move.AverageExpirationPrice)
		}
	}

	fmt.Printf("\nCurrent average daily FTX MOVE contracts expiration price for %d: "+color.Green("$%2.f\n"), time.Now().Year(), contracts.AverageExpirationPrice)
}
