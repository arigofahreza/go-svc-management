package routers

import (
	"go-svc-management/src/configs"
	"go-svc-management/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRouters() (*gin.Engine, *mongo.Client, *redis.Client) {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(CORSMiddleware())

	mongoCollection, err := configs.MongoClient()
	if err != nil {
		panic(err)
	}

	redisClient, err := configs.Redis()
	if err != nil {
		panic(err)
	}

	authRouter := controllers.InitAuthRouters(configs.MongoCollection)

	mainGroup := router.Group("/api/v1")
	{
		auth := mainGroup.Group("/auth")
		{
			auth.POST("/register", authRouter.RegisterRouter)
		}
	}
	return router, mongoCollection, redisClient
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
