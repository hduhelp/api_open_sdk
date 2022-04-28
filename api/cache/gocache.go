package cache

import (
	"time"

	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/marshaler"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	goCache "github.com/patrickmn/go-cache"
)

var marshal *marshaler.Marshaler

func SetUp(caches ...cache.SetterCacheInterface) {
	cacheManager := cache.NewChain(caches...)
	marshal = marshaler.New(cacheManager)
}

func DefaultMemCache() *cache.Cache {
	return MemCache(5*time.Minute, 10*time.Minute)
}

func MemCache(defaultExpiration time.Duration, cleanupInterval time.Duration) *cache.Cache {
	return cache.New(store.NewGoCache(goCache.New(defaultExpiration, cleanupInterval), nil))
}

func RedisCache(addr string) *cache.Cache {
	return cache.New(store.NewRedis(redis.NewClient(&redis.Options{
		Addr: addr,
	}), &store.Options{Expiration: 5 * time.Second}))
}
