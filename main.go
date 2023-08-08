package main

import (
	"context"
	"go-svc-management/src/configs"
	"go-svc-management/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(CORSMiddleware())

	mongoCollection, err := configs.MongoClient()
	if err != nil {
		panic(err)
	}
	defer mongoCollection.Disconnect(context.TODO())
	accountController := controllers.InitAccountControllers(configs.MongoCollection)

	redisClient, err := configs.Redis()
	if err != nil {
		panic(err)
	}
	defer redisClient.Close()
	sessionController := controllers.InitSessionControllers(redisClient)

	mainGroup := router.Group("/api/v1")
	{
		account := mainGroup.Group("/account")
		{
			account.POST("", accountController.CreateAccountControllers)
			account.GET("", accountController.GetAllAccountControllers)
			account.GET("/:id", accountController.GetAccountByIdController)
			account.PUT("", accountController.UpdateAccountController)
			account.DELETE("/:id", accountController.DeleteAccountByIdController)
		}

		session := mainGroup.Group("/session")
		{
			session.POST("", sessionController.GenerateSessionController)
		}
	}
	router.Run(":6004")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
