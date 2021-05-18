package cache

import (
	"encoding/json"
	"time"

	"github.com/haupc/cartransplant/auth/dto"
	"github.com/haupc/cartransplant/cache/loader"
)

var (
	userCacheName = "user"
	_userCache    *userCache
)

type userCache struct {
	baseCache
}

func (u *userCache) Get(key interface{}) (interface{}, error) {
	var retVal dto.UserDTO
	result, err := u.baseCache.Get(key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result), &retVal)
	return retVal, err
}

func (u *userCache) Evict(key interface{}) {
	u.baseCache.Evict(key)
}

func (u *userCache) Set(key, value string, ttl time.Duration) {
	u.baseCache.Set(key, value, ttl)
}

// GetUserCache ...
func GetUserCache() Cache {
	if _userCache == nil {
		_userCache = &userCache{
			baseCache{
				service:   serviceName,
				cacheName: userCacheName,
				LoadFunc:  loader.GetUserCacheLoader().Load,
			},
		}

	}
	return _userCache
}
