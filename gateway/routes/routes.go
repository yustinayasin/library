package routes

import (
	"log"
	"os"

	protoBooks "shared/proto/books"
	protoUsers "shared/proto/users"

	middleware "shared/app/middlewares"

	handlerBook "gateway/routes/handler/books"
	handlerBooksStocks "gateway/routes/handler/bookstocks"
	handlerBorrowRecord "gateway/routes/handler/borrowrecords"
	handlerUser "gateway/routes/handler/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouteRegister(JWTConfig middleware.ConfigJWT, userClient protoUsers.UserServiceClient, bookClient protoBooks.BookServiceClient, bookStockClient protoBooks.BooksStocksServiceClient, borrowRecordClient protoUsers.BorrowRecordsServiceClient) {
	router := gin.Default()

	router.Use(middleware.CacheControl(86400))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://poppy-florist.yustinayasin.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	user := router.Group("/user")
	{
		user.POST("/signup", handlerUser.SignupHandler(userClient))
		user.POST("/login", handlerUser.LoginHandler(userClient))
		user.PUT("/:userId", middleware.RequireAuth(handlerUser.EditUserHandler(userClient), JWTConfig, userClient))
		user.DELETE("/:userId", middleware.RequireAuth(handlerUser.DeleteUserHandler(userClient), JWTConfig, userClient))
		user.GET("/:userId", middleware.RequireAuth(handlerUser.GetUserHandler(userClient), JWTConfig, userClient))
	}

	book := router.Group("/book")
	{
		book.POST("/", middleware.RequireAuthAdmin(handlerBook.AddBook(bookClient), JWTConfig, userClient))
		book.PUT("/:bookId", middleware.RequireAuthAdmin(handlerBook.EditBook(bookClient), JWTConfig, userClient))
		book.DELETE("/:bookId", middleware.RequireAuthAdmin(handlerBook.DeleteBook(bookClient), JWTConfig, userClient))
		book.GET("/:bookId", middleware.RequireAuth(handlerBook.GetBook(bookClient), JWTConfig, userClient))
	}

	bookstock := router.Group("/bookstock")
	{
		bookstock.POST("/", middleware.RequireAuthAdmin(handlerBooksStocks.AddBooksStocks(bookStockClient, bookClient), JWTConfig, userClient))
		bookstock.PUT("/:bookStockId", middleware.RequireAuthAdmin(handlerBooksStocks.EditBooksStocks(bookStockClient), JWTConfig, userClient))
		bookstock.DELETE("/:bookStockId", middleware.RequireAuthAdmin(handlerBooksStocks.DeleteBooksStocks(bookStockClient), JWTConfig, userClient))
		bookstock.GET("/:bookStockId", middleware.RequireAuthAdmin(handlerBooksStocks.GetBooksStocks(bookStockClient), JWTConfig, userClient))
	}

	borrowrecord := router.Group("/borrowrecord")
	{
		borrowrecord.POST("/", middleware.RequireAuthAdmin(handlerBorrowRecord.AddBorrowRecords(borrowRecordClient, bookClient), JWTConfig, userClient))
		borrowrecord.PUT("/:borrowRecordId", middleware.RequireAuthAdmin(handlerBorrowRecord.EditBorrowRecords(borrowRecordClient), JWTConfig, userClient))
		borrowrecord.DELETE("/:borrowRecordId", middleware.RequireAuthAdmin(handlerBorrowRecord.DeleteBorrowRecords(borrowRecordClient), JWTConfig, userClient))
		borrowrecord.GET("/:borrowRecordId", middleware.RequireAuthAdmin(handlerBorrowRecord.GetBorrowRecords(borrowRecordClient), JWTConfig, userClient))
	}

	port := ":8080"
	if err := router.Run(port); err != nil {
		log.Println("Failed to start server:", err)
		os.Exit(1)
	}
}
