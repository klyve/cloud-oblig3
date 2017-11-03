package currency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func readTestFile(t *testing.T) DataList {
	path := "./data/test.json"
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		t.Failed()
	}
	var jsontype DataList
	jsonError := json.Unmarshal(file, &jsontype)
	if jsonError != nil {
		fmt.Printf("Failed to parse json: %v\n", jsonError)
		t.Fail()
	}
	return jsontype
}

func TestFormatJson(t *testing.T) {
	path := "./data/test.json"
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		t.Failed()
	}
	data := FormatJsonResponse(file)
	if len(data.Rates) != 31 {
		fmt.Printf("Failed to parse rates expected 31 got %v", len(data.Rates))
		t.Fail()
	}
	if data.Base != "EUR" {
		fmt.Printf("Failed expected base to be EUR got %s", data.Base)
		t.Fail()
	}
	if data.Date != "2017-11-02" {
		fmt.Printf("Failed to get date expected 2017-11-02 got %s", data.Date)
		t.Fail()
	}
	if data.Rates["NOK"] != 9.4838 {
		fmt.Printf("Failed to get accurate rate for NOK 9.4838 got %f", data.Rates["NOK"])
		t.Fail()
	}
}
func TestDataIntegrity(t *testing.T) {
	data := readTestFile(t)

	if len(data.Rates) != 31 {
		fmt.Printf("Failed to parse rates expected 31 got %v", len(data.Rates))
		t.Fail()
	}
	if data.Base != "EUR" {
		fmt.Printf("Failed expected base to be EUR got %s", data.Base)
		t.Fail()
	}
	if data.Date != "2017-11-02" {
		fmt.Printf("Failed to get date expected 2017-11-02 got %s", data.Date)
		t.Fail()
	}
	if data.Rates["NOK"] != 9.4838 {
		fmt.Printf("Failed to get accurate rate for NOK 9.4838 got %f", data.Rates["NOK"])
		t.Fail()
	}
}

func TestDataFetchIntegrity(t *testing.T) {
	data := FetchLatest()

	if len(data.Rates) == 0 {
		fmt.Printf("Failed to parse rates expected 31 got %v", len(data.Rates))
		t.Fail()
	}
	if data.Base != "EUR" {
		fmt.Printf("Failed expected base to be EUR got %s", data.Base)
		t.Fail()
	}
}

func TestFetchJSONData(t *testing.T) {
	body, err := FetchJSONData("http://api.fixer.io/latest")
	if err != nil {
		fmt.Println("Some error ")
		t.Fail()
	}
	data := FormatJsonResponse(body)
	if len(data.Rates) == 0 {
		fmt.Printf("Failed to parse rates expected 31 got %v", len(data.Rates))
		t.Fail()
	}
	if data.Base != "EUR" {
		fmt.Printf("Failed expected base to be EUR got %s", data.Base)
		t.Fail()
	}
}

func TestCurrencyAs(t *testing.T) {
	core := readTestFile(t)
	data := core.As("NOK")

	if len(data.Rates) != 31 {
		fmt.Printf("Failed to parse rates expected 31 got %v", len(data.Rates))
		t.Fail()
	}
	if data.Base != "NOK" {
		fmt.Printf("Failed expected base to be NOK got %s", data.Base)
		t.Fail()
	}
	rateSek := fmt.Sprintf("%.3f", data.Rates["SEK"])
	if rateSek != "1.029" {
		fmt.Printf("Failed to get accurate rate for SEK 1.02 got %s", rateSek)
		t.Fail()
	}
}

func TestCurrencyFrom(t *testing.T) {
	core := readTestFile(t)
	data := core.From("NOK")

	if len(data.Rates) != 31 {
		fmt.Printf("Failed to parse rates expected 31 got %v", len(data.Rates))
		t.Fail()
	}
	if data.Base != "NOK" {
		fmt.Printf("Failed expected base to be NOK got %s", data.Base)
		t.Fail()
	}
	rateSek := fmt.Sprintf("%.3f", data.Rates["SEK"])
	if rateSek != "1.029" {
		fmt.Printf("Failed to get accurate rate for SEK 1.02 got %s", rateSek)
		t.Fail()
	}
}

func TestCurrencyTo(t *testing.T) {
	core := readTestFile(t)
	data := core.From("NOK").To("SEK")

	if data.From != "NOK" {
		fmt.Printf("Failed expected FROM to be NOK got %s", data.From)
		t.Fail()
	}
	if data.FromValue != 1.0 {
		fmt.Printf("Failed expected FROMVALUE to be 1.0 got %v", data.FromValue)
		t.Fail()
	}
	if data.To != "SEK" {
		fmt.Printf("Failed expected TO to be NOK got %s", data.To)
		t.Fail()
	}
	rateSek := fmt.Sprintf("%.3f", data.ToValue)

	if rateSek != "1.029" {
		fmt.Printf("Failed to get accurate rate for SEK 1.029 got %s", rateSek)
		t.Fail()
	}
	rateSek2 := fmt.Sprintf("%.3f", data.Rate)

	if rateSek2 != "1.029" {
		fmt.Printf("Failed to get accurate rate SEK 1.029 got %s", rateSek2)
		t.Fail()
	}
}
