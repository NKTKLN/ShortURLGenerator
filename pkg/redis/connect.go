package redis

import (
	"github.com/go-redis/redis"
)
// Connecting to the db
func RedisClientConnect() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})
}