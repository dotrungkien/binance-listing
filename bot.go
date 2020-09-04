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

	// newSymbols := []string{}

	for _, s := range symbols[:4] {
		symbolData := s.(map[string]interface{})
		fmt.Println(symbolData["symbol"])
		// newSymbols = append(newSymbols, string())
	}

	// err = ioutil.WriteFile("go-symbol.json", body, 0644)
	// check(err)

	// fmt.Println(result2["symbol"])

	// for _, val := range symbols[:4] {
	// }
	// fmt.Printf("%s", body)

}
