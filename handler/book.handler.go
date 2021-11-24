package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	//"github.com/go-playground/validator/v10"
	"net/http"
	"pustaka-api/JWT"
	"pustaka-api/book"
	"pustaka-api/dto"
	"pustaka-api/helper"
	"pustaka-api/middleware"
	"pustaka-api/models"
)

type IBookHandlers interface {
	RootHandler(context *gin.Context)
	GetAllBookHandler(context *gin.Context)
	GetByIdHandler(context *gin.Context)
	UpdateBookHandler(context *gin.Context)
	DeleteBookHandler(context *gin.Context)
	PostBookHandler(context *gin.Context)
	QueryBookHandler(context *gin.Context)
}

type BookHandlers struct {
	bookService book.Iservice
	jwtService  JWT.IJwtService
}

func NewBookHandler(bookServ book.Iservice, jwtServ JWT.IJwtService) IBookHandlers {
	return &BookHandlers{bookService: bookServ, jwtService: jwtServ}
}

func (h *BookHandlers) RootHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"name": "damar",
		"usia": "17",
	})
}

func (h *BookHandlers) GetAllBookHandler(c *gin.Context) {

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

func (h *BookHandlers) GetByIdHandler(c *gin.Context) {

	idString := c.Param("id")

	id, err := strconv.ParseUint(idString, 10, 32)

	singleBook, err := h.bookService.FindById(uint(id))

	if (singleBook == models.Book{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		c.JSON(http.StatusNotFound, res)
		return
	}

	if err != nil {
		fmt.Println(err)
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	//	res := helper.MappingResponse(singleBook)

	response := helper.BuildResponse(true, "SUCCESS", singleBook)
	c.JSON(http.StatusCreated, response)
}

func (h *BookHandlers) UpdateBookHandler(c *gin.Context) {

	var bookDTO dto.BookUpdateDTO
	errDTO := c.ShouldBind(&bookDTO)

	if errDTO != nil {
		fmt.Println(errDTO)
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	idString := c.Param("id")
	userID := middleware.GetCurrentUser(c.Request.Context())
	tes , _:= strconv.ParseUint(userID.(string), 10, 32)
	id, err := strconv.ParseUint(idString, 10, 32)

	bookDTO.ID = uint(id)
	bookDTO.CreatedAt = time.Now()
	bookDTO.UserID = tes
	ok, err := h.bookService.IsAllowedToEdit(userID.(string), bookDTO.ID)

	//* if id not found
	if err != nil {
		response := helper.BuildErrorResponse("You dont have permission", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusForbidden, response)
		return
	}

	//* if book not belonged to user
	if !ok && err == nil {
		response := helper.BuildErrorResponse("You dont have permission", " You are not the owner", helper.EmptyObj{})
		c.JSON(http.StatusForbidden, response)
		return
	}

	b, _ := h.bookService.Update(bookDTO)
	res := helper.BuildResponse(true, "SUCCESS", b)
	c.JSON(http.StatusOK, res)
}

func (h *BookHandlers) DeleteBookHandler(c *gin.Context) {
	var book models.Book

	idString := c.Param("id")
	userID := middleware.GetCurrentUser(c.Request.Context())

	id, _ := strconv.ParseUint(idString, 0, 0)
	book.ID = uint(id)

	ok, err := h.bookService.IsAllowedToEdit(userID.(string), book.ID)

	//* if id not found
	if err != nil {
		response := helper.BuildErrorResponse("You dont have permission", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusForbidden, response)
		return
	}

	//* if book not belonged to user
	if !ok && err == nil {
		response := helper.BuildErrorResponse("You dont have permission", " You are not the owner", helper.EmptyObj{})
		c.JSON(http.StatusForbidden, response)
		return
	}

	h.bookService.Delete(book)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	c.JSON(http.StatusOK, res)

}

func (h *BookHandlers) PostBookHandler(c *gin.Context) {
	var bookDTO dto.BookRequest

	userID := middleware.GetCurrentUser(c.Request.Context())

	convertedUserID, _ := strconv.ParseUint(userID.(string), 10, 64)

	errDTO := c.ShouldBind(&bookDTO)

	if errDTO != nil {
		fmt.Println(errDTO)
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	//* masukin user id ke dto
	bookDTO.UserID = convertedUserID

	newBook, err := h.bookService.Create(bookDTO)

	if err != nil {
		fmt.Println(errDTO)
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "SUCCESS", newBook)
	c.JSON(http.StatusCreated, response)

}

func (h *BookHandlers) QueryBookHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title ": title,
		"price":  price,
	})
}
