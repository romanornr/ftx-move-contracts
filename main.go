package main

import (
	"fmt"
	"github.com/romanornr/ftx-move-contracts/csv"
)

func main() {
	file := csv.ScanFiles(".csv")
	records, _ := csv.ReadCSVFiles(file)
	moveRecords := csv.GetDailyMoveContractsRecords(records)
	stats := csv.AnalyzeDailyMoveContractRecords(moveRecords)
	fmt.Println(stats)
	//for _, stat := range stats.Static {
	//	const padding = 3
	//	w := tabwriter.NewWriter(os.Stdout, 10, 10, padding, ' ', tabwriter.TabIndent)
	//	fmt.Fprintf(w, "%s\t%s\t%s\t%.2f\t\n", stat.Type, stat.Day, stat.Time, stat.ExpirationPrice)
	//	w.Flush()
	//}
}
