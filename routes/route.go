package Routes

import (
	"pustaka-api/JWT"
	"pustaka-api/auth"
	"pustaka-api/book"
	"pustaka-api/config"
	"pustaka-api/handler"
	"pustaka-api/middleware"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
)

var (
	db, _          = config.ConnectDatabase()
	bookRepository = book.NewBookRepository(db.DbSQL)
	bookService    = book.NewBookService(bookRepository)
	bookController    = handler.NewBookHandler(bookService)


	userRepository user.IUserRepository = user.NewUserRepository(db.DbSQL)
	jwtService     JWT.IJwtService        = JWT.NewJWTService()
	authService    auth.IAuthService       = auth.NewAuthService(userRepository)
	authController handler.IAuthController = handler.NewAuthController(authService, jwtService)

	userService    user.IUserService       = user.NewUserService(userRepository)
	userController handler.IUserController = handler.NewUserController(userService, jwtService)

)

//SetupRouter ... Configure routes
func SetupRouter(db *config.DbConn) *gin.Engine {

	r := gin.Default()

	authRoutes := r.Group("/v1/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("v1/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	booksRoutes := r.Group("/v1/books")
	{
		//booksRoutes.GET("/", bookHandler.RootHandler)
		booksRoutes.GET("/", bookController.GetAllBookHandler)
		booksRoutes.GET("/:id", bookController.GetByIdHandler)
		booksRoutes.POST("/", bookController.PostBookHandler)
		booksRoutes.PUT("/:id", bookController.UpdateBookHandler)
		booksRoutes.DELETE("/:id", bookController.DeleteBookHandler)
		//v1.GET("/query", bookHandler.QueryHandler)
	}

	return r
}
