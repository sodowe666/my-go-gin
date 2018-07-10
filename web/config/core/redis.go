package core

import (
	"github.com/go-redis/redis"
	"sync"
	"time"
	"web/config"
)

func init() {
	initRedis()
}

var redisClient *redis.Client
var redisOnce sync.Once

func Redis() *redis.Client {
	if redisClient == nil {
		initRedis()
	}
	return redisClient
}

func initRedis() {
	redisOnce.Do(func() {
		if redisClient == nil {
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
			redisClient = redis.NewClient(&redis.Options{
				Addr:         host + ":" + port,
				DB:           database,
				IdleTimeout:  time.Duration(idleTime),
				PoolSize:     poolSize,
				ReadTimeout:  time.Duration(readTimeout) * time.Second,
				WriteTimeout: time.Duration(writeTimeout) * time.Second,
				Password:     password,
			})
			//defer Redis.Close()
			_, err := redisClient.Ping().Result()
			if err != nil {
				panic("redis init fail. " + string(err.Error()))
			}
		}
	})
}
