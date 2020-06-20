package main

import (
	"fmt"
	"github.com/romanornr/ftx-move-contracts/csv"
)

func main() {
	file := csv.ScanFiles(".csv")
	records, _ := csv.ReadCSVFiles(file)
	fmt.Println(csv.GetDailyMoveContractsRecords(records))
}
