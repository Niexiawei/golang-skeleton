package database

import (
	"context"
	"fmt"
	"github.com/Niexiawei/golang-skeleton/config"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

var RedisCtx = context.Background()

func SetupRedis() {
	redisConfig := config.Instance.Redis
	redisClient = redis.NewClient(&redis.Options{
		Password: redisConfig.Password,
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		DB:       redisConfig.Db,
		PoolSize: redisConfig.Pool,
	})
}

func GetRedisConn() *redis.Client {
	return redisClient
}
