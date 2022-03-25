package utils

import (
	"log"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func ConnectRedis() {
	var (
		host = GetEnv("REDIS_HOST", "localhost")
		port = string(GetEnv("REDIS_PORT", "6379"))
	)

	client := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})

	Redis = client

	pong, _ := Redis.Ping().Result()
	if pong == "PONG" {
		log.Println("redis: connected")
	}
}
