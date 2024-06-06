package priceconv

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"

	"strconv"
	"strings"

	"github.com/MirasDragonite/priceconv/logger"

	"golang.org/x/net/html/charset"
)

func (vc *Valutes) ConvertCurrency(amount float64, fromCode, toCode string) (float64, error) {

	fromValute, exist := vc.Valutes[fromCode]
	if !exist {
		return 0, errors.New("wrong valute code or data not comming in database")
	}

	toValute, exist := vc.Valutes[toCode]
	if !exist {
		return 0, errors.New("wrong valute code or data not comming in database")
	}

	fromRate, err := strconv.ParseFloat(strings.Replace(fromValute.Value, ",", ".", -1), 64)
	if err != nil {
		return 0, fmt.Errorf("invalid rate value for %s: %s", fromCode, fromValute.Value)
	}

	toRate, err := strconv.ParseFloat(strings.Replace(toValute.Value, ",", ".", -1), 64)
	if err != nil {
		return 0, fmt.Errorf("invalid rate value for %s: %s", toCode, toValute.Value)
	}

	fromNominal, err := strconv.ParseFloat(fromValute.Nominal, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid nominal value for %s: %s", fromCode, fromValute.Nominal)
	}
	toNominal, err := strconv.ParseFloat(toValute.Nominal, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid nominal value for %s: %s", toCode, toValute.Nominal)
	}
	convertedAmount := (amount * fromRate / fromNominal) * toNominal / toRate
	return roundToTwoDecimalPlaces(convertedAmount), nil
}

func (p *Valutes) GetDailyRate() {
	logger.Info.Println("Fetching daily rates...")

	url := "http://www.cbr.ru/scripts/XML_daily.asp"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error.Println("Error during creating new request:", err)
		return
	}

	// Set User-Agent header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		logger.Error.Println("Error during sending request:", err)
		return
	}
	defer resp.Body.Close()

	xmlData, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error.Println("Can't read the response from the api:", err)
		return
	}

	decoder := xml.NewDecoder(strings.NewReader(string(xmlData)))
	decoder.CharsetReader = charset.NewReaderLabel
	menu := new(ValCurs)
	err = decoder.Decode(menu)
	if err != nil {
		logger.Error.Println("Can't decode the comming data:", err)
		return
	}
	p.Valutes["RUB"] = Valute{
		ID:        "0",
		Nominal:   "1",
		Name:      "RUB",
		CharCode:  "RUB",
		Value:     "1,0",
		VunitRate: "1,0",
	}
	for _, el := range menu.Valutes {
		p.Valutes[el.CharCode] = el
	}
}

func (p Valutes) ConverStringPrices(price, currency string) (float64, error) {
	prices := strings.Split(price, " ")
	if len(prices) != 2 {
		return 0, errors.New("wrong price")
	}
	amount, err := strconv.ParseFloat(prices[0], 64)

	if amount <= 0 || err != nil {
		return 0, errors.New("wrong price")
	}
	result, err := p.ConvertCurrency(amount, prices[1], currency)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func roundToTwoDecimalPlaces(num float64) float64 {
	return math.Round(num*100) / 100
}
