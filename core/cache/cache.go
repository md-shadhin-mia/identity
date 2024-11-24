package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
)

// Cache interface to unify different cache implementations
type Cache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (interface{}, error)
}

// InMemoryCache implementation
type InMemoryCache struct {
	store *cache.Cache
}

func (imc *InMemoryCache) Set(key string, value interface{}, expiration time.Duration) error {
	imc.store.Set(key, value, expiration)
	return nil
}

func (imc *InMemoryCache) Get(key string) (interface{}, error) {
	value, found := imc.store.Get(key)
	if !found {
		return nil, fmt.Errorf("key not found")
	}
	return value, nil
}

// RedisCache implementation
type RedisCache struct {
	client *redis.Client
}

func (rc *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return rc.client.Set(context.Background(), key, value, expiration).Err()
}

func (rc *RedisCache) Get(key string) (interface{}, error) {
	val, err := rc.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("key not found")
	} else if err != nil {
		return nil, err
	}
	return val, nil
}

var CacheInstance Cache

func InitializeCache() {
	isRedis := os.Getenv("IS_REDIS")
	if isRedis == "true" {
		address := os.Getenv("REDIS_ADDRESS")
		password := os.Getenv("REDIS_PASSWORD")

		client := redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
		})
		CacheInstance = &RedisCache{client: client}
	} else {
		CacheInstance = &InMemoryCache{store: cache.New(5*time.Minute, 10*time.Minute)}
	}

}
