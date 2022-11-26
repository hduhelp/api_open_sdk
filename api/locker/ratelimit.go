package locker

import (
	"context"
	"fmt"
	"time"
)

const (
	// RateLimitTtl 请求记录在redis的存储时间
	RateLimitTtl = time.Minute
)

func AddRateLimitRecord(ctx context.Context, key string, period time.Duration, maxTimes int64) (hitRateLimit bool, err error) {
	// 初始化计数器（如果不存在）
	totalKey := getTotalKey(key, time.Now())
	_, err = instance.redisClient.SetNX(ctx, totalKey, 0, RateLimitTtl).Result()
	if err != nil {
		return false, err
	}

	// 得到这是第几次请求
	index, err := instance.redisClient.Incr(ctx, totalKey).Result()
	if err != nil {
		return false, err
	}

	// 存储当前请求时间
	timeKey := getTimeKey(key, time.Now(), index)
	_, err = instance.redisClient.Set(ctx, timeKey, time.Now(), RateLimitTtl).Result()
	if err != nil {
		return false, err
	}

	// 检查是否超过频率限制
	return checkRateLimit(ctx, key, time.Now(), index, period, maxTimes)
}

// 检查是否超过频率限制
func checkRateLimit(ctx context.Context, key string, reqTime time.Time, reqIndex int64, period time.Duration, maxTimes int64) (bool, error) {
	// 计算用于对比的频率标识
	compareIndex := reqIndex - maxTimes
	if compareIndex <= 0 {
		// 今日请求次数还未达到最大次数，通过
		return false, nil
	}

	// 获取对比请求的请求时间
	compareKey := getTimeKey(key, time.Now(), compareIndex)
	exist, err := instance.redisClient.Exists(ctx, compareKey).Result()
	if err != nil {
		return false, err
	}
	if exist == 0 {
		//记录时间的Key，已过期，由于过期时间大于period，故不用比较时间
		return false, nil
	}
	compareTime, err := instance.redisClient.Get(ctx, compareKey).Time()
	if err != nil {
		return false, err
	}
	if compareTime.Add(period).Before(reqTime) {
		return false, nil
	} else {
		return true, nil
	}
}

// 计算统计请求总次数的缓存key
func getTotalKey(key string, time time.Time) string {
	date := time.Format("2006-01-02")
	return fmt.Sprintf("TotalKey|%s|%s", date, key)
}

// 计算统计请求时间的缓存key
func getTimeKey(key string, time time.Time, index int64) string {
	date := time.Format("2006-01-02")
	return fmt.Sprintf("TimeKey|%s|%s|%d", date, key, index)
}
