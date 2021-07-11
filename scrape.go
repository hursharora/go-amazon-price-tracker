package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func scrapeItem(i item, ch chan item) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Print("Visiting: ", r.URL)
	})

	c.OnHTML("#priceblock_ourprice,#priceblock_dealprice", func(e *colly.HTMLElement) {
		log.Print("Getting price")
		priceText := e.Text
		priceText = priceText[1:]

		price, err := strconv.ParseFloat(priceText, 32)
		if err != nil {
			log.Fatal("Error converting price to float")
		}

		i.price = float32(price)
	})

	c.OnHTML("#productTitle", func(e *colly.HTMLElement) {
		log.Print("Getting title")
		i.name = strings.TrimSpace(e.Text)
	})

	c.Visit(i.url)

	ch <- i
}
