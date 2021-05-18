package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	tokenCacheName = "token"
	_tokenCache    *tokenCache
)

type tokenCache struct {
	baseCache
}

func (t *tokenCache) Get(k interface{}) (interface{}, error) {
	key := fmt.Sprintf("%v", k)
	cacheKey := t.getCacheKey(key)
	result, err := GetCacheManager().Get(ctx, cacheKey).Result()
	switch {
	case err == redis.Nil:
		return nil, nil
	case err != nil:
		return "", err
	default:
		return result, nil
	}
}

func (t *tokenCache) Evict(key interface{}) {
	t.baseCache.Evict(key)
}

func (t *tokenCache) Set(key, value string, ttl time.Duration) {
	t.baseCache.Set(key, value, ttl)
}

// GetTokenCache ...
func GetTokenCache() Cache {
	if _tokenCache == nil {
		_tokenCache = &tokenCache{
			baseCache: baseCache{
				service:   serviceName,
				cacheName: tokenCacheName,
				LoadFunc:  func(s string) (string, error) { return "", nil },
			},
		}
	}
	return _tokenCache
}
