package trip

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
	URL = "https://th.trip.com/hotels/pattaya-hotel-detail-25371680/x2-pattaya-oceanphere/"
)

func Scrap(wg *sync.WaitGroup) {

	var prices []int

	c := colly.NewCollector(
		colly.AllowedDomains("th.trip.com"),
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)

	c.OnHTML(".salecardB-priceinfo", func(e *colly.HTMLElement) {
		priceWithCurrency := e.ChildText("div.price-display")
		priceWithCurrency = strings.ReplaceAll(priceWithCurrency, ",", "")
		fmt.Println(priceWithCurrency)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(getUrl())
	c.Wait()

	fmt.Printf("The minimum price in Trip is %v\n", price.GetMinPrice(prices))
	wg.Done()
}

func getUrl() string {
	params := url.Values{}
	params.Set("curr", "THB")
	params.Set("checkin", "2021-02-24")
	params.Set("checkout", "2021-02-26")
	params.Set("RoomQuantity", "1")
	params.Set("adult", "1")
	params.Set("children", "0")
	params.Set("ages", "")

	return URL + "?" + params.Encode()
}
