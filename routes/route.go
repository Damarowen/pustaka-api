package Routes

import (
	"pustaka-api/book"
	"pustaka-api/config"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

var (
	db, _ = config.ConnectDatabase()
	bookRepository = book.NewBookRepository(db.DbSQL)
	bookService = book.NewBookService(bookRepository)
	bookHandler = handler.NewBookHandler(bookService)
)

//SetupRouter ... Configure routes
func SetupRouter(db *config.DbConn) *gin.Engine {

	r := gin.Default()
	
	v1 := r.Group("/v1")
	
	{
		v1.GET("/", bookHandler.RootHandler)
		v1.GET("/books", bookHandler.GetAllBookHandler)
		v1.GET("/books/:id", bookHandler.GetByIdHandler)
		v1.POST("/books", bookHandler.PostBookHandler)
		v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
		v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)
		//v1.GET("/query", bookHandler.QueryHandler)
	}

	return r
}
