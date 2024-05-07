package database

import (
	"fmt"
	"github.com/JimYcod3x/meter_server/config"
	redis "github.com/go-redis/redis/v8"
)

func ConnectionRedisDb(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})

	fmt.Println("Connected Successful to the dabase(redis)")

	return rdb
}