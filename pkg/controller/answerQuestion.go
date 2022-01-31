package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AnswerQuestion(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Access-Control-Allow-Origin")
	if c.Request.Method == "OPTIONS" {
		c.Writer.WriteHeader(http.StatusOK)
		return
	}

	token := c.Query("token")
	var url string = "https://api.mercadolibre.com/answers?access_token=" + token
	resp, err := http.Post(url, "application/json", c.Request.Body)
	if err != nil {
		c.JSON(resp.StatusCode, err.Error())
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	res := make(AnswerOut)
	err = json.Unmarshal(data, &res)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(resp.StatusCode, res)
}

/*
func AnswerQuestion(c *gin.Context) {
	token := c.Query("token")
	var url string = "https://api.mercadolibre.com/answers?access_token=" + token
	resp, err := http.Post(url, "application/json", c.Request.Body)
	if err != nil {
		c.JSON(resp.StatusCode, err.Error())
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(resp.StatusCode, data)
}
*/
