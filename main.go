package main

import (
	"encoding/json"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
	"log"
)

type ItemObject struct {
	Item   string         `json:"item"`
	Name   string         `json:"name"`
	Stock  bool           `json:"stock,omitempty"`
	Prices []CountryPrice `json:"prices"`
}

type CountryPrice struct {
	CountryCode string    `json:"country_code"`
	Pricings    []Pricing `json:"price"`
}

type Pricing struct {
	Name      string   `json:"name"`
	Current   bool     `json:"current"`
	Currency  string   `json:"currency"`
	Item      int      `json:"item"`
	Shippings Shipping `json:"shipping"`
}

type Shipping struct {
	GB int `json:"gb"`
	FR int `json:"fr"`
}

func main() {
	r := gin.Default()
	items, _ := Items()

	r.GET("", func(c *gin.Context) {
		c.JSON(200, items)
	})

	r.GET("/:item", func(c *gin.Context) {
		param_item := c.Params.ByName("item")

		for pos, item := range items {
			if item.Item == param_item {
				c.JSON(200, item)
				break
			} else if (pos+1) == len(items) && item.Item != param_item {
				c.JSON(404, gin.H{"error": "[" + param_item + "] doesn't exist!"})
				c.Abort()
			}
		}
	})

	r.Run(":3000")
}

func Items() ([]ItemObject, error) {
	data, err := ioutil.ReadFile("items.json")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var items []ItemObject
	err = json.Unmarshal(data, &items)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return items, err
}
