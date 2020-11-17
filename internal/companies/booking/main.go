package booking

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/sergiorra/scraping-prices-go/internal/shared/price"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

const (
	URL = "https://www.booking.com/hotel/no/spitsbergen.en-gb.html"
)

func Scrap(wg *sync.WaitGroup) {

	var prices []int

	c := colly.NewCollector(
		colly.AllowedDomains("www.booking.com"),
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)

	c.OnHTML(".hprt-price-block", func(e *colly.HTMLElement) {
		priceWithCurrency := e.ChildText("span.prco-valign-middle-helper")
		priceWithCurrency = strings.ReplaceAll(priceWithCurrency, ",", "")
		price, err := strconv.Atoi(priceWithCurrency[5:])
		if err != nil {
			fmt.Println(err)
		}
		prices = append(prices, price)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(getUrl())
	c.Wait()

	fmt.Printf("The minimum price in Booking is %v\n", price.GetMinPrice(prices))
	wg.Done()
}

func getUrl() string {
	params := url.Values{}
	params.Set("no_rooms", "1")
	params.Set("checkin", "2021-02-03")
	params.Set("checkout", "2021-02-07")
	params.Set("group_adults", "2")
	params.Set("group_children", "0")
	params.Set("req_adults", "2")
	params.Set("req_children", "0")
	params.Set("dist", "0")
	params.Set("type", "total")
	params.Set("selected_currency", "NOK")

	return URL + "?" + params.Encode()
}
