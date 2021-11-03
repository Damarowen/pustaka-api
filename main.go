package main

import (
	"encoding/json"
	"fmt"
	//	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func main() {

	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/books/:id", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", booksPostHandler)

	router.Run(":9090")

}

func rootHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"name": "damar",
		"usia": "17",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id ": id,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title ": title,
		"price":  price,
	})
}

type Bookinput struct {
	Title string ` json:"title" binding:"required" `
	Price json.Number   ` json:"price" binding:"required,number" `
	//* dari depan sub_title
	//Subtitle string `json:"sub_title"`
}

func booksPostHandler(c *gin.Context) {

	var bookInput Bookinput

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {
		//log.Fatal(err)
		//log.Panic(err)
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error %s, error %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusInternalServerError, errorMessage)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"title ":    bookInput.Title,
		"price":     bookInput.Price,
		//"sub_title": bookInput.Subtitle,
		"message":   "Success",
	})
}
