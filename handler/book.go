package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"
	"net/http"
	"pustaka-api/dto"
	"pustaka-api/helper"

)

func RootHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"name": "damar",
		"usia": "17",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id ": id,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title ": title,
		"price":  price,
	})
}

func BooksPostHandler(c *gin.Context) {
	var bookDTO dto.Bookinput

	errDTO := c.ShouldBind(&bookDTO)

	if errDTO != nil {
		fmt.Println(errDTO)
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title ": bookDTO.Title,
		"price":  bookDTO.Price,
		//"sub_title": bookInput.Subtitle,
		"message": "Success",
	})
}
