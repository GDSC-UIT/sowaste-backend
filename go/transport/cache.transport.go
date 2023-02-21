package transport

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Source *redis.Client
}

var Redis RedisCache

func (rc *RedisCache) InitRedis() {
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisConnectionURL := os.Getenv("REDIS_CONNECTION_URL")
	redisConnectionPort := os.Getenv("REDIS_CONNECTION_PORT")
	Redis.Source = redis.NewClient(&redis.Options{
		Addr:     redisConnectionURL + ":" + redisConnectionPort,
		Password: redisPassword,
		DB:       0, // use default DB
	})
	fmt.Println("Redis connected")
}

func (rc *RedisCache) DeleteValue(key string) {
	count := Redis.Source.Del(context.Background(), key)
	fmt.Println(count.Val())
}

func (rc *RedisCache) DeleteAllValue() {
	Redis.Source.FlushAll(context.Background())
}

func (rc *RedisCache) DisconnetRedis() {
	Redis.Source.Close()
}

func (rc *RedisCache) SetValue(key string, value string) {
	Redis.Source.Set(context.Background(), key, value, time.Hour*24*3)
}

func (rc *RedisCache) GetValue(key string) (string, error) {
	value, err := Redis.Source.Get(context.Background(), key).Result()
	return value, err
}
