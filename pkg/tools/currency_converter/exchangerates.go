package currency_converter

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/avito-test-case/pkg/tools/configer"
)

type convertResponse struct {
	Success bool               `json:"success"`
	Base    string             `json:"base"`
	Date    string             `json:"date"`
	Rates   map[string]float64 `json:"rates"`
}

func ConvertCurrency(balance float64, currencyToConvert string) (float64, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	req, _ := http.NewRequest(http.MethodGet, configer.AppConfig.Secret.HttpUrl, nil)
	params := url.Values{}
	params.Add("access_key", configer.AppConfig.Secret.AccessSecret)
	params.Add("format", configer.AppConfig.Secret.Format)
	req.URL.RawQuery = params.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var convertData convertResponse
	err = json.Unmarshal(body, &convertData)
	if err != nil {
		return 0, err
	}

	convertRate, ok := convertData.Rates[currencyToConvert]
	if !ok {
		return 0, errors.New("incorrect convert currency")
	}
	baseRate, ok := convertData.Rates[configer.AppConfig.Secret.BaseCurrency]
	if !ok {
		return 0, errors.New("incorrect base currency")
	}

	return math.Floor(balance/(baseRate/convertRate)*100) / 100, nil
}
