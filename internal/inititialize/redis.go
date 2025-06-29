package inititialize

import (
	"context"
	"fmt"

	"github.com/quockhanhcao/ecommerce/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	config := global.Config.RedisConfig
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis init failed", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Redis init ok", zap.String("ok", "success"))
	fmt.Println("Redis init ok")
	global.Rdb = rdb
    redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		global.Logger.Error("Failed to set value in Redis", zap.Error(err))
		return
	}

	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		global.Logger.Error("Failed to get value from Redis", zap.Error(err))
		return
	}

	global.Logger.Info("Value from Redis", zap.String("score", value))
}
