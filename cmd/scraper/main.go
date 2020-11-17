package main

import (
	"fmt"
	"sync"

	"github.com/sergiorra/scraping-prices-go/internal/companies/booking"
	"github.com/sergiorra/scraping-prices-go/internal/companies/pricetravel"
	"github.com/sergiorra/scraping-prices-go/internal/companies/trip"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	go booking.Scrap(&wg)
	go pricetravel.Scrap(&wg)
	go trip.Scrap(&wg)

	wg.Wait()

	fmt.Println("Done!")
}


