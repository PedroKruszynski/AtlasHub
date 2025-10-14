package udsbrl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type USDBRL struct {
	Code       string `json:"code"`
	Bid        string `json:"bid"` // valor de compra como string
	Ask        string `json:"ask"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	CreateDate string `json:"create_date"`
}

type APIResponse struct {
	USDBRL USDBRL `json:"USDBRL"`
}

func FetchDollar() (float64, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var apiResp map[string]USDBRL
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		return 0, err
	}
	dollarData, ok := apiResp["USDBRL"]
	if !ok {
		return 0, fmt.Errorf("campo USDBRL não encontrado")
	}
	// Converte bid (string) para float64
	var valor float64
	valor, err = strconv.ParseFloat(dollarData.Bid, 64)
	if err != nil {
		return 0, err
	}
	return valor, nil
}
