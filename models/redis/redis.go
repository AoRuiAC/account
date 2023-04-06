package redis

import (
	"account/dao/redis"
	"context"
)

func SetLink(id1, id2 string) (err error) {
	return redis.RedisDB.SAdd(context.Background(), id1, id2).Err()
}

func MSetLink(id1 string, id2 []string) (err error) {
	return redis.RedisDB.SAdd(context.Background(), id1, id2).Err()
}

func GetLinkById(id string) (res []string, err error) {
	res, err = redis.RedisDB.SMembers(context.Background(), id).Result()
	return res, err
}

func DeleteLink(id1, id2 string) (err error) {
	return redis.RedisDB.SRem(context.Background(), id1, id2).Err()
}
