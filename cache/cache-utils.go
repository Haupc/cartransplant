package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdm *redis.Client
	ctx = context.Background()
)

const (
	serviceName = "auth"
)

func init() {
	GetCacheManager()
}

// Cache ...
type Cache interface {
	Get(key interface{}) (interface{}, error)
	Evict(key interface{})
	Set(key, value string, ttl time.Duration)
}

// BaseCache ...
type baseCache struct {
	service   string
	cacheName string
	LoadFunc  func(string) (string, error)
}

func (c baseCache) getCacheKey(key string) string {
	return fmt.Sprintf("%s.%s.%s", c.service, c.cacheName, key)
}

func (c baseCache) Get(k interface{}) (string, error) {
	key := fmt.Sprintf("%v", k)
	cacheKey := c.getCacheKey(key)
	result, err := rdm.Get(ctx, cacheKey).Result()
	switch {
	case err == redis.Nil:
		result, err := c.LoadFunc(key)
		if err != nil {
			return "", err
		}
		GetCacheManager().Set(ctx, cacheKey, result, time.Hour)
		return result, err
	case err != nil:
		return "", err
	default:
		return result, nil
	}
}

func (c baseCache) Evict(k interface{}) {
	key := fmt.Sprintf("%v", k)
	cacheKey := c.getCacheKey(key)
	if err := GetCacheManager().Del(ctx, cacheKey).Err(); err != nil {
		log.Println(err.Error())
	}
}

func (c baseCache) Set(key, value string, ttl time.Duration) {
	cacheKey := c.getCacheKey(key)
	GetCacheManager().Set(ctx, cacheKey, value, ttl)
}

// GetCacheManager ...
func GetCacheManager() *redis.Client {
	if rdm == nil {
		rdm = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	}
	return rdm
}
