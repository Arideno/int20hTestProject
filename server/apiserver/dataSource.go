package apiserver

import (
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

func atb(result func([]product)) {
	c := colly.NewCollector()

	names := make(map[int]string)
	prices := make(map[int]float64)
	images := make(map[int]string)
	urls := make(map[int]string)

	indexNames := 0
	indexPrice := 0
	indexImages := 0
	indexUrls := 0

	c.OnHTML("div[data-productkey]", func(e *colly.HTMLElement) {
		namesDOM := e.DOM.Find(".product-detail div")
		for _, name := range namesDOM.Nodes {
			if strings.Contains(strings.ToLower(name.FirstChild.Data), "крупа") {
				names[indexNames] = name.FirstChild.Data
			}
		}
		indexNames++
		pricesDOM := e.DOM.Find("p.simple-price span.price")
		for _, price := range pricesDOM.Nodes {
			if _, ok := names[indexPrice]; ok {
				prices[indexPrice], _ = strconv.ParseFloat(price.FirstChild.Data + "." + price.LastChild.FirstChild.Data, 64)
			}
		}
		indexPrice++
		imagesDOM := e.DOM.Find("img.bg-img")
		for _, image := range imagesDOM.Nodes {
			for _, attr := range image.Attr {
				if attr.Key == "src" {
					images[indexImages] = attr.Val
					break
				}
			}
		}
		indexImages++
		urlsDOM := e.DOM.Find(".front a")
		for _, url := range urlsDOM.Nodes {
			for _, attr := range url.Attr {
				if attr.Key == "href" {
					urls[indexUrls] = attr.Val
					break
				}
			}
		}
		indexUrls++
	})

	c.OnScraped(func(response *colly.Response) {
		products := make([]product, 0)

		for i := range names {
			products = append(products, product{
				Name:  names[i],
				Price: prices[i],
				Image: images[i],
				Url:   "https://zakaz.atbmarket.com" + urls[i],
				Store: "atb",
			})
		}

		result(products)
	})

	_ = c.Visit("https://zakaz.atbmarket.com/search/531?text=гречана")
}
