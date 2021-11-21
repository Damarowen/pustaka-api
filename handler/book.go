package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	//"github.com/go-playground/validator/v10"
	"net/http"
	"pustaka-api/book"
	"pustaka-api/helper"
)

type BookHandlers struct {
	bookService book.Service
}

func NewBookHandler(bookService *book.Service) *BookHandlers {
	return &BookHandlers{*bookService}
}

func (h *BookHandlers) RootHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"name": "damar",
		"usia": "17",
	})
}
func (h *BookHandlers) GetAllBook(c *gin.Context) {

	allBook, err := h.bookService.FindAll()

	if err != nil {
		fmt.Println(err)
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data ":   allBook,
		"message": "Success",
	})

}
func (h *BookHandlers) GetById(c *gin.Context) {

	idString := c.Param("id")

	ID, _ := strconv.Atoi(idString)

	singleBook, err := h.bookService.FindById(ID)

	if err != nil {
		fmt.Println(err)
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data ":   singleBook,
		"message": "Success",
	})
}

func (h *BookHandlers) QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title ": title,
		"price":  price,
	})
}

func (h *BookHandlers) BooksPostHandler(c *gin.Context) {
	var bookDTO book.BookRequest

	errDTO := c.ShouldBind(&bookDTO)

	if errDTO != nil {
		fmt.Println(errDTO)
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	newBook, err := h.bookService.Create(bookDTO)

	if err != nil {
		fmt.Println(errDTO)
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data ":   newBook,
		"message": "Success",
	})

}
