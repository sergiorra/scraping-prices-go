package main

import (
	"github.com/sergiorra/scraping-prices-go/booking"
	"github.com/sergiorra/scraping-prices-go/pricetravel"
	"github.com/sergiorra/scraping-prices-go/trip"
)

func main() {
	booking.Scrap()
	pricetravel.Scrap()
	trip.Scrap()
}


