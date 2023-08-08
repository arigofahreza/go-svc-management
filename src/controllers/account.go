package controllers

import (
	"context"
	"go-svc-management/src/models"
	"go-svc-management/src/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountControllers struct {
	MongoCollection *mongo.Collection
	AccountServices *services.AccountServices
}

func InitAccountControllers(mongoCollection *mongo.Collection) *AccountControllers {
	return &AccountControllers{
		MongoCollection: mongoCollection,
		AccountServices: services.InitUserService(),
	}
}

func (controller AccountControllers) CreateAccountControllers(c *gin.Context) {
	start := time.Now()
	var baseResponse models.BaseResponse
	account := models.AccountModel{}
	err := c.BindJSON(&account)
	if err != nil {
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusInternalServerError, "FAILED", "parsing body error", time.Since(start).Seconds())
		c.AbortWithStatusJSON(http.StatusInternalServerError, baseResponse)
		return
	}
	datas, err := controller.AccountServices.CreateAccountService(context.TODO(), controller.MongoCollection, account)
	if err != nil {
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusInternalServerError, "FAILED", "insert data error", time.Since(start).Seconds())
		c.AbortWithStatusJSON(http.StatusInternalServerError, baseResponse)
		return
	}
	baseResponse.Data = datas
	baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusOK, "SUCCESS", "insert data success", time.Since(start).Seconds())
	c.JSON(http.StatusOK, baseResponse)
}

func (controller AccountControllers) GetAllAccountControllers(c *gin.Context) {
	start := time.Now()
	var baseResponse models.BaseResponse
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusBadRequest, "FAILED", "page required", time.Since(start).Seconds())
		c.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
		return
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusBadRequest, "FAILED", "size required", time.Since(start).Seconds())
		c.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
		return
	}
	datas, err := controller.AccountServices.GetAllAccountService(context.TODO(), controller.MongoCollection, page, size)
	if err != nil {
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusInternalServerError, "FAILED", "get data error", time.Since(start).Seconds())
		c.AbortWithStatusJSON(http.StatusInternalServerError, baseResponse)
		return
	}
	baseResponse.Data = datas
	baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusOK, "SUCCESS", "get data success", time.Since(start).Seconds())
	c.JSON(http.StatusOK, baseResponse)
}

func (controller AccountControllers) GetAccountByIdController(c *gin.Context) {
	start := time.Now()
	var baseResponse models.BaseResponse
	baseResponse.Data = []interface{}{}
	id := c.Param("id")
	if id == "" {
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusBadRequest, "FAILED", "id required", time.Since(start).Seconds())
		c.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
		return
	}
	datas, err := controller.AccountServices.GetAccountByIdService(context.TODO(), controller.MongoCollection, id)
	if err != nil {
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusInternalServerError, "FAILED", "get data error", time.Since(start).Seconds())
		c.AbortWithStatusJSON(http.StatusInternalServerError, baseResponse)
		return
	}
	if len(datas) > 0 {
		baseResponse.Data = datas
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusOK, "SUCCESS", "get data success", time.Since(start).Seconds())
		c.JSON(http.StatusOK, baseResponse)
	}
	baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusNotFound, "FAILED", "data not found", time.Since(start).Seconds())
	c.AbortWithStatusJSON(http.StatusNotFound, baseResponse)
}

func (controller AccountControllers) UpdateAccountController(c *gin.Context) {
	start := time.Now()
	var baseResponse models.BaseResponse
	baseResponse.Data = []interface{}{}
	account := models.AccountView{}
	err := c.BindJSON(&account)
	if err != nil {
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusInternalServerError, "FAILED", "parsing body error", time.Since(start).Seconds())
		c.AbortWithStatusJSON(http.StatusInternalServerError, baseResponse)
		return
	}
	datas, err := controller.AccountServices.UpdateAccountService(context.TODO(), controller.MongoCollection, account)
	if err != nil {
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusInternalServerError, "FAILED", "update data error", time.Since(start).Seconds())
		c.AbortWithStatusJSON(http.StatusInternalServerError, baseResponse)
		return
	}
	if len(datas) > 0 {
		baseResponse.Data = datas
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusOK, "SUCCESS", "update data success", time.Since(start).Seconds())
		c.JSON(http.StatusOK, baseResponse)
	}
	baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusNotFound, "FAILED", "data not found", time.Since(start).Seconds())
	c.AbortWithStatusJSON(http.StatusNotFound, baseResponse)
}

func (controller AccountControllers) DeleteAccountByIdController(c *gin.Context) {
	start := time.Now()
	var baseResponse models.BaseResponse
	baseResponse.Data = []interface{}{}
	id := c.Param("id")
	if id == "" {
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusBadRequest, "FAILED", "id required", time.Since(start).Seconds())
		c.AbortWithStatusJSON(http.StatusBadRequest, baseResponse)
		return
	}
	datas, err := controller.AccountServices.DeleteAccountService(context.TODO(), controller.MongoCollection, id)
	if err != nil {
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusInternalServerError, "FAILED", "delete data error", time.Since(start).Seconds())
		c.AbortWithStatusJSON(http.StatusInternalServerError, baseResponse)
		return
	}
	if len(datas) > 0 {
		baseResponse.Data = datas
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusOK, "SUCCESS", "delete data success", time.Since(start).Seconds())
		c.JSON(http.StatusOK, baseResponse)
	}
	baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusNotFound, "FAILED", "data not found", time.Since(start).Seconds())
	c.AbortWithStatusJSON(http.StatusNotFound, baseResponse)
}
