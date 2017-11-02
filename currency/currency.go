package currency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PrintData() {
	l := FetchLatest()
	latest := l.As("NOK")
	fmt.Printf("Base %s \n", latest.Base)
	fmt.Printf("Date %s \n", latest.Date)
	for key, val := range latest.Rates {
		fmt.Printf("%s => %f \n", key, val)
	}
}

func PrintTo() {
	l := FetchLatest()
	latest := l.From("NOK")
	PrintD(latest.To("SEK"))
	PrintD(latest.To("EUR"))
	PrintD(latest.To("CAD"))
}
func PrintD(curr Convertion) {
	fmt.Printf("From %s To %s\n", curr.From, curr.To)
	fmt.Printf("Value1 %f Value2 %f\n", curr.FromValue, curr.ToValue)
	fmt.Printf("Rate %f\n\n", curr.Rate)
}

func FormatJsonResponse(data []byte) DataList {
	var jsontype DataList
	jsonError := json.Unmarshal(data, &jsontype)
	if jsonError != nil {
		fmt.Printf("Failed to parse json: %v\n", jsonError)
	}
	return jsontype
}
func FetchLatest() DataList {
	body, err := FetchJSONData("http://api.fixer.io/latest")
	if err != nil {
		fmt.Println("Some error ")
	}
	resp := FormatJsonResponse(body)
	return resp
}

// FetchJSONData fetches the json data from the web
func FetchJSONData(url string) ([]byte, interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, "Could not fetch data from github api"
	}
	// defer resp.Body.Close()
	body, readError := ioutil.ReadAll(resp.Body)
	// fmt.Printf("%v", body)
	if readError != nil {
		return nil, "Could not parse response body"
	}

	return body, nil
}
