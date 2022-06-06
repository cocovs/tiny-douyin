package controller

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})
	ctx = context.Background()
)

//func initRedis(redisDb *redis.Client) (err error) {
//	redisDb = redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "123456",
//		DB:       0,
//	})
//	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
//	defer cancelFunc()
//	_, err = redisDb.Ping(timeoutCtx).Result()
//	if err != nil {
//		return err
//	}
//	return nil
//}
