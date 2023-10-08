package cache

import (
	"context"
	"gin-demo/config"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
	Rctx context.Context
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: config.RedisAddress,
		Password: config.RedisPassword,
		DB: config.RedisDb,
	})
	Rctx = context.Background()
}

func Zscore(id, score int) redis.Z {
    return redis.Z{Score: float64(score), Member: id}
}