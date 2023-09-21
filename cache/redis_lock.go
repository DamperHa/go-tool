package cache

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

// Client 是对redis.Cmdable的封装
type Client struct {
	client redis.Cmdable
}

// 需要考虑的问题
// 锁是否需要过期时间呢？
// 如果加了过期时间，但是自己完成的业务超过过期时间了，怎么办呢？
// 为什么用uuid作为值呢？保证这把锁，是属于某个实例
var (
	ErrFailedToPreemptLock = errors.New("redis-lock:抢锁失败")
	ErrLockNoExist         = errors.New("redis-lock:锁不存在")
)

func (c *Client) TryLock(ctx context.Context, key string, expiration time.Duration) (*Lock, error) {
	val := uuid.New().String()

	ok, err := c.client.SetNX(ctx, key, val, expiration).Result()
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrFailedToPreemptLock
	}

	return &Lock{
		client: c.client,
		key:    key,
		value:  val,
	}, nil
}

type Lock struct {
	client redis.Cmdable
	key    string
	value  string
}

// UnLock Check and do something
// check and do 需要保证是原子操作
//func (l *Lock) UnLock(ctx context.Context, key string) error {
//	// 先判断这把锁是不是我的锁
//
//	// 把键值对删除
//	cnt, err := l.client.Del(ctx, l.key).Result()
//	if err != nil {
//		return err
//	}
//
//	if cnt != 1 {
//		return ErrLockNoExist
//	}
//
//	return nil
//}
