package main

import (
	"github.com/sergiorra/scraping-prices-go/internal/booking"
	"github.com/sergiorra/scraping-prices-go/internal/pricetravel"
	"github.com/sergiorra/scraping-prices-go/internal/trip"
)

func main() {
	booking.Scrap()
	pricetravel.Scrap()
	trip.Scrap()
}


