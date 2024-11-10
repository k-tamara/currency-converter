package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type exchangeRatesResponse struct {
	Disclamer string             `json:"disclaimer"`
	License   string             `json:"license"`
	Timestamp int                `json:"timestamp"`
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	ss := strings.Fields(text)

	currencies := map[string]string{
		"C$":  "CAD",
		"€":   "EUR",
		"Ft":  "HUF",
		"£":   "GBP",
		"₴":   "UAH",
		"zł":  "PLN",
		"₺":   "TRY",
		"kr.": "DKK",
		"lei": "RON",
		"$":   "USD",
		"R$":  "BRL",
		"MX$": "MXN",
		"Rp":  "IDR",
		"A$":  "AUD",
		"NZ$": "NZD",
		"E£":  "EGP",
		"DH":  "MAD",
		"R":   "ZAR",
		"¥":   "JPY",
		"CN¥": "CNY",
		"₽":   "RUB",
		"₹":   "INR",
		"₪":   "ILS",
		"CHF": "CHF",
		"kn":  "HRK",
	}
	currencyToAmount := make(map[string]float64)
	currencyCodes := make([]string, 0, len(currencies))

	for i := 0; i < len(ss); i++ {
		item := ss[i]
		var amountString string
		var currencySymbol string
		for i, r := range item {
			if unicode.IsDigit(r) {
				amountString = item[i:]
				currencySymbol = item[:i]
				break
			}
		}
		amountNoCommas := strings.Replace(amountString, ",", "", -1)
		amountFloat, err := strconv.ParseFloat(strings.TrimSpace(amountNoCommas), 64)
		if err != nil {
			fmt.Printf("Error parsing amount %v: %v\n", amountNoCommas, err)
			continue
		}
		currencyCode := currencies[currencySymbol]
		if len(currencyCode) == 0 {
			fmt.Printf("Currency symbol %v is not supported\n", currencySymbol)
			continue
		}
		currencyToAmount[currencyCode] = amountFloat
		currencyCodes = append(currencyCodes, currencyCode)
	}
	fmt.Println(currencyCodes)
	fmt.Println(currencyToAmount)

	url := "https://openexchangerates.org/api/latest.json?app_id=4fc45f8092aa474788dde3b4169eeb4b&symbols='" + strings.Join(currencyCodes, ",") + "'&prettyprint=false&show_alternative=false"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	var exchangeRates exchangeRatesResponse

	if err := json.NewDecoder(res.Body).Decode(&exchangeRates); err != nil {
		fmt.Println("error reading json response body: ", err)
	}

	fmt.Println(exchangeRates.Rates)
	var total float64
	for currencyCode, exchangeRate := range exchangeRates.Rates {
		total += currencyToAmount[currencyCode] / exchangeRate
		fmt.Printf("%v %v = %v USD\n", currencyToAmount[currencyCode], currencyCode, currencyToAmount[currencyCode]/exchangeRate)
	}

	fmt.Println("----------------------------------------------------------------")
	fmt.Printf("Total: %v USD", total)
}
