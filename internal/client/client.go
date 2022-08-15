package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/mayuka-c/redis-go/internal/config"
	"github.com/mayuka-c/redis-go/internal/model"
)

type InterfaceClient interface {
	CreateItem(input model.Student)
	GetItem(id string) model.Student
}

type redisClient struct {
	rdb *redis.Client
}

func NewRedisClient(redisConfig config.RedisConfig) *redisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Endpoint,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &redisClient{
		rdb: rdb,
	}
}

func (r *redisClient) CreateItem(input model.Student) {

	json, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
	}

	err = r.rdb.Set("id-1", json, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (r *redisClient) GetItem(id string) model.Student {
	val, err := r.rdb.Get(id).Bytes()
	if err != nil {
		fmt.Println(err)
	}
	out := model.Student{}
	json.Unmarshal(val, &out)

	return out
}
