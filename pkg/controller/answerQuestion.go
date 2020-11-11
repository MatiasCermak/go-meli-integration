package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type AnswerOut struct {
	Question_id string `json:"id"`
	Status string `json:"status"`
}

func AnswerQuestion(c *gin.Context)  {
	token := c.Query("token")
 	var url string = "https://api.mercadolibre.com/answers?access_token=" + token
	resp, err :=http.Post(url, "application/json", c.Request.Body)
 	if err !=nil{
 		c.JSON(resp.StatusCode,err.Error())
		return
	}
	data,err := ioutil.ReadAll(resp.Body)
	if err !=nil{
		c.JSON(500,err.Error())
		return
	}
	var res AnswerOut
	err = json.Unmarshal(data, &res)
	if err != nil{
		c.JSON(500,err.Error())
	}
	c.JSON(200,res)
}
