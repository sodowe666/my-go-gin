package core

import (
	"github.com/go-redis/redis"
	"sync"
	"time"
	"web/config"
	"fmt"
)

func init() {
	initRedis()
}

type redisClient struct {
	once   sync.Once
	client *redis.Client
}

var redisC redisClient

func Redis() *redis.Client {
	if redisC.client == nil {
		initRedis()
	}
	return redisC.client
}

func initRedis() {
	redisC.once.Do(func() {
		if redisC.client == nil {
			cfg := config.GetConfig()
			redisCfg := cfg.Redis
			password := redisCfg.Password
			host := redisCfg.Host
			port := redisCfg.Port
			database := redisCfg.Database
			idleTime := redisCfg.IdleTime
			poolSize := redisCfg.MaxActive
			readTimeout := redisCfg.ReadTimeout
			writeTimeout := redisCfg.WriteTimeout
			redisC.client = redis.NewClient(&redis.Options{
				Addr:         host + ":" + port,
				DB:           database,
				IdleTimeout:  time.Duration(idleTime),
				PoolSize:     poolSize,
				ReadTimeout:  time.Duration(readTimeout) * time.Second,
				WriteTimeout: time.Duration(writeTimeout) * time.Second,
				Password:     password,
			})
			//defer Redis.Close()
			_, err := redisC.client.Ping().Result()
			if err != nil {
				panic("redis init fail. " + string(err.Error()))
			}
			fmt.Println("Redis Start Success")
		}
	})
}
