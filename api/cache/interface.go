package cache

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/eko/gocache/v2/store"
)

var debug bool

// Interface 缓存接口
type Interface interface {
	CacheKey() (cacheKeys []string)
	Option() *store.Options
}

// CustomCacheValue 自定义缓存value
type CustomCacheValue interface {
	CacheValue() interface{}
}

func SetCacheOf(ctx context.Context, c Interface) {
	var err error
	for _, cacheKey := range c.CacheKey() {
		if ci, ok := c.(CustomCacheValue); ok {
			err = marshal.Set(ctx, cacheKey, ci.CacheValue(), c.Option())
		} else {
			err = marshal.Set(ctx, cacheKey, c, c.Option())
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetCacheOf(ctx context.Context, c Interface) error {
	var err error
	for _, cacheKey := range c.CacheKey() {
		if ci, ok := c.(CustomCacheValue); ok {
			if !reflect.ValueOf(ci.CacheValue()).Elem().CanAddr() {
				break
			}
			_, err = marshal.Get(ctx, cacheKey, ci.CacheValue())
		} else {
			_, err = marshal.Get(ctx, cacheKey, c)
		}
		if err != nil {
			if debug {
				if ci, ok := c.(CustomCacheValue); ok {
					log.Printf("[cache] get cache faild: cacheKey(%s) value(%s) err(%s)", cacheKey, ci.CacheValue(), err)
				} else {
					log.Printf("[cache] get cache faild: cacheKey(%s) value(%s) err(%s)", cacheKey, c, err)
				}
			}
		} else {
			return nil
		}
	}
	return errors.New("cache not found")
}

func DeleteCacheOf(ctx context.Context, c Interface) {
	for _, cacheKey := range c.CacheKey() {
		err := marshal.Delete(ctx, cacheKey)
		if err != nil {
			if debug {
				log.Printf("[cache] delete cache error: key(%s) err(%s)", cacheKey, err)
			}
		}
	}
}
