package csv

import (
	"bufio"
	"encoding/csv"
	"github.com/romanornr/ftx-move-contracts/futures"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
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

			//fmt.Println(records)


			//if strings.Contains(record[0], "BTC-MOVE") {
			//	fmt.Println(record)
			//}

			//amountString := strings.TrimRight(record[2], " XBt")     // 88,055,513 XBt = 0.88 btc
			//amountString = strings.ReplaceAll(amountString, ",", "") // remove the  ,
			//amount, err := strconv.ParseFloat(amountString, 64)
			//if err != nil {
			//	panic(err)
			//}
			//
			//fee, _ := strconv.ParseFloat(record[3], 64)
			//
			//tx := account.Transaction{
			//	Time:          record[0],
			//	Type:          record[1],
			//	Amount:        amount, //amount,
			//	Fee:           fee,
			//	Address:       record[4],
			//	Status:        record[5],
			//	WalletBalance: record[6],
			//}
			//
			//transactions = append(transactions, tx)
		}
	}
return records, nil
}

func GetDailyMoveContractsRecords(records [][]string) [][]string {
	var expiredFutures futures.ExpiredFutures
	var moveContracts [][]string
	for _, record := range records {
		expiredFutures.Ticker = record[0]
		expiredFutures.Name = record[1]
		if strings.Contains(expiredFutures.Ticker, "BTC-MOVE") && !strings.Contains(expiredFutures.Name, "Weekly MOVE"){
			moveContracts = append(moveContracts, record)
		}
	}
	return moveContracts
}
