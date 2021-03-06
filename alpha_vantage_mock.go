package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Chart struct {
	Name  string
	Order int
}

type ResponseMetaData struct {
	Information   string
	Symbol        string
	LastRefreshed string
	OutputSize    string
	Timezone      string
}

type TimeSeriesDailyAdjusted struct {
	Open              float64 `json:"1. open"`
	High              float64 `json:"2. high"`
	Low               float64 `json:"3. low"`
	Close             float64 `json:"4. close"`
	Adjusted_Close    float64 `json:"5. adjusted close"`
	Volume            uint64  `json:"6. volume"`
	Dividend_Amount   float64 `json:"7. dividend amount"`
	Split_Coefficient float64 `json:"8. split coefficient"`
}

type Response struct {
	MetaData        ResponseMetaData                     `json:"Meta Data"`
	TimeSeriesDaily []map[string]TimeSeriesDailyAdjusted `json:"TimeSeries (Daily)"`
}

func (response *Response) String() string {
	output := "\n{\n"
	output += "    \"Meta Data\": {\n"
	output += "        "
	output += "\"1. Information\": \"" + response.MetaData.Information + "\",\n"
	output += "        "
	output += "\"2. Symbol\": \"" + response.MetaData.Symbol + "\",\n"
	output += "        "
	output += "\"3. Last Refreshed\": \"" + response.MetaData.LastRefreshed + "\",\n"
	output += "        "
	output += "\"4. Output Size\": \"" + response.MetaData.OutputSize + "\",\n"
	output += "        "
	output += "\"5. Time Zone\": \"" + response.MetaData.Timezone + "\"\n"
	output += "    },\n"
	output += "    \"TimeSeries (Daily)\": {\n"

	first_entry := true
	for _, timeseriesdaily := range response.TimeSeriesDaily {
		for datestamp := range timeseriesdaily {
			if !first_entry {
				output += "        },\n"
			} else {
				first_entry = false
			}

			output += "        \"" + datestamp + "\": {\n"
			output += "            \"1. open\": \"" + strconv.FormatFloat(timeseriesdaily[datestamp].Open, 'f', -1, 64) + "\",\n"
			output += "            \"2. high\": \"" + strconv.FormatFloat(timeseriesdaily[datestamp].High, 'f', -1, 64) + "\",\n"
			output += "            \"3. low\": \"" + strconv.FormatFloat(timeseriesdaily[datestamp].Low, 'f', -1, 64) + "\",\n"
			output += "            \"4. close\": \"" + strconv.FormatFloat(timeseriesdaily[datestamp].Close, 'f', -1, 64) + "\",\n"
			output += "            \"5. adjusted close\": \"" + strconv.FormatFloat(timeseriesdaily[datestamp].Adjusted_Close, 'f', -1, 64) + "\",\n"
			output += "            \"6. volume\": \"" + fmt.Sprintf("%d", timeseriesdaily[datestamp].Volume) + "\",\n"
			output += "            \"7. dividend amount\": \"" + strconv.FormatFloat(timeseriesdaily[datestamp].Dividend_Amount, 'f', 4, 64) + "\",\n"
			output += "            \"8. split coefficient\": \"" + strconv.FormatFloat(timeseriesdaily[datestamp].Split_Coefficient, 'f', 1, 64) + "\",\n"
		}
	}
	output += "        },\n    }\n}\n"

	return output
}

func GenerateTimeSeriesDailyAdjusted(symbol string, interval string) string {
	date := time.Now()

	// Skip today, AV's equiv API doesn't show today's data
	// if you query after-hours
	date = date.AddDate(0, 0, -1)

	original_date := date
	collected_series := make([]map[string]TimeSeriesDailyAdjusted, 1)
	for i := 0; i <= 100; i++ {
		year, _, day := date.Date()
		month := date.Month()
		date_string := fmt.Sprintf("%d-%d-%d", year, month, day)
		tsd := TimeSeriesDailyAdjusted{
			12.30,
			13.00,
			12.00,
			13.50,
			13.50,
			4074528,
			0.00,
			1.0,
		}
		collected_series = append(collected_series, map[string]TimeSeriesDailyAdjusted{date_string: tsd})

		if int(date.Weekday()) == 1 {
			date = date.AddDate(0, 0, -3)
		} else {
			date = date.AddDate(0, 0, -1)
		}
	}

	year, _, day := original_date.Date()
	// Yeah, this is a bit weird. I did it this way so that I can exactly
	// match AlphaVantage's responses, eg
	// "1. open": "133.9"
	// This lets me turn "Open" into "1. open".
	response := Response{
		ResponseMetaData{
			"Informational string goes here",
			symbol,
			fmt.Sprintf("%d-%d-%d", year, time.Now().Month(), day),
			"Compact",
			"US/Eastern",
		},
		collected_series,
	}
	return response.String()
}

func main() {

	r := setupRouter()
	r.Run(":8080")
}

//

func setupRouter() *gin.Engine {
	// Disable console color
	// gin.DisableConsoleColor
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:
	// /query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=IBM&interval=5min&apikey=xxx
	router.GET("/query", func(c *gin.Context) {
		function := c.Query("function")
		symbol := c.Query("symbol")
		interval := c.Query("interval")
		apikey := c.Query("apikey")

		if function == "" {
			c.String(http.StatusBadRequest, "Bad Request, 'function' parameter is empty")
			return
		}
		if symbol == "" {
			c.String(http.StatusBadRequest, "Bad Request, 'symbol' parameter is empty")
			return
		}
		if interval == "" {
			c.String(http.StatusBadRequest, "Bad Request, 'interval' parameter is empty")
			return
		}
		if apikey == "" {
			c.String(http.StatusBadRequest, "Bad Request, 'apikey' parameter is empty")
			return
		}

		c.String(http.StatusOK, GenerateTimeSeriesDailyAdjusted(symbol, interval))
	})

	return router
}
