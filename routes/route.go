package Routes

import (
	"pustaka-api/JWT"
	"pustaka-api/auth"
	"pustaka-api/book"
	"pustaka-api/config"
	"pustaka-api/handler"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
)

var (
	db, _          = config.ConnectDatabase()
	bookRepository = book.NewBookRepository(db.DbSQL)
	bookService    = book.NewBookService(bookRepository)
	bookHandler    = handler.NewBookHandler(bookService)


	userRepository user.IUserRepository = user.NewUserRepository(db.DbSQL)
	jwtService     JWT.IJwtService        = JWT.NewJWTService()
	authService    auth.IAuthService       = auth.NewAuthService(userRepository)
	authController handler.IAuthController = handler.NewAuthController(authService, jwtService)
)

//SetupRouter ... Configure routes
func SetupRouter(db *config.DbConn) *gin.Engine {

	r := gin.Default()

	authRoutes := r.Group("/v1/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	booksRoutes := r.Group("/v1/books")
	{
		//booksRoutes.GET("/", bookHandler.RootHandler)
		booksRoutes.GET("/", bookHandler.GetAllBookHandler)
		booksRoutes.GET("/:id", bookHandler.GetByIdHandler)
		booksRoutes.POST("/", bookHandler.PostBookHandler)
		booksRoutes.PUT("/:id", bookHandler.UpdateBookHandler)
		booksRoutes.DELETE("/:id", bookHandler.DeleteBookHandler)
		//v1.GET("/query", bookHandler.QueryHandler)
	}

	return r
}
