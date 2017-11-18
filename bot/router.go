package bot

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/klyve/cloud-oblig3/currency"
)

var callRoutes map[string]func(RouterData) RouterData

func CreateRoutes() {
	callRoutes = map[string]func(RouterData) RouterData{
		"replaceUsername":       replaceUsername,
		"findLatest":            findLatest,
		"replaceCurrency":       replaceCurrency,
		"verifyCurrencyRequest": verifyCurrencyRequest,
	}
}

func Route(recipe Recipe, data RouterData) RouterData {
	rand.Seed(time.Now().Unix())
	message := recipe.Messages[rand.Intn(len(recipe.Messages))]
	data.Message = message
	data.Count = 0
	data.Error = false
	data.ErrorTo = ""

	for _, route := range recipe.Route {
		if callRoutes[route] != nil {
			data = callRoutes[route](data)
		}
		if data.Error {
			res := FindRecipe(data.ErrorTo)
			return Route(res, data)
		}
	}
	return data

}

func ReplaceCtx(message string, num int, ctx string) string {
	regex := `\{%v\}`
	regex = fmt.Sprintf(regex, num)

	var re = regexp.MustCompile(regex)
	return re.ReplaceAllString(message, ctx)
}

func replaceUsername(data RouterData) RouterData {
	data.Message = ReplaceCtx(data.Message, data.Count, data.Data["username"])
	data.Count++
	return data
}

func findLatest(data RouterData) RouterData {
	curr := currency.FetchLatest()

	value := curr.From(data.Data["currency-from"]).To(data.Data["currency-to"])

	var amount string

	amount = "1"
	if data.Data["amount"] != "" {
		amount = data.Data["amount"]
	}
	data.Data["amount"] = amount
	totalAmount, err := strconv.Atoi(amount)

	if err != nil {
		data.Data["rate"] = fmt.Sprintf("%.2f", value.Rate)
		return data
	}

	data.Data["rate"] = fmt.Sprintf("%.2f", value.Rate*float32(totalAmount))
	return data
}

func replaceCurrency(data RouterData) RouterData {

	data.Message = ReplaceCtx(data.Message, data.Count, data.Data["amount"])
	data.Count++

	data.Message = ReplaceCtx(data.Message, data.Count, data.Data["currency-from"])
	data.Count++
	data.Message = ReplaceCtx(data.Message, data.Count, data.Data["currency-to"])
	data.Count++
	data.Message = ReplaceCtx(data.Message, data.Count, data.Data["rate"])
	data.Count++

	return data
}

func verifyCurrencyRequest(data RouterData) RouterData {
	if data.Data["currency-from"] == "" {
		data.Error = true
	}
	if data.Data["currency-to"] == "" {
		data.Error = true
	}
	if data.Error {
		data.ErrorTo = "405"
	}
	return data
}
