package router

import (
	"go-meli-integration/pkg/controller"
	"github.com/gin-gonic/gin"
)

func Run(){
	router := gin.Default()
	router.GET("/auth", controller.Auth)
	router.GET("/items/all", controller.ItemsAll)
	router.Run()
}
