package services

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"go-svc-management/src/models"
	"go-svc-management/src/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServices struct {
	C               *gin.Context
	Ctx             context.Context
	MongoCollection *mongo.Collection
}

func InitAuthServices(c *gin.Context, ctx context.Context, mongoCollection *mongo.Collection) *AuthServices {
	return &AuthServices{
		C:               c,
		Ctx:             ctx,
		MongoCollection: mongoCollection,
	}
}

func (service AuthServices) RegisterService(accountView models.AccountView) (*models.BaseResponse, error) {
	start := time.Now()
	datas := []*models.AccountView{}
	resp := models.BaseResponse{}
	hashId, err := bson.Marshal(accountView)
	if err != nil {
		resp.Metadata = models.OverrideMetadata(http.StatusInternalServerError, "FAILED", "hash id error", time.Since(start).Seconds())
		resp.Data = []interface{}{}
		return &resp, err
	}
	hash := md5.Sum(hashId)
	accountView.Id = hex.EncodeToString(hash[:])
	accountView.Password, err = utils.HashPassword(accountView.Password)
	if err != nil {
		resp.Metadata = models.OverrideMetadata(http.StatusInternalServerError, "FAILED", "hash password error", time.Since(start).Seconds())
		resp.Data = []interface{}{}
		return &resp, err
	}
	account := models.AccountView{}
	err = service.MongoCollection.FindOne(service.Ctx, bson.M{"email": accountView.Email}).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			_, err = service.MongoCollection.InsertOne(service.Ctx, accountView)
			if err != nil {
				resp.Metadata = models.OverrideMetadata(http.StatusInternalServerError, "FAILED", "insert data error", time.Since(start).Seconds())
				resp.Data = []interface{}{}
				return &resp, err
			}
			resp.Metadata = models.OverrideMetadata(http.StatusCreated, "CREATED", "success created data", time.Since(start).Seconds())
			resp.Data = append(datas, &accountView)
			return &resp, err
		}
		resp.Metadata = models.OverrideMetadata(http.StatusInternalServerError, "FAILED", "db mongo error", time.Since(start).Seconds())
		resp.Data = []interface{}{}
		return &resp, err
	}
	resp.Metadata = models.OverrideMetadata(http.StatusBadRequest, "FAILED", "email already used", time.Since(start).Seconds())
	resp.Data = []interface{}{}
	return &resp, err
}

func (service AuthServices) Login() {
	
}
