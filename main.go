package main

import (
	"context"
	"go-svc-management/src/routers"
)

func main() {
	router, mongo, redis := routers.InitRouters()
	defer mongo.Disconnect(context.TODO())
	defer redis.Close()
	router.Run(":8080")
}
