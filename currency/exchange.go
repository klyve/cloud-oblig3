package currency

import "fmt"

// From sets the base value of a currency
func (data DataList) From(name string) DataList {
	return data.As(name)
}

// To returns the value from a currency to another
func (data DataList) To(name string) Convertion {
	var rate float32
	rate = 1.0
	if data.Base != name {
		rate = data.Rates[name]
	}

	return Convertion{
		From:      data.Base,
		FromValue: 1.0,
		To:        name,
		ToValue:   rate,
		Rate:      rate,
	}
}

// As changes the base currency
func (data DataList) As(name string) DataList {
	if data.Base == name {
		return data
	}
	var baseCurrency float32
	if data.Base == "EUR" {
		baseCurrency = 1.0
	} else {
		baseCurrency = data.Rates[data.Base]
	}
	data.Rates[data.Base] = 1 * GetRates(data.Rates[name], baseCurrency)

	data.Base = name
	baseValue := data.Rates[name]
	for key, value := range data.Rates {
		data.Rates[key] = 1 * GetRates(baseValue, value)
	}
	delete(data.Rates, name)
	fmt.Println(data.Rates)

	return data
}

// GetRates returns the currency rates
func GetRates(from float32, to float32) float32 {
	if from == to {
		return 1.0
	}
	return to * (1 / from)
}
