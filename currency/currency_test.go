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
		t.Failed()
	}
	return jsontype
}
func TestJsonPrimaryParsing(t *testing.T) {
	data := readTestFile(t)

	if len(data.Rates) != 31 {
		fmt.Printf("Failed to parse rates expected 31 got %i", len(data.Rates))
		t.Failed()
	}
}
