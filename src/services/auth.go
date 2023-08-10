package services

import (
	"context"
	"go-svc-management/src/models"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthServices struct {
	Data           []*models.TokenModel `json:"data"`
	SessionService SessionServices
}

func InitAuthService() *AuthServices {
	return &AuthServices{}
}

func (service AuthServices) Login(c *gin.Context, redisClient *redis.Client, ctx context.Context, mongoCollection *mongo.Collection, loginModel models.LoginModel) ([]*models.TokenModel, error) {
	session, err := service.SessionService.GetSession(c, redisClient)
	if err != nil && session == "" {
		return nil, err
	}
	filter := bson.M{"username": loginModel.Username}
	var account models.AccountModel
	err = mongoCollection.FindOne(ctx, filter).Decode(&account)
	if err == mongo.ErrNoDocuments {
		return service.Data, nil
	} else if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(loginModel.Password))
	if err != nil {
		return nil, err
	}

}
