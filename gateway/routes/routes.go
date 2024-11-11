package routes

import (
	"log"
	"os"
	"shared/proto/users"

	middleware "shared/app/middlewares"

	"gateway/routes/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouteRegister(JWTConfig middleware.ConfigJWT, userClient proto.UserServiceClient) {
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
		user.POST("/signup", handler.SignupHandler(userClient))
		user.POST("/login", handler.LoginHandler(userClient))
		user.PUT("/:userId", middleware.RequireAuth(handler.EditUserHandler(userClient), JWTConfig, userClient))
		user.DELETE("/:userId", middleware.RequireAuth(handler.DeleteUserHandler(userClient), JWTConfig, userClient))
		user.GET("/:userId", middleware.RequireAuth(handler.GetUserHandler(userClient), JWTConfig, userClient))
	}

	port := ":8080"
	if err := router.Run(port); err != nil {
		log.Println("Failed to start server:", err)
		os.Exit(1)
	}
}
