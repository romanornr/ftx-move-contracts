package futures

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const LAYOUT = "2006-01-02T15:04:05Z07:00" // RFC3339

type ExpiredFuturesResponse struct {
	Result []struct {
		Ask                   interface{} `json:"ask"`
		Bid                   interface{} `json:"bid"`
		Description           string      `json:"description"`
		Enabled               bool        `json:"enabled"`
		Expired               bool        `json:"expired"`
		Expiry                time.Time   `json:"expiry"`
		ExpiryDescription     string      `json:"expiryDescription"`
		Group                 string      `json:"group"`
		ImfFactor             float64     `json:"imfFactor"`
		Index                 float64     `json:"index"`
		Last                  float64     `json:"last"`
		LowerBound            float64     `json:"lowerBound"`
		MarginPrice           float64     `json:"marginPrice"`
		Mark                  float64     `json:"mark"`
		MoveStart             interface{} `json:"moveStart"`
		Name                  string      `json:"name"`
		Perpetual             bool        `json:"perpetual"`
		PositionLimitWeight   float64     `json:"positionLimitWeight"`
		PostOnly              bool        `json:"postOnly"`
		PriceIncrement        float64     `json:"priceIncrement"`
		SizeIncrement         float64     `json:"sizeIncrement"`
		Type                  string      `json:"type"`
		Underlying            string      `json:"underlying"`
		UnderlyingDescription string      `json:"underlyingDescription"`
		UpperBound            float64     `json:"upperBound"`
	} `json:"result"`
	Success bool `json:"success"`
}

type ExpiredFutures struct {
	ExpiredFutures []ExpiredFuture
}

type ExpiredFuture struct {
	Ask                   interface{} `json:"ask"`
	Bid                   interface{} `json:"bid"`
	Description           string      `json:"description"`
	Enabled               bool        `json:"enabled"`
	Expired               bool        `json:"expired"`
	Expiry                time.Time   `json:"expiry"`
	ExpiryDescription     string      `json:"expiryDescription"`
	Group                 string      `json:"group"`
	ImfFactor             float64     `json:"imfFactor"`
	Index                 float64     `json:"index"`
	Last                  float64     `json:"last"`
	LowerBound            float64     `json:"lowerBound"`
	MarginPrice           float64     `json:"marginPrice"`
	Mark                  float64     `json:"mark"`
	MoveStart             interface{} `json:"moveStart"`
	Name                  string      `json:"name"`
	Perpetual             bool        `json:"perpetual"`
	PositionLimitWeight   float64     `json:"positionLimitWeight"`
	PostOnly              bool        `json:"postOnly"`
	PriceIncrement        float64     `json:"priceIncrement"`
	SizeIncrement         float64     `json:"sizeIncrement"`
	Type                  string      `json:"type"`
	Underlying            string      `json:"underlying"`
	UnderlyingDescription string      `json:"underlyingDescription"`
	UpperBound            float64     `json:"upperBound"`
}

func GetExpiredFutures() ExpiredFuturesResponse {
	apiEndpoint, err := url.Parse("https://ftx.com/api/expired_futures")
	if err != nil {
		logrus.Fatalf("Could not parse expired futures api endpoint: %s\n", err)
	}

	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Second * 15,
	}
	req, err := http.NewRequest(http.MethodGet, apiEndpoint.String(), nil)
	if err != nil {
		logrus.Fatalf("Could not make a new request to %s\n", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatalf("Client.Do failed for API endpoint %s %s\n", apiEndpoint.String(), err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("Could not read body response from API endpoint %s %s\n", apiEndpoint.String(), err)
	}

	expiredFutures := ExpiredFuturesResponse{}
	err = json.Unmarshal(body, &expiredFutures)
	if err != nil {
		logrus.Fatalf("Failed to unmarshal ExpiredFutures: %s\n", err)
	}

	return expiredFutures
}

func (expiredFuturesResp ExpiredFuturesResponse) GetDailyMOVEContracts() ExpiredFutures {
	response := expiredFuturesResp.Result
	expiredFutures := new(ExpiredFutures)

	for _, expiredFuture := range response {
		if expiredFuture.Type == "move" && expiredFuture.Underlying == "BTC" && expiredFuture.Expired == true {
			expiredFutures.ExpiredFutures = append(expiredFutures.ExpiredFutures, expiredFuture)
		}
	}
	return *expiredFutures
}

// Get total expiration price of Daily MOVE Contracts per year
func (expiredFutures ExpiredFutures) AverageDailyMOVEContractsThisYear() MOVEContracts {
	var MOVEContracts = new(MOVEContracts)
	var totalMOVEContractsThisYear float64
	currentYear := time.Now().Year()
	for _, expiredMOVEContract := range expiredFutures.ExpiredFutures {
		if expiredMOVEContract.Group == "daily" {
			if expiredMOVEContract.Expiry.Year() == currentYear {
				totalMOVEContractsThisYear += 1
				MOVEContracts.AverageExpirationPrice += expiredMOVEContract.Mark
				MOVEContracts.Expired = append(MOVEContracts.Expired, expiredMOVEContract)
			}
		}
	}
	MOVEContracts.AverageExpirationPrice = MOVEContracts.AverageExpirationPrice / totalMOVEContractsThisYear
	return *MOVEContracts
}

func (moveContracts MOVEContracts) AverageDayWeek(day time.Weekday) MOVEContracts {
	var MOVEContracts = new(MOVEContracts)
	var totalMOVEContractsWeekDay float64
	for _, moveContract := range moveContracts.Expired {
		if moveContract.Expiry.Weekday() == day {
			totalMOVEContractsWeekDay += 1
			MOVEContracts.Expired = append(MOVEContracts.Expired, moveContract)
			MOVEContracts.AverageExpirationPrice += moveContract.Mark
		}
	}
	MOVEContracts.AverageExpirationPrice = MOVEContracts.AverageExpirationPrice / totalMOVEContractsWeekDay
	return *MOVEContracts
}
