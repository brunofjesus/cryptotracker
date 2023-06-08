package yahoo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetCurrentValue(ticker string) (string, error) {

	response, err := http.Get(fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s", ticker))
	if err != nil {
		return "", err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return "", fmt.Errorf("got error response code: %d", response.StatusCode)
	}

	var data yahooResponse
	dec := json.NewDecoder(response.Body)
	err = dec.Decode(&data)

	if err != nil {
		return "", err
	}

	if len(data.Chart.Result) == 0 {
		return "", errors.New("unexpected response payload from yahoo finance")
	}

	return fmt.Sprintf("%f", data.Chart.Result[0].Meta.RegularMarketPrice), nil
}
