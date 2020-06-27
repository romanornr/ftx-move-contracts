package main

import (
	"fmt"
	_ "fmt"
	"github.com/romanornr/ftx-move-contracts/futures"
	"time"
)

func main() {
	expired := futures.GetExpiredFutures()
	MOVEContractsData := expired.GetDailyMOVEContracts()
	contracts := MOVEContractsData.AverageDailyMOVEContractsThisYear()
	days := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}

	for i := len(contracts.Expired)-1; i >= 0; i-- {
		fmt.Printf("%s\tGroup: %s\tDay:%s\tExpiration price:$%.2f\n",contracts.Expired[i].Description, contracts.Expired[i].Group, contracts.Expired[i].Expiry.Weekday().String(), contracts.Expired[i].Mark)
	}
	fmt.Println()

	for _, day := range days {
		move := contracts.AverageDayWeek(day)
		fmt.Printf("Average %s MOVE Contract expiration price on %s\t $%2.f\n", move.Expired[0].UnderlyingDescription, day, move.AverageExpirationPrice)
	}

	fmt.Printf("\nCurrent average daily FTX MOVE contracts expiration price for %d: $%2.f\n", time.Now().Year(), contracts.AverageExpirationPrice)
}
