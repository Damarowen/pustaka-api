package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", rootHandler)

	router.Run(":9090")

}


func rootHandler(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{
		"name": "damar",
		"usia": "17",
	})
}