package controller

import (
	"github.com/gin-gonic/gin"
)

type ItemIds struct {
	Results[] string `results`
}

type Questions[] struct {
	Date_created string `json:"date_created"`
	Item_id       string `json:"item_id"`
	Status       string `json:"status"`
	Text         string `json:"text"`
	Id           int64  `json:"id"`
	Answer       string `json:"answer"`
}

type Question struct {
	Questn Questions `json:"questions""`
}


type Items struct{
	Body struct {
		Id    string
		Title string
		Price float32
		Pictures[] map[string]string
		Available_quantity int
		Sold_quantity int
	}

}

type Item struct{
	Id    string
	Title string
	Price float32
	Quantity int
	SoldQuantity int
	Picture string
	Question Questions
}


func ItemsAll(c *gin.Context) {
	token := c.Query("token")
	userid := c.Query("userid")
	var url string = "https://api.mercadolibre.com/users/" + userid + "/items/search?access_token=" + token
	var res ItemIds
	getAndMarshall(url, &res, c)
	ch1 := make(chan Item)
	ch2 := make(chan Question)
	var response[] Item
	if len(res.Results) == 0 {
		c.JSON(400, struct {
			Error string}{
			Error: "No items being sold by this user",
		})
		return
	}
	for i := 0; i < len(res.Results); i++ {
		go itemCollector(token, res.Results[i], ch1, c)
		go questionCollector(token, res.Results[i], ch2, c)
		respContainer := <-ch1
		respContainer.Question = (<-ch2).Questn
		response = append(response, respContainer)
	}
	c.JSON(200, response)
}

func itemCollector(token, itemid string, ch1 chan Item, c *gin.Context){
	var url string = "https://api.mercadolibre.com/items?ids="+ itemid +"&attributes=id,price,available_quantity,title,pictures,sold_quantity&access_token=" + token
	var res[] Items
	getAndMarshall(url, &res, c)
	var resp Item
	resp.Quantity = res[0].Body.Available_quantity
	resp.SoldQuantity = res[0].Body.Sold_quantity
	resp.Id = res[0].Body.Id
	resp.Picture = res[0].Body.Pictures[0]["url"]
	resp.Price = res[0].Body.Price
	resp.Title = res[0].Body.Title
	ch1 <- resp
}

func questionCollector(token, itemid string, ch2 chan Question, c *gin.Context)  {
	var url string = "https://api.mercadolibre.com/questions/search?item="+ itemid +"&access_token=" + token
	var res Question
	getAndMarshall(url, &res, c)
	var unansweredQ Question
	if len(res.Questn) == 0 {
		ch2 <- Question{}
	}
	for i := len(res.Questn)-1; i >= 0 ; i-- {
		if res.Questn[i].Status == "UNANSWERED" {
			unansweredQ.Questn = append(unansweredQ.Questn, res.Questn[i])
		}
	}
	ch2 <- unansweredQ
}