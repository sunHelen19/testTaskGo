package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

/*
Сделать парсер.

https://hypeauditor.com/top-instagram-all-russia/

Берем первую страницу результата выборки.
Собираем данные по всем колонками (рейтинг, имя, ник и т.д.)

На выходе получаем csv файл с 50 строчками.
*/
type Ranks struct {
	Rank      string
	Nick      string
	Name      string
	Category  string
	Followers string
	Country   string
	EngAuth   string
	EngAvg    string
}

var (
	RanksTable []Ranks
	url        = "https://hypeauditor.com/top-instagram-all-russia/"
)

func main() {

	parseData(url)
}

func parseData(url string) {
	var ranksList []Ranks

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var nodes []*cdp.Node
	chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Nodes(".row__top", &nodes, chromedp.ByQueryAll),
	)

	for _, node := range nodes {
		var rank, nick, name, category, followers, country, engAuth, engAvg string
		var categoryItems []*cdp.Node

		chromedp.Run(ctx,
			chromedp.Text(".rank span", &rank, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".contributor__name-content", &nick, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".contributor__title", &name, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Nodes(".tag", &categoryItems, chromedp.ByQueryAll, chromedp.FromNode(node), chromedp.AtLeast(0)),
			chromedp.Text(".subscribers", &followers, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".audience", &country, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".authentic", &engAuth, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".engagement", &engAvg, chromedp.ByQuery, chromedp.FromNode(node)),
		)

		for id, categoryItem := range categoryItems {

			var tag string
			chromedp.Run(ctx,
				chromedp.Text(".tag__content", &tag, chromedp.ByQuery, chromedp.FromNode(categoryItem)),
			)
			if id > 0 {
				div := ", "
				category += div
			}
			category += tag
		}

		ranks := Ranks{}

		ranks.Rank = rank
		ranks.Nick = nick
		ranks.Name = name
		ranks.Category = category
		ranks.Followers = followers
		ranks.Country = country
		ranks.EngAuth = engAuth
		ranks.EngAvg = engAvg

		ranksList = append(ranksList, ranks)

	}
	ExportInCSV(ranksList)
}

func ExportInCSV(RanksTable []Ranks) {
	file, err := os.Create("ranks.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{
		"Рейтинг",
		"Ник",
		"Имя",
		"Категория",
		"Подписчики",
		"Страна",
		"Eng.(Auth.)",
		"Eng.(Avg.)",
	}

	writer.Write(headers)

	for _, ranks := range RanksTable {
		record := []string{
			ranks.Rank,
			ranks.Nick,
			ranks.Name,
			ranks.Category,
			ranks.Followers,
			ranks.Country,
			ranks.EngAuth,
			ranks.EngAvg,
		}
		writer.Write(record)

	}
	defer writer.Flush()

}
