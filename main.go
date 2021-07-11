package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
)

type item struct {
	url         string
	email       string
	name        string
	price       float32
	targetPrice float32
}

func main() {
	godotenv.Load()
	tracks := loadItemTrack()

	// fmt.Print(os.Getenv("EMAIL"))
	c := make(chan item)

	for _, track := range tracks {
		go scrapeItem(track, c)
	}

	// infinite loop to check price every scrape price of item every 5 minutes
	for i := range c {
		if i.price <= i.targetPrice {
			log.Printf("The price of \"%s\" is $%.2f! Below target price of $%.2f", i.name, i.price, i.targetPrice)
			log.Printf("Emailing %s", i.email)
		} else {
			log.Printf("The price of \"%s\" is above target, checking again in 5 minutes", i.name)
			go func(i item) {
				time.Sleep(5 * time.Minute)
				go scrapeItem(i, c)
			}(i)
		}
	}
}

func check(err error, msg string) {
	if err != nil {
		log.Fatal(msg)
	}
}

func assert(cond bool, msg string) {
	if !cond {
		log.Fatal(msg)
	}
}
