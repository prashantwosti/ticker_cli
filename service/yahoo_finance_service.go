package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type YahooResponse struct {
	QuoteResponse struct {
		Result []struct {
			Language                      string  `json:"language"`
			Region                        string  `json:"region"`
			QuoteType                     string  `json:"quoteType"`
			QuoteSourceName               string  `json:"quoteSourceName"`
			Triggerable                   bool    `json:"triggerable"`
			Currency                      string  `json:"currency"`
			MarketState                   string  `json:"marketState"`
			Exchange                      string  `json:"exchange"`
			LongName                      string  `json:"longName"`
			UUID                          string  `json:"uuid"`
			ExchangeTimezoneName          string  `json:"exchangeTimezoneName"`
			ExchangeTimezoneShortName     string  `json:"exchangeTimezoneShortName"`
			GmtOffSetMilliseconds         int     `json:"gmtOffSetMilliseconds"`
			Market                        string  `json:"market"`
			EsgPopulated                  bool    `json:"esgPopulated"`
			FiftyTwoWeekHighChangePercent float64 `json:"fiftyTwoWeekHighChangePercent"`
			FiftyTwoWeekLow               float64 `json:"fiftyTwoWeekLow"`
			FiftyTwoWeekHigh              float64 `json:"fiftyTwoWeekHigh"`
			SharesOutstanding             int64   `json:"sharesOutstanding"`
			MarketCap                     int64   `json:"marketCap"`
			SourceInterval                int     `json:"sourceInterval"`
			ExchangeDataDelayedBy         int     `json:"exchangeDataDelayedBy"`
			Tradeable                     bool    `json:"tradeable"`
			FirstTradeDateMilliseconds    int64   `json:"firstTradeDateMilliseconds"`
			PriceHint                     int     `json:"priceHint"`
			RegularMarketChange           float64 `json:"regularMarketChange"`
			RegularMarketChangePercent    float64 `json:"regularMarketChangePercent"`
			RegularMarketTime             int     `json:"regularMarketTime"`
			RegularMarketPrice            float64 `json:"regularMarketPrice"`
			RegularMarketDayHigh          float64 `json:"regularMarketDayHigh"`
			RegularMarketDayRange         string  `json:"regularMarketDayRange"`
			RegularMarketDayLow           float64 `json:"regularMarketDayLow"`
			RegularMarketVolume           int     `json:"regularMarketVolume"`
			RegularMarketPreviousClose    float64 `json:"regularMarketPreviousClose"`
			FullExchangeName              string  `json:"fullExchangeName"`
			RegularMarketOpen             float64 `json:"regularMarketOpen"`
			FiftyTwoWeekLowChange         float64 `json:"fiftyTwoWeekLowChange"`
			FiftyTwoWeekLowChangePercent  float64 `json:"fiftyTwoWeekLowChangePercent"`
			FiftyTwoWeekRange             string  `json:"fiftyTwoWeekRange"`
			FiftyTwoWeekHighChange        float64 `json:"fiftyTwoWeekHighChange"`
			Symbol                        string  `json:"symbol"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"quoteResponse"`
}

type ErrorCode string

const (
	InvalidTicker ErrorCode = "Ticker code could be invalid. Please try again."
	Unknown       ErrorCode = "Something went terribly wrong. Check if your symbol is valid."
	Ok
)

func GetTicker(symbol string) (string, ErrorCode) {
	req, _ := http.NewRequest(
		"GET",
		"https://query2.finance.yahoo.com/v7/finance/quote",
		nil,
	)
	// Build query
	q := req.URL.Query()
	q.Add("symbols", symbol) // Add a new value to the set.
	req.URL.RawQuery = q.Encode()
	// Call api
	response, err := http.Get(req.URL.String())

	if err != nil || response.StatusCode != http.StatusOK {
		return "", InvalidTicker
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	yahooData := &YahooResponse{}

	decoder.Decode(yahooData)
	if len(yahooData.QuoteResponse.Result) > 0 {
		quote := yahooData.QuoteResponse.Result[0]

		var sb strings.Builder
		fmt.Fprintf(&sb, "%s\n", quote.LongName)
		fmt.Fprintf(&sb, "Price: %s %.2f\n", quote.Currency, quote.RegularMarketPrice)
		fmt.Fprintf(&sb, "Last change: %.2f%%", quote.RegularMarketChangePercent)

		return sb.String(), Ok
	}
	return "", Unknown
}
