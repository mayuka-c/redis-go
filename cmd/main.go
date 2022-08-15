package main

import (
	"context"

	"github.com/mayuka-c/redis-go/internal/config"
	"github.com/mayuka-c/redis-go/internal/service"
)

func main() {
	ctx := context.Background()

	redisConfig := config.GetRedisConfig(ctx)
	svc := service.NewService(redisConfig)
	// svc.AddValues()
	svc.GetValues()
}
