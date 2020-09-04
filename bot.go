package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func sendCW(message string) {
	err := godotenv.Load()
	checkError((err))

	cwToken := os.Getenv("CW_TOKEN")
	client := &http.Client{}
	checkError((err))

	request, err := http.NewRequest("POST", "https://api.chatwork.com/v2/rooms/199044484/messages?body="+url.QueryEscape(message), nil)
	checkError((err))

	request.Header.Set("X-ChatworkToken", cwToken)

	resp, err := client.Do(request)
	checkError((err))

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func main() {
	// url := "https://api.binance.com/api/v3/exchangeInfo"
	// resp, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(resp)

	// body, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s", body)
	sendCW("hi Kien")

}
