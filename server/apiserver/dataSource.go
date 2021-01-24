package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/gocolly/colly"
	"log"
	"net/http"
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

func Silpo()  []product {
	data := map[string]interface{}{ "customFilter": "гречана крупа", "filialId": 2382, "From":1}
	values := map[string]interface{}{ "method":"GetSimpleCatalogItems", "data": data}
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Panic(err)
	}
	resp, err := http.Post("https://api.catalog.ecom.silpo.ua/api/2.0/exec/EcomCatalogGlobal", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Panic(err)
	}
	var res map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Panic(err)
	}
	allItems := res["items"].([]interface{})

	products := make([]product, 0)
	for index := range allItems {
		item := allItems[index].(map[string]interface{})
		products = append(products, product{
			Name:  item["name"].(string),
			Price: item["price"].(float64),
			Image: item["mainImage"].(string),
			Url:   "https://shop.silpo.ua/detail/" + strconv.Itoa(int(item["id"].(float64))),
			Store: "silpo",
		})
	}
	return products
}

func Auchan() []product {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://auchan.ua/graphql/?query=query%20($searchQuery:%20String!,%20$currentPage:%20Int!,%20$pageSize:%20Int!,%20$category:%20[Int],%20$max_price:%20Float,%20$min_price:%20Float,%20$sort:%20sortInput,%20$filter:%20[attributeInput])%20{%20%20GetProductSearchResults(search_query:%20$searchQuery,%20current_page:%20$currentPage,%20page_size:%20$pageSize,%20category_ids:%20$category,%20min_price:%20$min_price,%20max_price:%20$max_price,%20sort:%20$sort,%20attributes:%20$filter)%20{%20%20%20%20products%20{%20%20%20%20%20%20...ProductBase%20%20%20%20%20%20__typename%20%20%20%20}%20%20%20%20__typename%20%20}}fragment%20ProductBase%20on%20ProductInterface%20{%20%20id%20%20sku%20%20name%20%20url_key%20%20type_id%20%20stock_status%20%20thumbnail%20{%20%20%20%20url%20%20%20%20__typename%20%20}%20%20special_price%20%20special_to_date%20%20price%20{%20%20%20%20regularPrice%20{%20%20%20%20%20%20amount%20{%20%20%20%20%20%20%20%20currency%20%20%20%20%20%20%20%20value%20%20%20%20%20%20%20%20__typename%20%20%20%20%20%20}%20%20%20%20%20%20__typename%20%20%20%20}%20%20%20%20__typename%20%20}%20%20offers%20{%20%20%20%20from_date%20%20%20%20to_date%20%20%20%20__typename%20%20}%20%20attributes%20{%20%20%20%20code%20%20%20%20label%20%20%20%20value%20%20%20%20__typename%20%20}%20%20__typename}&variables={%22searchQuery%22:%22%D0%B3%D1%80%D0%B5%D1%87%D0%B0%D0%BD%D0%B0%20%D0%BA%D1%80%D1%83%D0%BF%D0%B0%22,%22pageSize%22:24,%22currentPage%22:1,%22category%22:[]}", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("store", "ua")
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	var res map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Panic(err)
	}
	allItems := res["data"].(map[string]interface{})["GetProductSearchResults"].(map[string]interface{})["products"].([]interface{})
	products := make([]product, 0)
	for index := range allItems {
		item := allItems[index].(map[string]interface{})
		products = append(products, product{
			Name:  item["name"].(string),
			Price: item["price"].(map[string]interface{})["regularPrice"].(map[string]interface{})["amount"].(map[string]interface{})["value"].(float64),
			Image: item["thumbnail"].(map[string]interface{})["url"].(string),
			Url:   "https://auchan.ua/ua/" + item["url_key"].(string) + "-" + strconv.Itoa(int(item["id"].(float64))),
			Store: "auchan",
		})
	}
	return products
}
