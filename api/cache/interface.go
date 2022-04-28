package cache

import (
	"context"
	"errors"
	"log"
	"reflect"

	"github.com/eko/gocache/v2/store"
	"github.com/gin-gonic/gin"
)

var debug bool

func Debug() {
	debug = true
}

//缓存接口
type Interface interface {
	CacheKey() (cacheKeys []string)
	Option() *store.Options
}

//自定义缓存value
type CustomCacheValue interface {
	CacheValue() interface{}
}

func SetCacheOf(ctx context.Context, c Interface) {
	var err error
	for _, cacheKey := range c.CacheKey() {
		if ci, ok := c.(CustomCacheValue); ok {
			err = marshal.Set(ctx, cacheKey, ci.CacheValue(), c.Option())
		} else {
			if gc, o := ctx.(*gin.Context); o {
				gc.Set(cacheKey, c)
			}
			err = marshal.Set(ctx, cacheKey, c, c.Option())
		}
		if debug && err != nil {
			log.Printf("[cache] set cache error: cacheKey(%s) value(%s) err(%s)", cacheKey, c, err)
		}
	}
}

func GetCacheOf(ctx context.Context, c Interface) error {
	var err error
	for _, cacheKey := range c.CacheKey() {
		//对于自定义内容的缓存不做处理
		if _, ok := c.(CustomCacheValue); !ok {
			if gc, o := ctx.(*gin.Context); o {
				if v, exists := gc.Get(cacheKey); exists {
					c = v.(Interface)
				}
			}
		}
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
	for _, key := range c.CacheKey() {
		if gc, ok := ctx.(*gin.Context); ok {
			gc.Set(key, nil)
		}
		err := marshal.Delete(ctx, c.CacheKey())
		if err != nil {
			if debug {
				log.Printf("[cache] delete cache error: key(%s) err(%s)", key, err)
			}
		}
	}
}
