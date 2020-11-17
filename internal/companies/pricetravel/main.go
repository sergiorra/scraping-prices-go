package pricetravel

import (
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/sergiorra/scraping-prices-go/internal/shared/price"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

const (
	URL = "https://www.pricetravel.com/hotel/occidental-at-xcaret-destination.xcaret-quintana-roo-mexico"
)

func Scrap(wg *sync.WaitGroup) {

	var prices []int

	c := colly.NewCollector(
		colly.AllowedDomains("www.pricetravel.com"),
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)

	c.OnHTML(".room-table-price", func(e *colly.HTMLElement) {
		priceWithCurrency := e.ChildText("h3.product-price-final")
		priceWithCurrency = strings.ReplaceAll(priceWithCurrency, ",", "")
		fmt.Println(priceWithCurrency)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(getUrl())
	c.Wait()

	fmt.Printf("The minimum price in PriceTravel is %v\n", price.GetMinPrice(prices))
	wg.Done()
}

func getUrl() string {
	params := url.Values{}
	params.Set("checkin", "2021-03-14")
	params.Set("checkout", "2021-03-23")
	params.Set("rooms", "1")
	params.Set("room1.adults", "2")
	params.Set("room1.kids", "1")
	params.Set("room1.agekids", "10")

	return URL + "?" + params.Encode()
}