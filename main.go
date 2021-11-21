package main

import (
	"pustaka-api/book"
	"pustaka-api/config"
	"pustaka-api/routes"
)

func main() {

	db, _ := config.ConnectDatabase()

	bookRepository := book.NewRepository(db.DbSQL)
	bookService := book.NewService(bookRepository)
	
	r := Routes.SetupRouter(bookService)

	r.Run("127.0.0.1:9090")

}
