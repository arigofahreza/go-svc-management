package services

import (
	"go-svc-management/src/models"
	"go-svc-management/src/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type SessionServices struct {
	Data []*models.SessionModel `json:"data,omitempty"`
}

func InitSessionServices() *SessionServices {
	return &SessionServices{}
}

const (
	SESSION_EXPIRE        = 24 * time.Hour
	SESSION_KEY           = "dsip"
	CACHE_PREFIX_DEVICE   = "session:device"
	CACHE_PREFIX_LOGIN    = "session:login"
	CACHE_PREFIX_REGISTER = "session:register"
	PATH                  = "/"
	DOMAIN                = "localhost"
)

func (service SessionServices) GenerateSession(ctx *gin.Context, redisClient *redis.Client) ([]*models.SessionModel, error) {
	var session string
	err := service.DeleteSession(ctx, redisClient)
	if err != nil {
		return nil, err
	}
	userAgent := ctx.Request.Header["User-Agent"]
	origin := ctx.Request.Header["Origin"]
	device := map[string]interface{}{
		"user_agent": userAgent,
		"origin":     origin,
	}
	session, err = utils.EncryptDevice(device)
	cacheKey := strings.Join([]string{CACHE_PREFIX_DEVICE, session}, ":")
	if err != nil {
		return nil, err
	}
	if session != "" {
		redisClient.Set(cacheKey, session, SESSION_EXPIRE)
		ctx.SetCookie(SESSION_KEY, session, int(time.Duration(SESSION_EXPIRE).Seconds()), PATH, DOMAIN, false, false)
		service.Data = append(service.Data, &models.SessionModel{
			SessionId: session,
		})
		return service.Data, nil
	}
	return nil, err
}

func (service SessionServices) GetSession(ctx *gin.Context, redisClient *redis.Client) (string, error) {
	session, err := ctx.Cookie(SESSION_KEY)
	if err != nil {
		return "", err
	}
	cacheKey := strings.Join([]string{CACHE_PREFIX_DEVICE, session}, ":")
	deviceData, err := redisClient.Get(cacheKey).Result()
	if err != redis.Nil {
		return "", nil
	}
	return deviceData, nil

}

func (service SessionServices) DeleteSession(ctx *gin.Context, redisClient *redis.Client) error {
	session, err := ctx.Cookie(SESSION_KEY)
	if session != "" {
		if err != nil && err != http.ErrNoCookie {
			return err
		}
		cacheKey := strings.Join([]string{CACHE_PREFIX_DEVICE, session}, ":")
		_, err := redisClient.Get(cacheKey).Result()
		if err != redis.Nil {
			redisClient.Del(cacheKey)
		}
		return nil
	}
	return nil
}
