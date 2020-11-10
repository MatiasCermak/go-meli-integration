package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func getAndMarshall(url string, res interface{}, c *gin.Context)  {
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
	erro = json.Unmarshal(data, &res)
	if erro != nil {
		c.JSON(req.StatusCode, erro.Error())
		return
	}
}