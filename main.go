package main

import (
	"pustaka-api/book"
	"pustaka-api/models"
	"pustaka-api/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	db, _ := models.ConnectDatabase()

	bookRepository := book.NewRepository(db.DbSQL)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	v1 := router.Group("/v1")

	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/books", bookHandler.GetAllBookHandler)
	v1.GET("/books/:id", bookHandler.GetByIdHandler)
	v1.POST("/books", bookHandler.PostBookHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	//v1.GET("/query", bookHandler.QueryHandler)

	router.Run("127.0.0.1:9090")

}
