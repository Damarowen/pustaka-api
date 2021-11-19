package main

import (
	//"encoding/json"

	//	"log"
	"github.com/gin-gonic/gin"
	"pustaka-api/handler"
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

	router.GET("/", handler.RootHandler)
	router.GET("/books/:id", handler.BooksHandler)
	router.GET("/query", handler.QueryHandler)
	router.POST("/books", handler.BooksPostHandler)

	router.Run(":9090")

}
