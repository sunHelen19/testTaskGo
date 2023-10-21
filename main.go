package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

/*
Сделать клиента для получения курсов.

https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1

Добавить возможность получать курс для определенной криптовалюты.

Курсы обновляем не чаще чем раз в 10 минут.

*/

type Data struct {
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
}

var url = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"

func main() {
	timeDownload := time.Now().Add(-10 * time.Minute)
	var dataStore []Data

	for {
		fmt.Println("Какие курсы криптовалют Вы хотите получить?:")
		fmt.Println("1. Все курсы криптовалют")
		fmt.Println("2. Выбрать криптовалюту")
		fmt.Print("Введите номер варианта ответа:")
		var response int
		fmt.Scan(&response)

		currentTime := time.Now()

		nextTimeDownload := timeDownload.Add(10 * time.Minute)

		if currentTime.After(nextTimeDownload) {

			data, err := getData(url)
			dataStore = data
			timeDownload = time.Now()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			timeRange := time.Until(nextTimeDownload)
			dur := timeRange.Round(time.Minute).Minutes()
			//dur := timeRange.Minutes()
			fmt.Printf("\nСледующая загрузка курсов не ранее чем через %v минут!\n\n", dur)
		}

		if response == 1 {

			for _, currency := range dataStore {
				printData(currency)

			}
		} else if response == 2 {
		ForLabel:
			for {
				fmt.Print("Введите название криптовалюты:")
				var respCur string
				fmt.Scan(&respCur)

				respCur = prepareStr(respCur)

				var isCurrency = false
				for _, currency := range dataStore {
					cripto := prepareStr(currency.Name)
					if cripto == respCur {
						printData(currency)
						isCurrency = true
						break ForLabel
					}
				}
				if isCurrency == false {
					fmt.Println("Некорректный ввод криптовалюты. Попробуйте еще раз!")
				}
			}

		}
		fmt.Println()
	}

}

func getData(path string) ([]Data, error) {
	dataStore := []Data{}

	resp, err := http.Get(path)
	if err != nil {

		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, errReadBody := io.ReadAll(resp.Body)
		if errReadBody != nil {

			return nil, errReadBody
		}

		if errJson := json.Unmarshal(body, &dataStore); errJson != nil {
			return nil, errJson
		}

	}
	return dataStore, nil
}
func printData(currency Data) {
	fmt.Println(currency.Name, "-", currency.CurrentPrice)
}

func prepareStr(str string) string {
	str = strings.Trim(str, " ")
	str = strings.ToLower(str)
	return str
}
