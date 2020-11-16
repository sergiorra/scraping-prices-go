package booking

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/url"
	"strconv"
	"strings"
)

const (
	URL = "https://www.booking.com/hotel/no/spitsbergen.en-gb.html"
)

func Scrap() {

	var prices []int

	c := colly.NewCollector(
		colly.AllowedDomains("www.booking.com"),
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)

	c.OnHTML(".hprt-price-block", func(e *colly.HTMLElement) {
		price := e.ChildText("span.prco-valign-middle-helper")
		price = strings.ReplaceAll(price, ",", "")
		pr, err := strconv.Atoi(price[5:])
		if err != nil {
			fmt.Println(err)
		}
		prices = append(prices, pr)
	})

	c.Visit(getUrl())
	c.Wait()

	var minPrice int
	for i, price := range prices {
		if i==0 || price < minPrice {
			minPrice = price
		}
	}
	fmt.Println(prices)
	fmt.Printf("The minimum price is %d\n", minPrice)
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
