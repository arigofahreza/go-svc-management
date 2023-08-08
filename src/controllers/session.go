package controllers

import (
	"go-svc-management/src/models"
	"go-svc-management/src/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type SessionControllers struct {
	RedisClient     *redis.Client
	SessionServices *services.SessionServices
}

func InitSessionControllers(redisClient *redis.Client) *SessionControllers {
	return &SessionControllers{
		RedisClient:     redisClient,
		SessionServices: services.InitSessionServices(),
	}
}

func (controller SessionControllers) GenerateSessionController(c *gin.Context) {
	start := time.Now()
	var baseResponse models.BaseResponse
	sessionId, err := controller.SessionServices.GenerateSession(c, controller.RedisClient)
	if err != nil {
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusInternalServerError, "FAILED", "generate session error", time.Since(start).Seconds())
		c.AbortWithStatusJSON(http.StatusInternalServerError, baseResponse)
		return
	}
	if sessionId != nil {
		baseResponse.Data = sessionId
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusOK, "SUCCESS", "generate session success", time.Since(start).Seconds())
		c.JSON(http.StatusOK, baseResponse)
	} else {
		baseResponse.Metadata = baseResponse.OverrideMetadata(http.StatusInternalServerError, "FAILED", "session_id null", time.Since(start).Seconds())
		c.AbortWithStatusJSON(http.StatusInternalServerError, baseResponse)
		return
	}
}
