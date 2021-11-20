package main

import (
	//"encoding/json"

	//	"log"
	//	"log"
	//	"pustaka-api/book"
	"log"
	"pustaka-api/DB"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
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
	err := DB.NewConnDb()

	if err != nil {
		log.Fatal(err.Error())
	}
	
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.BooksPostHandler)

	router.Run(":9090")

}
