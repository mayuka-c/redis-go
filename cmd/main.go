package main

import (
	"context"

	"github.com/mayuka-c/redis-go/internal/client"
	"github.com/mayuka-c/redis-go/internal/config"
	"github.com/mayuka-c/redis-go/internal/model"
)

func main() {
	ctx := context.Background()

	item := model.Student{
		Name:    "Elliot",
		Age:     "25",
		Hobbies: []string{"music", "games"},
	}

	redisConfig := config.GetRedisConfig(ctx)
	// svc := service.NewService(redisConfig)
	// svc.AddValues()
	// svc.GetValues()
	// svc.AddValues()

	client.DifferentPackage(redisConfig, item)
}
