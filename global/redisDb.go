package global

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	RedisDb *redis.Client
)

//创建redis链接
func SetupRedisDb() (error) {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
