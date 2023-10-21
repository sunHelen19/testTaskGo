package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

	for {
		fmt.Println("Какие курсы криптовалют Вы хотите получить?:")
		fmt.Println("1. Все курсы криптовалют")
		fmt.Println("2. Выбрать криптовалюту:")
		fmt.Print("Введите номер варианта ответа:")
		var response int
		fmt.Scan(&response)
		data, err := getData(url)
		if err != nil {
			fmt.Println(err)
		}
		if response == 1 {

			for _, currency := range data {
				printData(currency)

			}
		} else if response == 2 {
		ForLabel:
			for {
				fmt.Print("Введите название валюты:")
				var respCur string
				fmt.Scan(&respCur)

				respCur = prepareStr(respCur)

				var isCurrency = false
				for _, currency := range data {
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
			fmt.Println("урс")
			return nil, errJson
		}

	} else {
		fmt.Println("превышено количество запросов")
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
