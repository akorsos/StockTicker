package main

import (
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
	"encoding/json"
	"sort"
	"strconv"
)

func main() {
	nDays, err := strconv.Atoi(os.Getenv("NDAYS"))
	symbol := os.Getenv("SYMBOL")
	apikey := os.Getenv("APIKEY")

	//Run GET API call and error check response validity
	res, err := http.Get("https://www.alphavantage.co/query?apikey=" + apikey + "&function=TIME_SERIES_DAILY&symbol=" + symbol)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//Decode the json to a map of unknown type
	Response := map[string]interface{}{}
    json.Unmarshal([]byte(body), &Response)

	//Remove the Meta Data portion, not necessary but shows understanding
	delete(Response,"Meta Data")

	//Assert the Days var is same type as Response
	//This allows us to pull out the map within a map
	Days := Response["Time Series (Daily)"].(map[string]interface {})

	//Maps are unsorted, so we sort keys by date in a string map
	keys := make([]string, 0, len(Days))
	for k := range Days {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	
	//Define average closing var
	averageClosing := 0.0

	//Iterate over sorted key map and use key to find value
	//There is no descending sort, so we iterate backwards
	for i := len(keys)-1; i >= len(keys)-nDays; i-- {

		//Pull daily data into map
		Data := Days[keys[i]].(map[string]interface {})
		// fmt.Println(Data["4. close"])

		//Get closing price from each day, multiple vars for ease of reading
		//Assert type
		closingString := Data["4. close"].(string)
		//Convert string to float32
		closingFloat, err := strconv.ParseFloat(closingString, 32)
		if err != nil {
			log.Fatal(err)
		}

		averageClosing = averageClosing + closingFloat
	}

	//Calculate aver closing price
	var nDaysFloat float64 = float64(nDays)
	averageClosing = averageClosing / nDaysFloat

	//Format float32 to 123.4567 and print
	average := fmt.Sprintf("%.4f", averageClosing)
	fmt.Println(average)
}