package update_rate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ducktyst/goapi/internal/database"
	"io/ioutil"
	"net/http"
	"strconv"
)

var codesCurrencides = map[int]string{
	840: "USD",
	978: "EUR",
	643: "RUB",
}

var currencyCodes = map[string]int{
	"USD": 840,
	"EUR": 978,
	"RUB": 643,
}

type NullString string

func (ns NullString) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if len(string(ns)) == 0 {
		buf.WriteString("null")
	} else {
		buf.WriteString(`"` + string(ns) + `"`)
	}
	return buf.Bytes(), nil
}

func (ns *NullString) UnmarshalJSON(in []byte) error {
	str := string(in)
	if str == "null" {
		*ns = ""
		return nil
	}

	res := NullString(str)
	if len(res) >= 2 {
		res = res[1 : len(res)-1] //remove the wrapped qutation
	}
	*ns = res
	return nil
}

type ExchangeRateResp struct {
	Amount    float64    `json:"amount"`
	CrossRate NullString `json:"crossRate,omitempty"`
}

func UpdateRates(d *database.DB) {
	fmt.Println("UpdateRates")

	rates, err := d.GetExchangeRates(context.Background())
	if err != nil {
		fmt.Println("error", err)
		return
	}

	for _, er := range rates {
		from := currencyCodes[er.CurrencyFrom]
		fromCode := strconv.Itoa(from)
		if fromCode == "643" { // специфика сбера
			fromCode = ""
		}
		to := currencyCodes[er.CurrencyTo]
		toCode := strconv.Itoa(to)
		if toCode == "643" {
			toCode = ""
		}

		requestUrl := "https://www.sberbank.ru/portalserver/proxy/?pipe=shortCachePipe&url=http%3A%2F%2Flocalhost%2Frates-web%2FrateService%2Frate%2Fconversion%3FregionId%3D77%26sourceCode%3Dcard%26destinationCode%3Daccount%26exchangeType%3Dibank%26servicePack%3Dempty%26fromCurrencyCode%3D" + fromCode + "%26toCurrencyCode%3D" + toCode + "%26amount%3D1"
		resp, err := http.Get(requestUrl)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error", err)
		}

		var exchangeRate ExchangeRateResp
		err = json.Unmarshal(responseData, &exchangeRate)
		if err != nil {
			fmt.Println("error", err)
			return
		}

		er.Cource = exchangeRate.Amount
		err = d.UpdateExchangeRate(context.Background(), er)
		if err != nil {
			fmt.Println("error", err)
			return
		}
	}
}
