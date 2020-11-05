package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type ItemIds struct {
	Results[] string `results`
}

type Items struct{
	Body struct {
		Id    string  `id`
		Title string  `title`
		Price float32 `price`
		Pictures[] map[string]string `pictures`

	} `body`
	Category string `category_id`
}

type Item struct{
	Id    string
	Title string
	Price float32
	Category string
	Picture string
}


func ItemsAll(c *gin.Context) {
	token := c.Query("token")
	userid := c.Query("userid")
	var url string = "https://api.mercadolibre.com/users/" + userid + "/items/search?access_token=" + token
	req, erro := http.Get(url)
	if req.StatusCode != 200 || erro != nil{
		c.JSON(req.StatusCode, erro.Error())
		return
	}
	defer req.Body.Close()
	data, erro := ioutil.ReadAll(req.Body)
	if erro != nil {
		c.JSON(req.StatusCode, erro.Error())
		return
	}
	var res ItemIds
	//fmt.Println(string(data))
	erro = json.Unmarshal(data, &res)
	if erro != nil {
		c.JSON(500, erro.Error())
		return
	}
	ch := make(chan Item)
	var response[] Item

	for i := 0; i < len(res.Results); i++ {
		go itemCollector(token, res.Results[i], ch, c)
	}
	for i := 0; i < len(res.Results); i++ {
		response = append(response, <-ch)
	}
	c.JSON(200, response)

}

func itemCollector(token, userid string, ch chan Item, c *gin.Context ){
	var url string = "https://api.mercadolibre.com/items?ids="+ userid +"&attributes=id,price,category_id,title,pictures&access_token=" + token
	req, erro := http.Get(url)
	if req.StatusCode != 200 || erro != nil{
		c.JSON(req.StatusCode, erro.Error())
		return
	}
	defer req.Body.Close()
	data, erro := ioutil.ReadAll(req.Body)
	if erro != nil {
		c.JSON(req.StatusCode, erro.Error())
		return
	}
	var res[] Items
	fmt.Println(string(data))
	erro = json.Unmarshal(data, &res)
	if erro != nil {
		c.JSON(req.StatusCode, erro.Error())
		return
	}
	var resp Item
	resp.Category = res[0].Category
	resp.Id = res[0].Body.Id
	resp.Picture = res[0].Body.Pictures[0]["url"]
	resp.Price = res[0].Body.Price
	resp.Title = res[0].Body.Title
	ch <- resp
}