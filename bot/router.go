package bot

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/klyve/cloud-oblig2/currency"
)

var callRoutes map[string]func(RouterData) RouterData

func CreateRoutes() {
	callRoutes = map[string]func(RouterData) RouterData{
		"testFunction1":   testFunction1,
		"testFunction2":   testFunction2,
		"replaceUsername": replaceUsername,
		"findLatest":      findLatest,
		"replaceCurrency": replaceCurrency,
	}
}

func Route(recipe Recipe, data RouterData) RouterData {
	rand.Seed(time.Now().Unix())
	message := recipe.Messages[rand.Intn(len(recipe.Messages))]
	data.Message = message
	data.Count = 0
	// data := RouterData{
	// 	Message: message,
	// 	Count:   0,
	// }
	for _, route := range recipe.Route {
		if callRoutes[route] != nil {
			data = callRoutes[route](data)
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

	value := curr.From(data.Data["baseCurrency"]).To(data.Data["targetCurrency"])

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

	data.Message = ReplaceCtx(data.Message, data.Count, data.Data["baseCurrency"])
	data.Count++
	data.Message = ReplaceCtx(data.Message, data.Count, data.Data["targetCurrency"])
	data.Count++
	data.Message = ReplaceCtx(data.Message, data.Count, data.Data["rate"])
	data.Count++

	return data
}

func testFunction1(data RouterData) RouterData {
	data.Message = ReplaceCtx(data.Message, data.Count, "TestFunction1")
	data.Count++
	return data
}

func testFunction2(data RouterData) RouterData {
	data.Message = ReplaceCtx(data.Message, data.Count, "TestFunction2")
	data.Count++
	return data
}
