package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client

func InitRedis() (err error) {
	RedisDB = redis.NewClient(&redis.Options{
		Addr: "124.221.179.105:6379",
	})
	_, err = RedisDB.Ping(context.Background()).Result()
	return err
}

func RedisClose() {
	if err := RedisDB.Close(); err != nil {
		panic(err)
	}
}
