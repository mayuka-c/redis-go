package service

import (
	"fmt"

	"github.com/mayuka-c/redis-go/internal/client"
	"github.com/mayuka-c/redis-go/internal/config"
	"github.com/mayuka-c/redis-go/internal/model"
)

type service struct {
	client client.InterfaceClient
}

func NewService(redisConfig config.RedisConfig) *service {
	return &service{
		client: client.NewRedisClient(redisConfig),
	}
}

func (s *service) AddValues() {

	item := model.Student{
		Name:    "Elliot",
		Age:     "25",
		Hobbies: []string{"music", "games"},
	}
	s.client.CreateItemWithHash(item)
}

func (s *service) GetValues() {
	out := s.client.GetItemWithHash("id-1")
	fmt.Printf("Received Item: %#v\n", out)
}
