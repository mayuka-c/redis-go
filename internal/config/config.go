package config

import (
	"context"

	"github.com/kelseyhightower/envconfig"
)

type RedisConfig struct {
	Endpoint string `envconfig:"REDIS_ENDPOINT" required:"true"`
}

func GetRedisConfig(ctx context.Context) RedisConfig {
	redisConfig := RedisConfig{}
	err := envconfig.Process("", &redisConfig)
	if err != nil {
		panic(err)
	}
	return redisConfig
}
