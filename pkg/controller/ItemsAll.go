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
	Code int `code`
	Body struct {
		Category string `category_id`
		Pictures[] struct {
			Url string `url`
		} `pictures`
		Price float32 `price`
		Title string  `title`
		Id    string  `id`
	} `body`
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
	ch := make(chan Items)
	itemCollector(token, userid, ch, c)
}

func itemCollector(token, userid string, ch chan Items, c *gin.Context ){
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
	var res Items
	//fmt.Println(string(data))
	erro = json.Unmarshal(data, &res)
	if erro != nil {
		c.JSON(req.StatusCode, erro.Error())
		return
	}
	c.JSON(req.StatusCode, res)
	fmt.Println(res)
}