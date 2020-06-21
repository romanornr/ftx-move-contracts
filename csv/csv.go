package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/romanornr/ftx-move-contracts/futures"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ScanFiles(extension string) []string {

	csvFiles := []string{}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			r, err := regexp.MatchString(extension, f.Name())
			if err == nil && r {
				csvFiles = append(csvFiles, f.Name())
			}
		}
	}
	return csvFiles
}

func ReadCSVFiles(csvfiles []string) ([][]string, error) {

	//var transactions []account.Transaction

	// Range over multiple CSV files
	var records [][]string
	for _, c := range csvfiles {
		csvFile, _ := os.Open("./" + c)
		r := csv.NewReader(bufio.NewReader(csvFile))

		r.LazyQuotes = true
		r.Comma = ','

		_, err := r.Read()
		if err != nil && err != io.EOF {
			return nil, err
		}

		//var error error
		for {
			record, error := r.Read()
			if error == io.EOF {
				break
			} else if error != nil {
				log.Fatal(error)
			}

			records = append(records, record)
		}
	}
	return records, nil
}

func GetDailyMoveContractsRecords(records [][]string) [][]string {
	var expiredFutures futures.ExpiredFutures
	var moveContractsRecords [][]string
	for _, record := range records {
		expiredFutures.Ticker = record[0]
		expiredFutures.Name = record[1]
		if strings.Contains(expiredFutures.Ticker, "BTC-MOVE") && !strings.Contains(expiredFutures.Name, "Weekly MOVE") {
			moveContractsRecords = append(moveContractsRecords, record)
		}
	}
	return moveContractsRecords
}

func AnalyzeDailyMoveContractRecords(records [][]string) futures.Statistics{
	var statistics futures.Statistic
	var totalStats futures.Statistics
	layout := "2006-01-02T15:04:05Z07:00" // RFC3339
	var averageExpirationPrice float64
	var amountContracts float64
	for i := len(records)-1; i >= 0; i-- {
		statistics.Type = "BTC-MOVE"
		recordTime, _ := time.Parse(layout, records[i][2])
		statistics.Day = recordTime.Weekday()
		statistics.Time = fmt.Sprintf("%d-%s-%d", recordTime.Day(), recordTime.Month().String(), recordTime.Year())

		currentYear := time.Now().Year()
		if recordTime.Year() <  currentYear {
			continue
		}
		amountContracts += 1

		price, _ := strconv.ParseFloat(records[i][3], 64)
		statistics.ExpirationPrice = math.Round(price*100)/100

		averageExpirationPrice += price
		totalStats.Static = append(totalStats.Static, statistics)
	}

	return totalStats
}

