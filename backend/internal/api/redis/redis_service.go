package redis

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/solomonjdavid001/Dockernetes/backend/config"
)

var (
	redisClient *redis.Client
	once        sync.Once
)

func getRedisClient() *redis.Client {
	once.Do(func() {
		redisClient = initRedisClient()
	})
	return redisClient
}

func initRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.GlobalConfig.Redis.Server,
	})
	return redisClient
}

func SetValue(key string, value string) error {
	client := getRedisClient()
	ctx := context.Background()

	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatal("Error setting Redis value: ", err)
		return err
	}
	return nil
}
