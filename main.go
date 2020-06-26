package main

import (
	_"fmt"
	"github.com/romanornr/ftx-move-contracts/futures"
)

func main() {
	//file := csv.ScanFiles(".csv")
	//records, _ := csv.ReadCSVFiles(file)
	//moveRecords := csv.GetDailyMoveContractsRecords(records)
	//moveContractsData := csv.AnalyzeDailyMoveContractRecords(moveRecords)
	//stats := moveContractsData.DailyAverage("Wednesday")
	//
	//const padding = 3
	//w := tabwriter.NewWriter(os.Stdout, 10, 10, padding, ' ', tabwriter.TabIndent)
	//fmt.Printf("Average expiration price: %s\t%.2f\n", stats.Static[0].Type, moveContractsData.AverageExpirationPrice)
	//fmt.Fprintf(w, "Average expiration price: %s\t%s\t%.2f\t\n", "BTC-MOVE", stats.Day, stats.AverageExpirationPrice)
	//w.Flush()
	expired := futures.GetExpiredFutures()
	expired.DailyMOVE()

}
