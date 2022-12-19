package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mfturkcanoglu/go-mux-clean/entities"
)

type redisCache struct {
	host    string
	db      int // 0 - 15
	expires time.Duration
}

type RedisConfig struct {
	Host       string `mapstructure:"host"`
	DB         uint   `mapstructure:"db"`
	Expiration uint   `mapstructure:"expiration"`
}

var (
	ctx context.Context = context.Background()
)

func NewRedisCache(host string, db int, exp time.Duration) ProductCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		DB:       cache.db,
		Addr:     cache.host,
		Password: "",
	})
}

func (cache *redisCache) Set(key string, value *entities.Product) {
	client := cache.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(ctx, key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *entities.Product {
	client := cache.getClient()
	var bytes, err = client.Get(ctx, key).Result()
	if err != nil {
		// log.Panicln(err, "Client getting value error.", "Key", key)
		log.Println("Cannot find value for key", key)
		return nil
	}
	var value entities.Product
	err = json.Unmarshal([]byte(bytes), &value)
	if err != nil {
		panic(err)
	}
	return &value
}
