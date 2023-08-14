package controllers

import (
	"context"
	"go-svc-management/src/models"
	"go-svc-management/src/services"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRouters struct {
	MongoCollection *mongo.Collection
	AuthServices    services.AuthServices
}

func InitAuthRouters(mongoCollection *mongo.Collection) *AuthRouters {
	return &AuthRouters{
		MongoCollection: mongoCollection,
		AuthServices:    *services.InitAuthServices(&gin.Context{}, context.TODO(), mongoCollection),
	}
}

func (router AuthRouters) RegisterRouter(c *gin.Context) {
	start := time.Now()
	account := models.AccountView{}
	baseResponse := models.BaseResponse{}
	err := c.BindJSON(&account)
	if err != nil {
		baseResponse.Metadata = models.OverrideMetadata(http.StatusBadRequest, "FAILED", "fetch body error", time.Since(start).Seconds())
		c.AbortWithStatusJSON(400, baseResponse)
		return
	}
	datas, err := router.AuthServices.RegisterService(account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(201, datas)
			return
		}
		c.AbortWithStatusJSON(500, datas)
		return
	}
	c.JSON(400, datas)
}
