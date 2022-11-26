package locker

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Locker struct {
	redisClient *redis.Client
}

var instance *Locker

func Init(redisClient *redis.Client) error {
	if redisClient == nil {
		return errors.New("init Locker error: redisClient can't be nil")
	}
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	instance = &Locker{redisClient: redisClient}
	return nil
}

func Lock(ctx context.Context, lockKey string, options ...LockOption) (token string, err error) {
	acquireTimeOut := LockOptions(options).acquireTimeOut()
	lockTimeOut := LockOptions(options).lockTimeOut()

	token = uuid.NewV4().String()
	var acquireDuration = acquireTimeOut

	for acquireDuration > 0 {
		ok, err := instance.redisClient.SetNX(ctx, lockKey, token, lockTimeOut).Result()
		if err != nil {
			fmt.Printf("redis锁 - 加锁失败 key: %s, token: %s, ttl: %v, err: %v\r\n", lockKey, token, lockTimeOut, err)
			return "", err
		}

		if ok {
			//fmt.Printf("redis锁 - 加锁成功 key: %s, token: %s, ttl: %v\r\n", lockKey, token, lockTimeOut)
			return token, nil
		}
		time.Sleep(time.Millisecond * 100) // next: sleep 500ms
		acquireDuration -= time.Millisecond * 100
	}
	fmt.Printf("redis锁 - 加锁超时 key: %s, token: %s, ttl: %v\r\n", lockKey, token, acquireTimeOut)
	return "", fmt.Errorf("获取redis锁超时: %v\r\n", acquireTimeOut)
}

func Unlock(ctx context.Context, lockKey, token string) error {
	const script = `if redis.call('GET', KEYS[1]) == ARGV[1] then return redis.call('DEL', KEYS[1]) else return 0 end`
	ret, err := instance.redisClient.Eval(ctx, script, []string{lockKey}, token).Result()
	if err != nil {
		fmt.Printf("redis锁 - 解锁失败 key: %s, token: %s,  err: %v\r\n", lockKey, token, err)
		return err
	}

	if v, ok := ret.(int64); ok && v == 1 {
		//fmt.Printf("redis锁 - 解锁成功 key: %s, token: %s\r\n", lockKey, token)
		return nil
	}

	fmt.Printf("redis锁 - 解锁失败 key: %s, token: %s, key已经过期\r\n", lockKey, token)
	return nil
}
