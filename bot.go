package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func sendCW(message string) {
	err := godotenv.Load()
	check((err))

	cwToken := os.Getenv("CW_TOKEN")
	client := &http.Client{}
	check((err))

	request, err := http.NewRequest("POST", "https://api.chatwork.com/v2/rooms/199044484/messages?body="+url.QueryEscape(message), nil)
	check((err))

	request.Header.Set("X-ChatworkToken", cwToken)

	resp, err := client.Do(request)
	check((err))

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func readSymbols() ([]string, error) {
	jsonData, err := ioutil.ReadFile("go-symbol.json")
	if err != nil {
		return nil, err
	}
	currentSymbols := []string{}
	err = json.Unmarshal(jsonData, &currentSymbols)
	if err != nil {
		return nil, err
	}
	return currentSymbols, nil
}

func writeSymbols(symbols []string) {
	jsonData, err := json.Marshal(symbols)
	check(err)
	err = ioutil.WriteFile("go-symbol.json", jsonData, 0644)
	check(err)
}

func inslice(n string, h []string) bool {
	for _, v := range h {
		if v == n {
			return true
		}
	}
	return false
}

func main() {
	url := "https://api.binance.com/api/v3/exchangeInfo"
	resp, err := http.Get(url)
	check(err)

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	check(err)

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	symbols := result["symbols"].([]interface{})

	newSymbols := []string{}

	for _, s := range symbols {
		symbolData := s.(map[string]interface{})
		newSymbols = append(newSymbols, symbolData["symbol"].(string))
	}

	currentSymbols, err := readSymbols()
	check(err)

	isNewFound := false

	for _, symbol := range newSymbols {
		if !inslice(symbol, currentSymbols) {
			fmt.Println("New Pair Listing Found: ", symbol)
			sendCW(symbol)
			isNewFound = true
		}
	}

	if isNewFound {
		writeSymbols(newSymbols)
	} else {
		fmt.Println("No new pair found")
	}
}
