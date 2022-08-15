package client

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/fatih/structs"
	"github.com/go-redis/redis"
	"github.com/mayuka-c/redis-go/internal/config"
	"github.com/mayuka-c/redis-go/internal/model"
	"github.com/mayuka-c/redis-go/internal/utils"
	"github.com/nitishm/go-rejson"
)

type InterfaceClient interface {
	CreateItem(value string)
	GetItem(id string) string
	CreateItemWithHash(input model.Student)
	GetItemWithHash(id string) model.Student
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

func (r *redisClient) CreateItem(value string) {

	err := r.rdb.Set("id-1", value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (r *redisClient) GetItem(id string) string {

	val := r.rdb.Get(id).String()
	return val
}

func (r *redisClient) CreateItemWithHash(input model.Student) {

	hashInput := structs.Map(input)

	fmt.Printf("Hashinput: %#v\n", hashInput)

	val, err := r.rdb.HMSet("id-1", hashInput).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HMSET Val: ", val)
}

func (r *redisClient) GetItemWithHash(id string) model.Student {

	val, err := r.rdb.HGetAll(id).Result()
	if err != nil {
		panic(err)
	}

	out := model.Student{}
	utils.DeepCopy(&val, &out)

	return out
}

func DifferentPackage(redisConfig config.RedisConfig, input model.Student) {

	var addr = flag.String("Server", redisConfig.Endpoint, "Redis server address")
	rh := rejson.NewReJSONHandler()
	flag.Parse()
	cli := redis.NewClient(&redis.Options{Addr: *addr})
	rh.SetGoRedisClient(cli)

	res, err := rh.JSONSet("id-`1", ".", input)
	if err != nil {
		log.Fatalf("Failed to JSONSet: %v", err)
		return
	}

	res, err = rh.JSONGet("id-`1", ".")
	if err != nil {
		log.Fatalf("Failed to JSONGet")
	}

	out := model.Student{}
	json.Unmarshal(res.([]byte), &out)

	fmt.Printf("Values: #%v\n", out)
}
