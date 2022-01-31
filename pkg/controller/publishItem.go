package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InBound struct {
	Description string `json:"description"`
	ItemPub     struct {
		Title             string `json:"title"`
		CategoryID        string `json:"category_id"`
		Price             int    `json:"price"`
		CurrencyID        string `json:"currency_id"`
		AvailableQuantity int    `json:"available_quantity"`
		BuyingMode        string `json:"buying_mode"`
		Condition         string `json:"condition"`
		ListingTypeID     string `json:"listing_type_id"`
		SaleTerms         []struct {
			ID        string `json:"id"`
			ValueName string `json:"value_name"`
		} `json:"sale_terms"`
		Pictures []struct {
			Source string `json:"source"`
		} `json:"pictures"`
		Attributes []struct {
			ID        string `json:"id"`
			ValueName string `json:"value_name"`
		} `json:"attributes"`
	}
}

func PublishItem(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")

	if c.Request.Method == "OPTIONS" {
		c.Writer.WriteHeader(http.StatusOK)
		return
	}

	token := c.Query("token")
	var url string = "https://api.mercadolibre.com/items"
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	var res InBound
	err = json.Unmarshal(data, &res)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	dat, err := json.Marshal(res.ItemPub)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dat))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	req.Header.Add("Authorization", " Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(resp.StatusCode, err.Error())
		return
	}
	dta, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	//Envío de la descripción

	var str struct {
		Id string `json:"id"`
	}
	err = json.Unmarshal(dta, &str)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	var send struct {
		PlainText string `json:"plain_text"`
	}

	send.PlainText = res.Description

	dat2, err := json.Marshal(send)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	fmt.Printf("%+v", string(dat2))

	req2, err := http.NewRequest("POST", url+"/"+str.Id+"/description", bytes.NewBuffer(dat2))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	fmt.Printf("%+v 2", string(url+"/"+str.Id+"/description"))
	req2.Header.Add("Authorization", " Bearer "+token)
	req2.Header.Add("Content-Type", "application/json")

	resp2, err := http.DefaultClient.Do(req2)
	if err != nil {
		c.JSON(resp2.StatusCode, err.Error())
		return
	}

	fmt.Printf("%+v 3", string(str.Id))

	dta2, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	fmt.Printf("%+v 4", string(str.Id))
	fmt.Printf("this is %+v", string(dta2))

	out := make(map[string]string)
	err = json.Unmarshal(dta2, &out)
	if err != nil && len(dta2) != 0 {
		c.JSON(500, err.Error())
		return
	}

	fmt.Printf("%+v 5", string(str.Id))
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	c.JSON(resp.StatusCode, out)

}
