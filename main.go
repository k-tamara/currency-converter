package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

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
	}

	fmt.Println(currencyToAmount)
}
